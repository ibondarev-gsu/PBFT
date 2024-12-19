package message

type RequestMessage struct {
	Timestamp  int64  `json:"timestamp"`
	ClinetID   string `json:"clientID"`
	Operation  string `json:"operation"`
	SequenceID int64  `json:"sequenceID"`
}

type PrePrepareMessage struct {
	ViewID         int64           `json:"viewID"`
	SequenceID     int64           `json:"sequenceID"`
	Digest         string          `json:"digest"`
	RequestMessage *RequestMessage `json:"requestMessage"`
}

type VoteMessage struct {
	ViewID      int64  `json:"viewID"`
	SequenceID  int64  `json:"sequenceID"`
	Digest      string `json:"digest"`
	NodeID      string `json:"nodeID"`
	MessageType `json:"messageType"`
}

type ReplyMessage struct {
	ViewID    int64  `json:"viewID"`
	Timestamp int64  `json:"timestamp"`
	ClientID  string `json:"clientID"`
	NodeID    string `json:"nodeID"`
	Result    string `json:"result"`
}

type MessageType int

const (
	PrePrepare     MessageType = iota
	PrepareMessage MessageType = iota
	CommitMessage
)

type PBFT interface {
	StartConsensus(request *RequestMessage) (*PrePrepareMessage, error)
	PrePrepare(prePrepareMsg *PrePrepareMessage) (*VoteMessage, error)
	Prepare(prepareMsg *VoteMessage) (*VoteMessage, error)
	Commit(commitMsg *VoteMessage) (*ReplyMessage, *RequestMessage, error)
}

type State struct {
	ViewID         int64
	MsgLogs        *MsgLogs
	LastSequenceID int64
	CurrentStage   Stage
}
type MsgLogs struct {
	ReqMsg      *RequestMessage
	PrepareMsgs map[string]*VoteMessage
	CommitMsgs  map[string]*VoteMessage
}

type Stage int
