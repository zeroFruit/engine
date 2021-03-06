package mock

import (
	"github.com/it-chain/engine/p2p"
)

type MockPeerService struct {
	SaveFunc          func(peer p2p.Peer) error
	RemoveFunc        func(peerId p2p.PeerId) error
	FindByIdFunc      func(peerId p2p.PeerId) (p2p.Peer, error)
	FindByAddressFunc func(ipAddress string) (p2p.Peer, error)
	FindAllFunc       func() ([]p2p.Peer, error)
}

type MockLeaderService struct {
	SetFunc func(leader p2p.Leader) error
}

func (ls *MockLeaderService) Set(leader p2p.Leader) error {

	return ls.SetFunc(leader)
}

type MockCommunicationService struct {
	DialFunc           func(ipAddress string) error
	DeliverPLTableFunc func(connectionId string, pLTable p2p.PLTable) error
}

func (mcs *MockCommunicationService) Dial(ipAddress string) error {

	return mcs.DialFunc(ipAddress)
}

func (mcs *MockCommunicationService) DeliverPLTable(connectionId string, pLTable p2p.PLTable) error {

	return mcs.DeliverPLTableFunc(connectionId, pLTable)
}

type MockPeerQueryService struct {
	GetPLTableFunc        func() (p2p.PLTable, error)
	GetLeaderFunc         func() (p2p.Leader, error)
	FindPeerByIdFunc      func(peerId p2p.PeerId) (p2p.Peer, error)
	FindPeerByAddressFunc func(ipAddress string) (p2p.Peer, error)
}

func (mpltqs *MockPeerQueryService) GetPLTable() (p2p.PLTable, error) {

	return mpltqs.GetPLTableFunc()
}

func (mpltqs *MockPeerQueryService) GetLeader() (p2p.Leader, error) {

	return mpltqs.GetLeaderFunc()
}

func (mpltqs *MockPeerQueryService) FindPeerById(peerId p2p.PeerId) (p2p.Peer, error) {

	return mpltqs.FindPeerByIdFunc(peerId)
}

func (mpltqs *MockPeerQueryService) FindPeerByAddress(ipAddress string) (p2p.Peer, error) {

	return mpltqs.FindPeerByAddressFunc(ipAddress)
}

type MockPLTableService struct {
	GetPLTableFromCommandFunc func(command p2p.GrpcReceiveCommand) (p2p.PLTable, error)
}

func (mplts *MockPLTableService) GetPLTableFromCommand(command p2p.GrpcReceiveCommand) (p2p.PLTable, error) {

	return mplts.GetPLTableFromCommandFunc(command)
}

type MockElectionService struct {
	VoteFunc             func(connectionId string) error
	BroadcastLeaderFunc  func(peer p2p.Peer) error
	DecideToBeLeaderFunc func(command p2p.GrpcReceiveCommand) error
}

func (mes *MockElectionService) Vote(connectionId string) error {

	return mes.VoteFunc(connectionId)

}
func (mes *MockElectionService) BroadcastLeader(peer p2p.Peer) error {

	return mes.BroadcastLeaderFunc(peer)

}
func (mes *MockElectionService) DecideToBeLeader(command p2p.GrpcReceiveCommand) error {

	return mes.DecideToBeLeaderFunc(command)

}
