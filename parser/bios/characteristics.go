package bios

// Characteristics bios字符集
type Characteristics uint64

// BIOS Characteristics
const (
	Reserved0 Characteristics = 1 << iota
	Reserved1
	Unknown
	NotSupported
	ISASupported
	MCASupported
	EISASupported
	PCISupported
	PCMCIASupported
	PlugPlaySupported
	APMSupported
	Upgradeable
	ShadowingIsAllowed
	VLVESASupported
	ESCDSupported
	BootFromCDSupported
	SelectableBootSupported
	BIOSROMIsSockectd
	BootFromPCMCIASupported
	EDDSupported
	JPFloppyNECSupported
	JPFloppyToshibaSupported
	Floppy525_360KBSupported
	Floppy525_1_2MBSupported
	Floppy35_720KBSupported
	Floppy35_2_88MBSupported
	PrintScreenSupported
	Keyboard8042Supported
	SerialSupported
	PrinterSupported
	CGAMonoSupported
	NECPC98
	//Bit32:47 Reserved for BIOS vendor
	//Bit47:63 Reserved for system vendor
)

func (b Characteristics) String() string {
	var s string
	chars := [...]string{
		"BIOS characteristics not supported", /* 3 */
		"ISA is supported",
		"MCA is supported",
		"EISA is supported",
		"PCI is supported",
		"PC Card (PCMCIA) is supported",
		"PNP is supported",
		"APM is supported",
		"BIOS is upgradeable",
		"BIOS shadowing is allowed",
		"VLB is supported",
		"ESCD support is available",
		"Boot from CD is supported",
		"Selectable boot is supported",
		"BIOS ROM is socketed",
		"Boot from PC Card (PCMCIA) is supported",
		"EDD is supported",
		"Japanese floppy for NEC 9800 1.2 MB is supported (int 13h)",
		"Japanese floppy for Toshiba 1.2 MB is supported (int 13h)",
		"5.25\"/360 kB floppy services are supported (int 13h)",
		"5.25\"/1.2 MB floppy services are supported (int 13h)",
		"3.5\"/720 kB floppy services are supported (int 13h)",
		"3.5\"/2.88 MB floppy services are supported (int 13h)",
		"Print screen service is supported (int 5h)",
		"8042 keyboard services are supported (int 9h)",
		"Serial services are supported (int 14h)",
		"Printer services are supported (int 17h)",
		"CGA/mono video services are supported (int 10h)",
		"NEC PC-98", /* 31 */
	}

	for i := uint32(4); i < 32; i++ {
		if b&(1<<i) != 0 {
			s += "\n\t\t" + chars[i-3]
		}
	}
	return s
}
