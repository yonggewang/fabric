/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package deliverservice

import (
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/orderer"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/common/util"
	bftBlocksprovider "github.com/hyperledger/fabric/core/deliverservice/blocksprovider"
	"github.com/hyperledger/fabric/internal/pkg/comm"
	"github.com/hyperledger/fabric/internal/pkg/identity"
	"github.com/hyperledger/fabric/internal/pkg/peer/blocksprovider"
	"github.com/hyperledger/fabric/internal/pkg/peer/orderers"
	"google.golang.org/grpc"
)

var logger = flogging.MustGetLogger("deliveryClient")

// DeliverService used to communicate with orderers to obtain
// new blocks and send them to the committer service
type DeliverService interface {
	// StartDeliverForChannel dynamically starts delivery of new blocks from ordering service
	// to channel peers.
	// When the delivery finishes, the finalizer func is called
	StartDeliverForChannel(chainID string, ledgerInfo blocksprovider.LedgerInfo, finalizer func()) error

	// StopDeliverForChannel dynamically stops delivery of new blocks from ordering service
	// to channel peers.
	StopDeliverForChannel(chainID string) error

	// UpdateEndpoints updates the ordering endpoints for the given chain.
	UpdateEndpoints(chainID string, endpoints []*orderers.Endpoint) error

	// Stop terminates delivery service and closes the connection
	Stop()
}

// deliverServiceImpl the implementation of the delivery service
// maintains connection to the ordering service and maps of
// blocks providers
type deliverServiceImpl struct {
	conf           *Config
	blockProviders map[string]*blocksprovider.Deliverer
	lock           sync.RWMutex
	stopping       bool
}

// Config dictates the DeliveryService's properties,
// namely how it connects to an ordering service endpoint,
// how it verifies messages received from it,
// and how it disseminates the messages to other peers
type Config struct {
	IsStaticLeader bool
	// CryptoSvc performs cryptographic actions like message verification and signing
	// and identity validation.
	CryptoSvc blocksprovider.BlockVerifier
	// Gossip enables to enumerate peers in the channel, send a message to peers,
	// and add a block to the gossip state transfer layer.
	Gossip blocksprovider.GossipServiceAdapter
	// OrdererSource provides orderer endpoints, complete with TLS cert pools.
	OrdererSource *orderers.ConnectionSource
	// Signer is the identity used to sign requests.
	Signer identity.SignerSerializer
	// GRPC Client
	DeliverGRPCClient *comm.GRPCClient
	// Configuration values for deliver service.
	// TODO: merge 2 Config struct
	// DeliverServiceConfig is the configuration object.
	DeliverServiceConfig *DeliverServiceConfig
}

// NewDeliverService construction function to create and initialize
// delivery service instance. It tries to establish connection to
// the specified in the configuration ordering service, in case it
// fails to dial to it, return nil
func NewDeliverService(conf *Config) DeliverService {
	ds := &deliverServiceImpl{
		conf:           conf,
		blockProviders: make(map[string]*blocksprovider.Deliverer),
	}
	return ds
}

// DialerAdapter implements the creation of a new gRPC connection
type DialerAdapter struct {
	Client *comm.GRPCClient
}

// Dial creates a new gRPC connection
func (da DialerAdapter) Dial(address string, certPool *x509.CertPool) (*grpc.ClientConn, error) {
	return da.Client.NewConnection(address, comm.CertPoolOverride(certPool))
}

type DeliverAdapter struct{}

// Deliver returns a stream client
func (DeliverAdapter) Deliver(ctx context.Context, clientConn *grpc.ClientConn) (blocksprovider.StreamClient, error) {
	abc, err := orderer.NewAtomicBroadcastClient(clientConn).Deliver(ctx)
	if err != nil {
		return nil, err
	}
	return &deliverClient{abc: abc}, nil
}

type deliverClient struct {
	abc orderer.AtomicBroadcast_DeliverClient
}

// Send sends an envelope to the ordering service
func (d deliverClient) Send(envelope *common.Envelope) error {
	return d.abc.Send(envelope)
}

// Recv receives a chaincode message
func (d deliverClient) Recv() (*orderer.DeliverResponse, error) {
	return d.abc.Recv()
}

// CloseSend closes the client connection
func (d deliverClient) CloseSend() error {
	d.abc.CloseSend()
	return nil
}

// Disconnect does nothing
func (d deliverClient) Disconnect() {
}

// UpdateReceived does nothing
func (d deliverClient) UpdateReceived(blockNumber uint64) {
}

// StartDeliverForChannel starts blocks delivery for channel
// initializes the grpc stream for given chainID, creates blocks provider instance
// that spawns in go routine to read new blocks starting from the position provided by ledger
// info instance.
func (d *deliverServiceImpl) StartDeliverForChannel(chainID string, ledgerInfo blocksprovider.LedgerInfo, finalizer func()) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.stopping {
		errMsg := fmt.Sprintf("Delivery service is stopping cannot join a new channel %s", chainID)
		logger.Errorf(errMsg)
		return errors.New(errMsg)
	}
	if _, exist := d.blockProviders[chainID]; exist {
		errMsg := fmt.Sprintf("Delivery service - block provider already exists for %s found, can't start delivery", chainID)
		logger.Errorf(errMsg)
		return errors.New(errMsg)
	}
	logger.Info("This peer will retrieve blocks from ordering service and disseminate to other peers in the organization for channel", chainID)

	var bp blocksprovider.BlocksProvider

	dialer := DialerAdapter{
		Client: d.conf.DeliverGRPCClient,
	}


	if d.conf.DeliverServiceConfig.IsBFT {
		bftClient, _ := NewBFTDeliveryClient(chainID, d.conf.OrdererSource, ledgerInfo, d.conf.CryptoSvc, d.conf.Signer, d.conf.DeliverGRPCClient, dialer)
		bp = bftBlocksprovider.NewBlocksProvider(chainID, bftClient, d.conf.Gossip, d.conf.CryptoSvc)
	} else {
		dc := &blocksprovider.Deliverer{
			ChannelID:     chainID,
			Gossip:        d.conf.Gossip,
			Ledger:        ledgerInfo,
			BlockVerifier: d.conf.CryptoSvc,
			Dialer: DialerAdapter{
				ClientConfig: comm.ClientConfig{
					DialTimeout: d.conf.DeliverServiceConfig.ConnectionTimeout,
					KaOpts:      d.conf.DeliverServiceConfig.KeepaliveOptions,
					SecOpts:     d.conf.DeliverServiceConfig.SecOpts,
				},
			},
			Orderers:            d.conf.OrdererSource,
			DoneC:               make(chan struct{}),
			Signer:              d.conf.Signer,
			DeliverStreamer:     DeliverAdapter{},
			Logger:              flogging.MustGetLogger("peer.blocksprovider").With("channel", chainID),
			MaxRetryDelay:       d.conf.DeliverServiceConfig.ReConnectBackoffThreshold,
			MaxRetryDuration:    d.conf.DeliverServiceConfig.ReconnectTotalTimeThreshold,
			BlockGossipDisabled: !d.conf.DeliverServiceConfig.BlockGossipEnabled,
			InitialRetryDelay:   100 * time.Millisecond,
			YieldLeadership:     !d.conf.IsStaticLeader,
		}	
		if d.conf.DeliverGRPCClient.MutualTLSRequired() {
			dc.TLSCertHash = util.ComputeSHA256(d.conf.DeliverGRPCClient.Certificate().Certificate[0])
		}
		bp = dc
	}

	if dc.BlockGossipDisabled {
		logger.Infow("This peer will retrieve blocks from ordering service (will not disseminate them to other peers in the organization)", "channel", chainID)
	} else {
		logger.Infow("This peer will retrieve blocks from ordering service and disseminate to other peers in the organization", "channel", chainID)
	}

	if d.conf.DeliverServiceConfig.SecOpts.RequireClientCert {
		cert, err := d.conf.DeliverServiceConfig.SecOpts.ClientCertificate()
		if err != nil {
			return fmt.Errorf("failed to access client TLS configuration: %w", err)
		}
		dc.TLSCertHash = util.ComputeSHA256(cert.Certificate[0])
	}

	d.blockProviders[chainID] = dc
	go func() {
		dc.DeliverBlocks()
		finalizer()
	}()
	return nil
}

// StopDeliverForChannel stops blocks delivery for channel by stopping channel block provider
func (d *deliverServiceImpl) StopDeliverForChannel(chainID string) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.stopping {
		errMsg := fmt.Sprintf("Delivery service is stopping, cannot stop delivery for channel %s", chainID)
		logger.Errorf(errMsg)
		return errors.New(errMsg)
	}
	client, exist := d.blockProviders[chainID]
	if !exist {
		errMsg := fmt.Sprintf("Delivery service - no block provider for %s found, can't stop delivery", chainID)
		logger.Errorf(errMsg)
		return errors.New(errMsg)
	}
	client.Stop()
	delete(d.blockProviders, chainID)
	logger.Debug("This peer will stop pass blocks from orderer service to other peers")
	return nil
}

// Stop all service and release resources
func (d *deliverServiceImpl) Stop() {
	d.lock.Lock()
	defer d.lock.Unlock()
	// Marking flag to indicate the shutdown of the delivery service
	d.stopping = true

	for _, client := range d.blockProviders {
		client.Stop()
	}
}

// UpdateEndpoints assigns the new endpoints for the block provider
func (d *deliverServiceImpl) UpdateEndpoints(chainID string, endpoints []*orderers.Endpoint) error {
	d.lock.RLock()
	defer d.lock.RUnlock()

	// Use chainID to obtain blocks provider and pass endpoints
	// for update
	if dc, ok := d.blockProviders[chainID]; ok {
		// We have found specified channel so we can safely update it
		if bp, ok := dc.(interface {
			UpdateEndpoints(endpoints []*orderers.Endpoint)
		}); ok {
			logger.Infof("UpdateEndpoints for %s", chainID)
			bp.UpdateEndpoints(endpoints)
		} else {
			logger.Infof("No UpdateEndpoints for %s", chainID)
		}
		return nil
	}
	return errors.New(fmt.Sprintf("Channel with %s id was not found", chainID))
}
