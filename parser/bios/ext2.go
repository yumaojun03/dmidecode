package bios

// Ext2 ext2
type Ext2 byte

// BIOS Characteristics Extension Bytes(Ext2)
// Byte 2
const (
	Ext2BIOSBootSpecSupported Ext2 = 1 << iota
	Ext2FuncKeyInitiatedNetworkBootSupported
	Ext2EnableTargetedContentDistribution
	Ext2UEFISpecSupported
	Ext2VirtualMachine
	// Bits 5:7 Reserved for future assignment
)

func (b Ext2) String() string {
	var s string
	chars := [...]string{
		"BIOS boot specification is supported", /* 0 */
		"Function key-initiated network boot is supported",
		"Targeted content distribution is supported",
		"UEFI is supported",
		"System is a virtual machine", /* 4 */
	}

	for i := uint32(0); i < 5; i++ {
		if b&(1<<i) != 0 {
			s += "\n\t\t" + chars[i]
		}
	}
	return s
}
