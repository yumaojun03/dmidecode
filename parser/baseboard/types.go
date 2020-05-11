package baseboard

import "fmt"

// FeatureFlags 主板功能标签
type FeatureFlags byte

// Baseboard feature flags
const (
	HostingBoard FeatureFlags = 1 << iota
	AtLeastOneDaughter
	Removable
	Repleaceable
	HotSwappable
	//FeatureFlagsReserved = 000b
)

func (f FeatureFlags) String() string {
	fmt.Printf("xxx,%d", f)
	features := [...]string{
		"Board is a hosting board", /* 0 */
		"Board requires at least one daughter board",
		"Board is removable",
		"Board is replaceable",
		"Board is hot swappable", /* 4 */
	}
	var s string
	for i := uint32(0); i < 5; i++ {
		if f&(1<<i) != 0 {
			s += "\n\t\t" + features[i]
		}
	}
	return s
}

// Type 主板类型
type Type byte

const (
	Unknown Type = 1 + iota
	Other
	ServerBlade
	ConnectivitySwitch
	ManagementModule
	ProcessorModule
	IOModule
	MemModule
	DaughterBoard
	Motherboard
	ProcessorMemmoryModule
	ProcessorIOModule
	InterconnectBoard
)

func (b Type) String() string {
	types := [...]string{
		"Unknown", /* 0x01 */
		"Other",
		"Server Blade",
		"Connectivity Switch",
		"System Management Module",
		"Processor Module",
		"I/O Module",
		"Memory Module",
		"Daughter Board",
		"Motherboard",
		"Processor+Memory Module",
		"Processor+I/O Module",
		"Interconnect Board", /* 0x0D */
	}
	if b > Unknown && b < InterconnectBoard {
		return types[b-1]
	}
	return "Out Of Spec"
}
