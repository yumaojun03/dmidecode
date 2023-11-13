package tpm

import "fmt"

type 	SpecificationVersion struct {
	Major byte 
	Minor byte 
} 

func (v SpecificationVersion) String() string {
	return fmt.Sprintf("%d.%d", v.Major, v.Minor)
}

type 	FirmwareRevision struct {
	Major uint32 
	Minor uint32 
} 

func (v 	FirmwareRevision) String() string {
	return fmt.Sprintf("%d.%d", v.Major, v.Minor)
}

type TpmDeviceCharacteristics struct {
	Raw uint64
}

func (t TpmDeviceCharacteristics) Reserved0() bool {
	return t.Raw&0x0000000000000001 == 0x0000000000000001
}

func (t TpmDeviceCharacteristics) Reserved1() bool {
	return t.Raw&0x0000000000000002 == 0x0000000000000002
}

func (t TpmDeviceCharacteristics) NotSupported() bool {
	return t.Raw&0x0000000000000004 == 0x0000000000000004
}

func (t TpmDeviceCharacteristics) FamilyConfigurableViaFirmware() bool {
	return t.Raw&0x0000000000000008 == 0x0000000000000008
}

func (t TpmDeviceCharacteristics) FamilyConfigurableViaSoftware() bool {
	return t.Raw&0x0000000000000010 == 0x0000000000000010
}

func (t TpmDeviceCharacteristics) FamilyConfigurableViaOem() bool {
	return t.Raw&0x0000000000000020 == 0x0000000000000020
}

func (t TpmDeviceCharacteristics) String() string {
	characteristics := ""
	if t.Reserved0() {
		characteristics += fmt.Sprintf("\t\tReserved0\n")
	}
	if t.Reserved1() {
		characteristics +=  fmt.Sprintf("\t\tReserved1\n")
	}
	if t.NotSupported() {
		characteristics += fmt.Sprintf("\t\tTPM Device characteristics not supported\n")
	}
	if t.FamilyConfigurableViaFirmware() {
		characteristics += fmt.Sprintf("\t\tFamily configurable via firmware update\n")
	}
	if t.FamilyConfigurableViaSoftware() {
		characteristics += fmt.Sprintf("\t\tFamily configurable via platform software support\n")
	}
	if  t.FamilyConfigurableViaOem() {
		characteristics += fmt.Sprintf("\t\tFamily configurable via OEM proprietary mechanism\n")
	}
	return characteristics
}

type OEMSpecificInformation struct {
	Oem uint32
}

func (o OEMSpecificInformation) String() string {
	return fmt.Sprintf("0x%08x", o.Oem)
}