package api_gateway

import (
	"sync"

	"errors"

	"github.com/it-chain/engine/p2p"
)

var ErrPeerExists = errors.New("peer already exists")

type PeerQueryApi struct {
	mux            sync.Mutex
	peerRepository PeerRepository
}

func (pqa *PeerQueryApi) GetPLTable() (p2p.PLTable, error) {

	return pqa.peerRepository.GetPLTable()
}

func (pqa *PeerQueryApi) GetLeader() (p2p.Leader, error) {

	return pqa.peerRepository.GetLeader()
}

func (pqa *PeerQueryApi) FindPeerById(peerId p2p.PeerId) (p2p.Peer, error) {

	return pqa.peerRepository.FindPeerById(peerId)
}

func (pqa *PeerQueryApi) FindPeerByAddress(ipAddress string) (p2p.Peer, error) {

	return pqa.peerRepository.FindPeerByAddress(ipAddress)
}

type PeerRepository struct {
	mux     sync.Mutex
	pLTable p2p.PLTable
}

func (pltrepo *PeerRepository) GetPLTable() (p2p.PLTable, error) {

	pltrepo.mux.Lock()
	defer pltrepo.mux.Unlock()

	return pltrepo.pLTable, nil
}

func (pltrepo *PeerRepository) GetLeader() (p2p.Leader, error) {

	return pltrepo.pLTable.Leader, nil
}

func (pltrepo *PeerRepository) FindPeerById(peerId p2p.PeerId) (p2p.Peer, error) {

	pltrepo.mux.Lock()
	defer pltrepo.mux.Unlock()
	v, exist := pltrepo.pLTable.PeerTable[peerId.Id]

	if peerId.Id == "" {
		return v, p2p.ErrEmptyPeerId
	}
	//no matching id
	if !exist {
		return v, p2p.ErrNoMatchingPeerId
	}

	return v, nil
}

func (pltrepo *PeerRepository) FindPeerByAddress(ipAddress string) (p2p.Peer, error) {

	pltrepo.mux.Lock()
	defer pltrepo.mux.Unlock()

	for _, peer := range pltrepo.pLTable.PeerTable {

		if peer.IpAddress == ipAddress {
			return peer, nil
		}
	}

	return p2p.Peer{}, nil
}

func (pltrepo *PeerRepository) Save(peer p2p.Peer) error {

	pltrepo.mux.Lock()
	defer pltrepo.mux.Unlock()

	pLTable, _ := pltrepo.GetPLTable()

	_, exist := pLTable.PeerTable[peer.PeerId.Id]

	if exist {
		return ErrPeerExists
	}

	pLTable.PeerTable[peer.PeerId.Id] = peer

	return nil
}

func (pltrepo *PeerRepository) SetLeader(peer p2p.Peer) error {

	pltrepo.mux.Lock()
	defer pltrepo.mux.Unlock()

	leader := p2p.Leader{
		LeaderId: p2p.LeaderId{
			Id: peer.PeerId.Id,
		},
	}

	pltrepo.pLTable.Leader = leader

	return nil
}

func (pltrepo *PeerRepository) Delete(id string) error {

	pltrepo.mux.Lock()
	defer pltrepo.mux.Unlock()

	delete(pltrepo.pLTable.PeerTable, id)

	return nil
}

type P2PEventHandler struct {
	peerRepository PeerRepository
}

func (peh *P2PEventHandler) PeerCreatedEventHandler(event p2p.PeerCreatedEvent) error {

	peer := p2p.Peer{
		PeerId: p2p.PeerId{
			Id: event.ID,
		},
		IpAddress: event.IpAddress,
	}

	peh.peerRepository.Save(peer)

	return nil
}

func (peh *P2PEventHandler) PeerDeletedEventHandler(event p2p.PeerCreatedEvent) error {

	peh.peerRepository.Delete(event.ID)

	return nil
}

func (peh *P2PEventHandler) HandleLeaderUpdatedEvent(event p2p.LeaderUpdatedEvent) error {

	peer := p2p.Peer{
		PeerId: p2p.PeerId{Id: event.ID},
	}

	peh.peerRepository.SetLeader(peer)

	return nil

}
