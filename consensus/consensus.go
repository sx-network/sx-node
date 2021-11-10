package consensus

import (
	"context"
	"log"

	"github.com/0xPolygon/polygon-sdk/blockchain"
	"github.com/0xPolygon/polygon-sdk/chain"
	"github.com/0xPolygon/polygon-sdk/network"
	"github.com/0xPolygon/polygon-sdk/state"
	"github.com/0xPolygon/polygon-sdk/txpool"
	"github.com/0xPolygon/polygon-sdk/types"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

// Consensus is the public interface for consensus mechanism
// Each consensus mechanism must implement this interface in order to be valid
type Consensus interface {
	// VerifyHeader verifies the header is correct
	VerifyHeader(parent, header *types.Header) error

	// GetBlockCreator retrieves the block creator (or signer) given the block header
	GetBlockCreator(header *types.Header) (types.Address, error)

	// Start starts the consensus
	Start() error

	// Close closes the connection
	Close() error

	// Returns whether or not ibft node is stale
	IsIbftStateStale() bool
}

// Config is the configuration for the consensus
type Config struct {
	// Logger to be used by the backend
	Logger *log.Logger

	// Params are the params of the chain and the consensus
	Params *chain.Params

	// Config defines specific configuration parameters for the backend
	Config map[string]interface{}

	// Path is the directory path for the consensus protocol tos tore information
	Path string
}

// Factory is the factory function to create a discovery backend
type Factory func(
	context.Context,
	bool, *Config,
	*txpool.TxPool,
	*network.Server,
	*blockchain.Blockchain,
	*state.Executor,
	*grpc.Server,
	hclog.Logger,
) (Consensus, error)
