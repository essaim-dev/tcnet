package tcnet

import (
	"essaim.dev/tcnet/flame"
)

type MessageType uint8

const (
	MessageTypeOptIn  MessageType = 002
	MessageTypeOptOut MessageType = 003
	MessageTypeStatus MessageType = 005
)

type Header struct {
	NodeID      NodeID
	Version     flame.Version
	MessageType MessageType
	NodeName    string
	Sequence    uint8
	NodeType    NodeType
	Timestamp   uint32
}
