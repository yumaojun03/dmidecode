package chassis

import "fmt"

const (
	OUT_OF_SPEC = "<OUT OF SPEC>"
)

// Type Chssis 类型
type Type byte

const (
	Other Type = 1 + iota
	Unknown
	Desktop
	LowProfileDesktop
	PizzaBox
	MiniTower
	Tower
	Portable
	Laptop
	Notebook
	HandHeld
	DockingStation
	AllinOne
	SubNotebook
	SpaceSaving
	LunchBox
	MainServerChassis
	ExpansionChassis
	SubChassis
	BusExpansionChassis
	PeripheralChassis
	RAIDChassis
	RackMountChassis
	SealedcasePC
	MultiSystem
	CompactPCI
	AdvancedTCA
	Blade
	BladeEnclosure
)

func (c Type) String() string {
	types := [...]string{
		"Other",
		"Unknown",
		"Desktop",
		"LowProfileDesktop",
		"PizzaBox",
		"MiniTower",
		"Tower",
		"Portable",
		"Laptop",
		"Notebook",
		"HandHeld",
		"Docking Station",
		"AllinOne",
		"SubNotebook",
		"SpaceSaving",
		"LunchBox",
		"MainServer Chassis",
		"Expansion Chassis",
		"Sub Chassis",
		"BusExpansion Chassis",
		"Peripheral Chassis",
		"RAID Chassis",
		"Rack Mount Chassis",
		"SealedcasePC",
		"MultiSystem",
		"CompactPCI",
		"AdvancedTCA",
		"Blade",
		"BladeEnclosure",
	}
	c &= 0x7F
	if c >= 0x01 && c < 0x1D {
		return types[c-1]
	}
	return OUT_OF_SPEC
}

// Lock todo
type Lock byte

func (c Lock) String() string {
	locks := [...]string{
		"Not Present", /* 0x00 */
		"Present",     /* 0x01 */
	}
	return locks[c]
}

// State todo
type State byte

const (
	StateOther State = 1 + iota
	StateUnknown
	StateSafe
	StateWarning
	StateCritical
	StateNonRecoverable
)

func (c State) String() string {
	states := [...]string{
		"Other",
		"Unknown",
		"Safe",
		"Warning",
		"Critical",
		"NonRecoverable",
	}
	return states[c-1]
}

type ContainedElementType byte

type ContainedElements struct {
	Type    ContainedElementType
	Minimum byte
	Maximum byte
}

type SecurityStatus byte

const (
	ChassisSecurityStatusOther SecurityStatus = 1 + iota
	ChassisSecurityStatusUnknown
	ChassisSecurityStatusNone
	ChassisSecurityStatusExternalInterfaceLockedOut
	ChassisSecurityStatusExternalInterfaceEnabled
)

func (s SecurityStatus) String() string {
	status := [...]string{
		"Other",
		"Unknown",
		"None",
		"ExternalInterfaceLockedOut",
		"ExternalInterfaceEnabled",
	}
	return status[s-1]
}

// Height 高度
type Height byte

func (s Height) String() string {
	if s == 0 {
		return fmt.Sprintf("Not Specified")
	}

	return fmt.Sprintf("%d U", s)
}
