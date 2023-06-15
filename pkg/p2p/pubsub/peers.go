package pubsub

import (
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/libp2p/go-libp2p/core/peer"
)

type info struct {
	// start time of the peer
	start time.Time
	// version information of the peer
	version string
	// address of the peer
	address common.Address
	// stake amount held by the peer
	stake *big.Int
	// protect info
	sync.RWMutex
}

func (i *info) getStart() time.Time {
	i.RLock()
	defer i.RUnlock()

	return i.start
}

func (i *info) getVersion() string {
	i.RLock()
	defer i.RUnlock()

	return i.version
}

func (i *info) getAddress() common.Address {
	i.RLock()
	defer i.RUnlock()

	return i.address
}

func (i *info) getStake() *big.Int {
	i.RLock()
	defer i.RUnlock()

	return i.stake
}

func (i *info) setStart(start time.Time) {
	i.Lock()
	defer i.Unlock()

	i.start = start
}

func (i *info) setVersion(version []byte) {
	i.Lock()
	defer i.Unlock()

	i.version = string(version)
}

func (i *info) setAddress(address common.Address) {
	i.Lock()
	defer i.Unlock()

	i.address = address
}

func (i *info) setStake(stake *big.Int) {
	i.Lock()
	defer i.Unlock()

	i.stake = stake
}

type approvedPeersMap struct {
	peers map[peer.ID]*info
	sync.RWMutex
}

func newApprovedPeersMap() *approvedPeersMap {
	return &approvedPeersMap{
		peers: make(map[peer.ID]*info, 1024),
	}
}

func (a *approvedPeersMap) AddPeer(peer peer.ID) {
	a.Lock()
	defer a.Unlock()
	a.peers[peer] = &info{}
}

func (a *approvedPeersMap) DelPeer(peer peer.ID) {
	a.Lock()
	defer a.Unlock()
	delete(a.peers, peer)
}

func (a *approvedPeersMap) InPeers(peer peer.ID) bool {
	a.RLock()
	defer a.RUnlock()

	_, ok := a.peers[peer]
	return ok
}

func (a *approvedPeersMap) GetPeers() map[peer.ID]*info {
	a.RLock()
	defer a.RUnlock()

	var peers = make(map[peer.ID]*info)
	for k, v := range a.peers {
		infoCopy := &info{
			start:   v.getStart(),
			version: v.getVersion(),
			address: v.getAddress(),
			stake:   v.getStake(),
		}

		peers[k] = infoCopy
	}

	return peers
}

func (a *approvedPeersMap) ListApprovedPeers() []peer.ID {
	a.RLock()
	defer a.RUnlock()

	approvedPeers := []peer.ID{}
	for k, _ := range a.peers {
		approvedPeers = append(approvedPeers, k)
	}

	return approvedPeers
}

// (start) set peer info options
func (a *approvedPeersMap) SetPeerInfoStart(peer peer.ID, start time.Time) {
	a.Lock()
	defer a.Unlock()
	if val, ok := a.peers[peer]; ok {
		val.setStart(start)
	}
}

// (version) set peer info options
func (a *approvedPeersMap) SetPeerInfoVersion(peer peer.ID, version []byte) {
	a.Lock()
	defer a.Unlock()
	if val, ok := a.peers[peer]; ok {
		val.setVersion(version)
	}
}

// (address) set peer info options
func (a *approvedPeersMap) SetPeerInfoAddress(peer peer.ID, address common.Address) {
	a.Lock()
	defer a.Unlock()
	if val, ok := a.peers[peer]; ok {
		val.setAddress(address)
	}
}

// (stake) set peer info options
func (a *approvedPeersMap) SetPeerInfoStake(peer peer.ID, stake *big.Int) {
	a.Lock()
	defer a.Unlock()
	if val, ok := a.peers[peer]; ok {
		val.setStake(stake)
	}
}
