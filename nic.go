package virtualbox

// NIC represents a virtualized network interface card.
type NIC struct {
	Network       NICNetwork
	Hardware      NICHardware
	HostInterface string // The host interface name to bind to in 'hostonly' and 'bridged' mode
	NatNetwork string // The nat network to use in 'natnetwork' mode
	MacAddr       string
}

// NICNetwork represents the type of NIC networks.
type NICNetwork string

const (
	NICNetAbsent       = NICNetwork("none")
	NICNetDisconnected = NICNetwork("null")
	NICNetNAT          = NICNetwork("nat")
	NICNetNATNetwork   = NICNetwork("natnetwork")
	NICNetBridged      = NICNetwork("bridged")
	NICNetInternal     = NICNetwork("intnet")
	NICNetHostonly     = NICNetwork("hostonly")
	NICNetGeneric      = NICNetwork("generic")
)

// NICHardware represents the type of NIC hardware.
type NICHardware string

const (
	AMDPCNetPCIII         = NICHardware("Am79C970A")
	AMDPCNetFASTIII       = NICHardware("Am79C973")
	IntelPro1000MTDesktop = NICHardware("82540EM")
	IntelPro1000TServer   = NICHardware("82543GC")
	IntelPro1000MTServer  = NICHardware("82545EM")
	VirtIO                = NICHardware("virtio")
)
