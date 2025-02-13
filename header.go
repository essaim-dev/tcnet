package tcnet

import (
	"essaim.dev/tcnet/flame"
)

type MessageType uint8

const (
	// Opt-in/opt-out messages.
	//
	// Broadcast of Realtime Status messages.

	MessageTypeOptIn  MessageType = 002 // Present and keep alive a node into a TCNet network.
	MessageTypeOptOut MessageType = 003 // Notifies other nodes that node leaves network.

	// Status messages.

	MessageTypeStatus MessageType = 005 // Status packet of current settings on node.

	// Sync messages.

	MessageTypeTimeSync MessageType = 010 // Send and Receive Time Sync Data.

	// Notification messages.
	MessageTypeErrorNotification   MessageType = 013 // Notifies that a request is not handled.
	MessageTypeRequest             MessageType = 020 // Request Data from other Node.
	MessageTypeApplicationSpecific MessageType = 030

	// Control messages.
	//
	// Control messages are special messages that allow remote control TCNet nodes.

	MessageTypeControl      MessageType = 101 // Send and Receive Control Packets to control nodes remotely.
	MessageTypeTextData     MessageType = 128 // Send and Receive Text Data Packets to control nodes remotely.
	MessageTypeKeyboardData MessageType = 133 // Send and Receive Realtime Keyboard Data Packets to control nodes remotely.

	// Data packets.
	//
	// Data message types are messages containing data such as metadata, timing data, waveform data, cues etc.

	MessageTypeData MessageType = 200 // Updates Metrics Data for Layer.

	// File packets.
	//
	// File packet types are packets containing data such as images and audio files.

	MessageTypeFileData MessageType = 204 // Contains Low Res Artwork file (JPEG Format).

	// Application specific data packets.
	//
	// Application Specific Data packet types are packets containing data exchanged between applications.

	MessageTypeApplicationSpecifcData MessageType = 213 // Application Specific Broadcasted Data.

	// Timing Packets.
	//
	// Time Packets are time critical and updated at high rates.

	MessageTypeTime MessageType = 254 // Constant stream of timing data of layers.
)

type Header struct {
	// Unique Node ID. When multiple applications/services are running on same IP, this number must be unique.
	NodeID NodeID

	// Protocol version of source that sends the packet. - Example: 1.00 would be: 01 00
	Version flame.Version

	// TCNet Protocol Header (Must be “TCN”)
	Header string

	// Message type of packet.
	MessageType MessageType

	// GW Code of software/machine/source that sends packet. (8 Characters).
	NodeName string

	// Sequence number of packet.
	Sequence uint8

	// Node Type.
	NodeType NodeType

	// Node options. (Bit field).
	NodeOptions NodeOption

	// Time stamp in microseconds that is used to calculate network latency.
	Timestamp uint32
}
