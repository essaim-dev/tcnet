package tcnet

// Opt-in message data.
//
// In order to correctly implement the Opt-In usage, the following steps are needed.
//
// Step 1:
// Create a Opt-IN packet and make sure your Node Listeners Port value is correct.
//
// Step 2:
// Broadcast every 1000ms a Opt-IN packet to port 60000
//
// Step 3:
// Unicast every 1000ms a Opt-IN packet to each discovered node, targeting that node’s port. This ensures that when a node doesn’t receive broadcast messages, it still can discover your node.
type OptInData struct {
	// Number of nodes registered by system.
	NodeCount uint16

	// Listener port of node (Used to receive unicast messages).
	NodeListenerPort uint16

	// Up time of Node in seconds. (!) Must Roll over / Reset every 12 hours.
	Uptime uint16

	// Name of Vendor of Node.
	VendorName string

	// Name of Application/Device (Node).
	NodeName string

	// Major Version of Node.
	NodeVersionMajor uint8

	// Minor Version of Node.
	NodeVersionMinor uint8

	// Bug Version of Node.
	NodeVersionBug uint8
}

// Opt-out message data.
//
// TIP:
// In case of a disconnect of a Master Node in the network, the next master is chosen by looking at all Nodes running as Node Type 1 (Auto Master).
// The node that has the highest Uptime including Timestamp becomes the new master. This node changes its type to 2 (Master) and starts its services as such.
type OptOutData struct {
	// Number of nodes registered by system.
	NodeCount uint16

	// Listener port of node (Used to receive unicast messages).
	NodeListenerPort uint16
}
