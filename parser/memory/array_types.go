package memory

type PhysicalMemoryArrayLocation byte

const (
	PhysicalMemoryArrayLocationOther PhysicalMemoryArrayLocation = 1 + iota
	PhysicalMemoryArrayLocationUnknown
	PhysicalMemoryArrayLocationSystemboardormotherboard
	PhysicalMemoryArrayLocationISAadd_oncard
	PhysicalMemoryArrayLocationEISAadd_oncard
	PhysicalMemoryArrayLocationPCIadd_oncard
	PhysicalMemoryArrayLocationMCAadd_oncard
	PhysicalMemoryArrayLocationPCMCIAadd_oncard
	PhysicalMemoryArrayLocationProprietaryadd_oncard
	PhysicalMemoryArrayLocationNuBus
	PhysicalMemoryArrayLocationPC_98C20add_oncard
	PhysicalMemoryArrayLocationPC_98C24add_oncard
	PhysicalMemoryArrayLocationPC_98Eadd_oncard
	PhysicalMemoryArrayLocationPC_98Localbusadd_oncard
)

func (p PhysicalMemoryArrayLocation) String() string {
	locations := [...]string{
		"Other",
		"Unknown",
		"System board or motherboard",
		"ISA add-on card",
		"EISA add-on card",
		"PCI add-on card",
		"MCA add-on card",
		"PCMCIA add-on card",
		"Proprietary add-on card",
		"NuBus",
		"PC-98/C20 add-on card",
		"PC-98/C24 add-on card",
		"PC-98/E add-on card",
		"PC-98/Local bus add-on card",
	}
	return locations[p-1]
}

type PhysicalMemoryArrayUse byte

const (
	PhysicalMemoryArrayUseOther PhysicalMemoryArrayUse = 1 + iota
	PhysicalMemoryArrayUseUnknown
	PhysicalMemoryArrayUseSystemmemory
	PhysicalMemoryArrayUseVideomemory
	PhysicalMemoryArrayUseFlashmemory
	PhysicalMemoryArrayUseNon_volatileRAM
	PhysicalMemoryArrayUseCachememory
)

func (p PhysicalMemoryArrayUse) String() string {
	uses := [...]string{
		"Other",
		"Unknown",
		"System memory",
		"Video memory",
		"Flash memory",
		"Non-volatile RAM",
		"Cache memory",
	}
	return uses[p-1]
}

type PhysicalMemoryArrayErrorCorrection byte

const (
	PhysicalMemoryArrayErrorCorrectionOther PhysicalMemoryArrayErrorCorrection = 1 + iota
	PhysicalMemoryArrayErrorCorrectionUnknown
	PhysicalMemoryArrayErrorCorrectionNone
	PhysicalMemoryArrayErrorCorrectionParity
	PhysicalMemoryArrayErrorCorrectionSingle_bitECC
	PhysicalMemoryArrayErrorCorrectionMulti_bitECC
	PhysicalMemoryArrayErrorCorrectionCRC
)

func (p PhysicalMemoryArrayErrorCorrection) String() string {
	types := [...]string{
		"Other",
		"Unknown",
		"None",
		"Parity",
		"Single-bit ECC",
		"Multi-bit ECC",
		"CRC",
	}
	return types[p-1]
}
