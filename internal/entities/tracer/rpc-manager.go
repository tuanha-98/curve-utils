package tracer

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

// RPCManager handles RPC endpoint rotation and connection management
type RPCManager struct {
	rpcs       []string
	currentIdx int
	mu         sync.Mutex
	clients    sync.Map // Cache of active clients
}

func NewRPCManager(rpcs []string) *RPCManager {
	return &RPCManager{
		rpcs:       rpcs,
		currentIdx: 0,
	}
}

func (rm *RPCManager) GetNextRPC() string {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	rpc := rm.rpcs[rm.currentIdx]
	rm.currentIdx = (rm.currentIdx + 1) % len(rm.rpcs)
	return rpc
}

func (rm *RPCManager) Dial(ctx context.Context) (*ethclient.Client, error) {
	// Check cached clients
	rm.clients.Range(func(key, value interface{}) bool {
		rpc, _ := key.(string)
		client, ok := value.(*ethclient.Client)
		if !ok {
			rm.clients.Delete(rpc)
			return true
		}
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		if _, err := client.ChainID(ctx); err == nil {
			return false // Stop iteration, client is alive
		}
		rm.clients.Delete(rpc)
		return true
	})

	// Try connecting to a new RPC
	for i := 0; i < len(rm.rpcs); i++ {
		rpc := rm.GetNextRPC()
		log.Printf("Attempting to connect to RPC %s", rpc)
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		client, err := ethclient.DialContext(ctx, rpc)
		if err == nil {
			rm.clients.Store(rpc, client)
			log.Printf("Connected to RPC %s", rpc)
			return client, nil
		}
		log.Printf("RPC connection attempt %d failed for %s: %v", i+1, rpc, err)
		time.Sleep(100 * time.Millisecond)
	}
	return nil, fmt.Errorf("failed to connect to any RPC endpoint after %d attempts", len(rm.rpcs))
}
