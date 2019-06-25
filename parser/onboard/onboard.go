package onboard

import (
	"dmidecode/smbios"
	"fmt"
)

// ExtendedInformation 扩展设备信息
type ExtendedInformation struct {
	smbios.Header
	ReferenceDesignation string
	DeviceType           ExtendedInformationType
	DeviceStatus         DeviceStatus
	DeviceTypeInstance   byte
	SegmentGroupNumber   uint16
	BusNumber            byte
	DeviceFunctionNumber byte
}

// SlotSegment 主板设备对应的slot
func (o ExtendedInformation) SlotSegment() string {
	if o.SegmentGroupNumber == 0xFFFF || o.BusNumber == 0xFF || o.DeviceFunctionNumber == 0xFF {
		return "Not of types PCI/AGP/PCI-X/PCI-Express"
	}
	return fmt.Sprintf("Bus Address: %04x:%02x:%02x.%x",
		o.SegmentGroupNumber,
		o.BusNumber,
		o.DeviceFunctionNumber>>3,
		o.DeviceFunctionNumber&0x7)
}

func (o ExtendedInformation) String() string {
	return fmt.Sprintf("On Board Devices Extended Information\n"+
		"\tReference Designation: %s\n"+
		"\tDevice Type: %s\n"+
		"\tDevice Status: %s\n"+
		"\tDevice Type Instance: %d\n"+
		"%s\n",
		o.ReferenceDesignation,
		o.DeviceType,
		o.DeviceStatus,
		o.DeviceTypeInstance,
		o.SlotSegment())
}
