package onboard

// ExtendedInformationType 主板上的设备类型
type ExtendedInformationType byte

const (
	ExtendedInformationTypeOther ExtendedInformationType = 1 + iota
	ExtendedInformationTypeUnknown
	ExtendedInformationTypeVideo
	ExtendedInformationTypeSCSIController
	ExtendedInformationTypeEthernet
	ExtendedInformationTypeTokenRing
	ExtendedInformationTypeSound
	ExtendedInformationTypePATAController
	ExtendedInformationTypeSATAController
	ExtendedInformationTypeSASController
)

func (o ExtendedInformationType) String() string {
	types := [...]string{
		"Other",
		"Unknown",
		"Video",
		"SCSI Controller",
		"Ethernet",
		"Token Ring",
		"Sound",
		"PATA Controller",
		"SATA Controller",
		"SAS Controller",
	}
	return types[o-1]
}

// DeviceStatus 设备状态
type DeviceStatus byte

const (
	DeviceDisabled DeviceStatus = 0
	DeviceEnabled               = 1
)

func (d DeviceStatus) String() string {
	types := [...]string{
		"Disabled",
		"Enabled",
	}
	return types[d]
}
