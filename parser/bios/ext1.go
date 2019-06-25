package bios

// Ext1 ext1
type Ext1 byte

// BIOS Characteristics Extension Bytes(Ext1)
// Byte 1
const (
	Ext1ACPISupported Ext1 = 1 << iota
	Ext1USBLegacySupported
	Ext1AGPSupported
	Ext1I2OBootSupported
	Ext1LS120SupperDiskBootSupported
	Ext1ATAPIZIPDriveBootSupported
	Ext11394BootSupported
	Ext1SmartBatterySupported
)

func (b Ext1) String() string {
	var s string
	chars := [...]string{
		"ACPI is supported", /* 0 */
		"USB legacy is supported",
		"AGP is supported",
		"I2O boot is supported",
		"LS-120 boot is supported",
		"ATAPI Zip drive boot is supported",
		"IEEE 1394 boot is supported",
		"Smart battery is supported", /* 7 */
	}

	for i := uint32(0); i < 7; i++ {
		if b&(1<<i) != 0 {
			s += "\n\t\t" + chars[i]
		}
	}
	return s
}
