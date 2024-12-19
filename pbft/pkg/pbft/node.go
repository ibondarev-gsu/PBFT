package pbft

import "pbft/pkg/message"

type View struct {
	ID      int64
	Primary string
}

type MessageBuffer struct {
	ReqMsgs        []*message.RequestMessage
	PrePrepareMsgs []*message.PrePrepareMessage
	PrepareMsgs    []*message.VoteMessage
	CommitMsgs     []*message.VoteMessage
}

type Node struct {
	NodeID       string
	NodeTable    map[string]string
	View         *View
	CurrentState *message.State
	CommitMsgs   []*message.RequestMessage
	MsgBuffer    *MessageBuffer
	MsgEntrance  chan interface{}
	MsgDelivery  chan interface{}
	Alarm        chan bool
}
