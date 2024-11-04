package tcnet

type NodeID uint16

type NodeType uint8

const (
	NodeTypeAuto NodeType = 1 << iota
	NodeTypeMaster
	NodeTypeSlave
	NodeTypeRepeater
)

type NodeOption uint16

const (
	// NodeOptionNeedsAuthentication signifies that authentication for extended communication with the node is needed.
	NodeOptionNeedsAuthentication NodeOption = 1 << iota
	// NodeOptionSupportsTCNCM signifies that the node listens to TCNet Control Messages.
	NodeOptionSupportsTCNCM
	// NodeOptionSupportsTCNASDP signifies that the node listens to TCNet Application Specific Data Packet.
	NodeOptionSupportsTCNASDP
	// NodeOptionDND signifies that the node is in do not disturb/sleeping mode. The node will request data itself if needed to avoid traffic.
	NodeOptionDND
)
