package p2p

import "github.com/it-chain/midgard"

//publish
type LeaderChangedEvent struct {
	midgard.EventModel
}

//handle
type ConnectionCreatedEvent struct {
	midgard.EventModel
	Address string
}

//handle
type ConnectionDisconnectedEvent struct {
	midgard.EventModel
}

// node created event
type PeerCreatedEvent struct {
	midgard.EventModel
	IpAddress string
}

type PeerDeletedEvent struct {
	midgard.EventModel
}

// handle leader received event
type LeaderUpdatedEvent struct {
	midgard.EventModel
}

type LeaderDeliveredEvent struct {
	midgard.EventModel
}

//todo add to event doc
type LeaderDeletedEvent struct {
	midgard.EventModel
}
