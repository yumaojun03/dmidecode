package smbios

// StructureType Structure definition type
type StructureType byte

const (
	BIOS StructureType = iota
	System
	BaseBoard
	Chassis
	Processor
	Controller
	Module
	Cache
	PortConnector
	SystemSlots
	OnBoardDevices
	OEMStrings
	SystemConfigurationOptions
	BIOSLanguage
	GroupAssociations
	SystemEventLog
	PhysicalMemoryArray
	MemoryDevice
	Bit32MemoryError
	MemoryArrayMappedAddress
	MemoryDeviceMappedAddress
	BuiltInPointingDevice
	PortableBattery
	SystemReset
	HardwareSecurity
	SystemPowerControls
	VoltageProbe
	CoolingDevice
	TemperatureProbe
	ElectricalCurrentProbe
	OutOfBandRemoteAccess
	BootIntegrityServices
	SystemBoot
	Bit64MemoryError
	ManagementDevice
	ManagementDeviceComponent
	ManagementDeviceThresholdData
	MemoryChannel
	IPMIDevice
	PowerSupply
	AdditionalInformation
	OnBoardDevicesExtendedInformation
	ManagementControllerHostInterface               /*42*/
	Inactive                          StructureType = 126
	EndOfTable                        StructureType = 127
)

func (b StructureType) String() string {
	types := [...]string{
		"BIOS", /* 0 */
		"System",
		"Base Board",
		"Chassis",
		"Processor",
		"Memory Controller",
		"Memory Module",
		"Cache",
		"Port Connector",
		"System Slots",
		"On Board Devices",
		"OEM Strings",
		"System Configuration Options",
		"BIOS Language",
		"Group Associations",
		"System Event Log",
		"Physical Memory Array",
		"Memory Device",
		"32-bit Memory Error",
		"Memory Array Mapped Address",
		"Memory Device Mapped Address",
		"Built-in Pointing Device",
		"Portable Battery",
		"System Reset",
		"Hardware Security",
		"System Power Controls",
		"Voltage Probe",
		"Cooling Device",
		"Temperature Probe",
		"Electrical Current Probe",
		"Out-of-band Remote Access",
		"Boot Integrity Services",
		"System Boot",
		"64-bit Memory Error",
		"Management Device",
		"Management Device Component",
		"Management Device Threshold Data",
		"Memory Channel",
		"IPMI Device",
		"Power Supply",
		"Additional Information",
		"Onboard Device",
		"Management Controller Host Interface", /* 42 */
	}
	return types[b]
}
