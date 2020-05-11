package slot

type Type byte

const (
	TypeOther Type = 1 + iota
	TypeUnknown
	TypeISA
	TypeMCA
	TypeEISA
	TypePCI
	TypePCCardPCMCIA
	TypeVL_VESA
	TypeProprietary
	TypeProcessorCardSlot
	TypeProprietaryMemoryCardSlot
	TypeIORiserCardSlot
	TypeNuBus
	TypePCI_66MHzCapable
	TypeAGP
	TypeAGP2X
	TypeAGP4X
	TypePCI_X
	TypeAGP8X
	TYPEM2Socket1DP
	TYPEM2Socket1SP
	TYPEM2Socket2
	TYPEM2Socket3
	TYPEMXMTypeI
	TYPEMXMTypeII
	TYPEMXMTypeIIISTD
	TYPEMXMTypeIIIHE
	TYPEMXMTypeIV
	TYPEMXM30TypeA
	TYPEMXM30TypeB
	TYPEPCIExpressGen2SFF8639
	TYPEPCIExpressGen3SFF8639
)

const (
	TypePC_98C20 = 160 + iota
	TypePC_98C24
	TypePC_98E
	TypePC_98LocalBus
	TypePC_98Card
	TypePCIExpress
	TypePCIExpressx1
	TypePCIExpressx2
	TypePCIExpressx4
	TypePCIExpressx8
	TypePCIExpressx16
	TypePCIExpressGen2
	TypePCIExpressGen2x1
	TypePCIExpressGen2x2
	TypePCIExpressGen2x4
	TypePCIExpressGen2x8
	TypePCIExpressGen2x16
	TypePCIExpressGen3
	TypePCIExpressGen3x1
	TypePCIExpressGen3x2
	TypePCIExpressGen3x4
	TypePCIExpressGen3x8
	TypePCIExpressGen3x16
)

func (s Type) String() string {
	types := [...]string{
		"Other",
		"Unknown",
		"ISA",
		"MCA",
		"EISA",
		"PCI",
		"PC Card (PCMCIA)",
		"VL-VESA",
		"Proprietary",
		"Processor Card Slot",
		"Proprietary Memory Card Slot",
		"I/O Riser Card Slot",
		"NuBus",
		"PCI – 66MHz Capable",
		"AGP",
		"AGP 2X",
		"AGP 4X",
		"PCI-X",
		"AGP 8X",
		"M.2 Socket 1-DP (Mechanical Key A)",
		"M.2 Socket 1-SD (Mechanical Key E)",
		"M.2 Socket 2 (Mechanical Key B)",
		"M.2 Socket 3 (Mechanical Key M)",
		"MXM Type I",
		"MXM Type II",
		"MXM Type III (standard connector)",
		"MXM Type III (HE connector)",
		"MXM Type IV",
		"MXM 3.0 Type A",
		"MXM 3.0 Type B",
		"PCI Express Gen 2 SFF-8639",
		"PCI Express Gen 3 SFF-8639",
	}

	typesExt := [...]string{
		"PC-98/C20",
		"PC-98/C24",
		"PC-98/E",
		"PC-98/Local Bus",
		"PC-98/Card",
		"PCI Express",
		"PCI Express x1",
		"PCI Express x2",
		"PCI Express x4",
		"PCI Express x8",
		"PCI Express x16",
		"PCI Express Gen 2",
		"PCI Express Gen 2 x1",
		"PCI Express Gen 2 x2",
		"PCI Express Gen 2 x4",
		"PCI Express Gen 2 x8",
		"PCI Express Gen 2 x16",
		"PCI Express Gen 3",
		"PCI Express Gen 3 x1",
		"PCI Express Gen 3 x2",
		"PCI Express Gen 3 x4",
		"PCI Express Gen 3 x8",
		"PCI Express Gen 3 x16",
	}

	if int(s) >= 160 {
		return typesExt[s-160]
	}

	return types[s-1]

}

type DataBusWidth byte

const (
	DataBusWidthOther DataBusWidth = 1 + iota
	DataBusWidthUnknown
	DataBusWidth8bit
	DataBusWidth16bit
	DataBusWidth32bit
	DataBusWidth64bit
	DataBusWidth128bit
	DataBusWidth1xorx1
	DataBusWidth2xorx2
	DataBusWidth4xorx4
	DataBusWidth8xorx8
	DataBusWidth12xorx12
	DataBusWidth16xorx16
	DataBusWidth32xorx32
)

func (s DataBusWidth) String() string {
	widths := [...]string{
		"Other",
		"Unknown",
		"8 bit",
		"16 bit",
		"32 bit",
		"64 bit",
		"128 bit",
		"1x or x1",
		"2x or x2",
		"4x or x4",
		"8x or x8",
		"12x or x12",
		"16x or x16",
		"32x or x32",
	}
	return widths[s-1]
}

type Usage byte

const (
	UsageOther Usage = 1 + iota
	UsageUnknown
	UsageAvailable
	UsageInuse
)

func (s Usage) String() string {
	usages := [...]string{
		"Other",
		"Unknown",
		"Available",
		"In use",
	}
	return usages[s-1]
}

type Length byte

const (
	LengthOther Length = 1 + iota
	LengthUnknown
	LengthShortLength
	LengthLongLength
)

func (s Length) String() string {
	lengths := [...]string{
		"Other",
		"Unknown",
		"Short Length",
		"Long Length",
	}
	return lengths[s-1]
}

type ID uint16

type Characteristics1 byte

const (
	Characteristicsunknown Characteristics1 = 1 << iota
	CharacteristicsProvides5_0volts
	CharacteristicsProvides3_3volts
	CharacteristicsSlotsopeningissharedwithanotherslot
	CharacteristicsPCCardslotsupportsPCCard_16
	CharacteristicsPCCardslotsupportsCardBus
	CharacteristicsPCCardslotsupportsZoomVideo
	CharacteristicsPCCardslotsupportsModemRingResume
)

func (s Characteristics1) String() string {
	chars := [...]string{
		"Characteristics unknown.",
		"Provides 5.0 volts.",
		"Provides 3.3 volts.",
		"Slot’s opening is shared with another slot (for example, PCI/EISA shared slot).",
		"PC Card slot supports PC Card-16.",
		"PC Card slot supports CardBus.",
		"PC Card slot supports Zoom Video.",
		"PC Card slot supports Modem Ring Resume.",
	}
	return chars[s>>1]
}

type Characteristics2 byte

const (
	Characteristics2PCIslotsupportsPowerManagementEventsignal Characteristics2 = 1 << iota
	Characteristics2Slotsupportshot_plugdevices
	Characteristics2PCIslotsupportsSMBussignal
	Characteristics2Reserved
)

func (s Characteristics2) String() string {
	chars := [...]string{
		"PCI slot supports Power Management Event (PME#) signal.",
		"Slot supports hot-plug devices.",
		"PCI slot supports SMBus signal.",
		"Reserved",
	}
	return chars[s>>1]
}

type SegmengGroupNumber uint16

type Number byte
