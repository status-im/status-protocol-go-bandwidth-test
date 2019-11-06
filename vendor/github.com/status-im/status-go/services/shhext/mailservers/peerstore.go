package mailservers

import (
	"errors"
	"sync"

	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	whispertypes "github.com/status-im/status-protocol-go/transport/whisper/types"
)

var (
	// ErrNoConnected returned when mail servers are not connected.
	ErrNoConnected = errors.New("no connected mail servers")
)

// PeersProvider is an interface for requesting list of peers.
type PeersProvider interface {
	Peers() []*p2p.Peer
}

// NewPeerStore returns an instance of PeerStore.
func NewPeerStore(cache *Cache) *PeerStore {
	return &PeerStore{
		nodes: map[whispertypes.EnodeID]*enode.Node{},
		cache: cache,
	}
}

// PeerStore stores list of selected mail servers and keeps N of them connected.
type PeerStore struct {
	mu    sync.RWMutex
	nodes map[whispertypes.EnodeID]*enode.Node

	cache *Cache
}

// Exist confirms that peers was added to a store.
func (ps *PeerStore) Exist(nodeID whispertypes.EnodeID) bool {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	_, exist := ps.nodes[nodeID]
	return exist
}

// Get returns instance of the node with requested ID or nil if ID is not found.
func (ps *PeerStore) Get(nodeID whispertypes.EnodeID) *enode.Node {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	return ps.nodes[nodeID]
}

// Update updates peers locally.
func (ps *PeerStore) Update(nodes []*enode.Node) error {
	ps.mu.Lock()
	ps.nodes = map[whispertypes.EnodeID]*enode.Node{}
	for _, n := range nodes {
		ps.nodes[whispertypes.EnodeID(n.ID())] = n
	}
	ps.mu.Unlock()
	return ps.cache.Replace(nodes)
}
