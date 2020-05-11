package onboard

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

// ExtendedInformation 扩展设备信息
type ExtendedInformation struct {
	smbios.Header
	ReferenceDesignation string                  `json:"reference_designation,omitempty"`
	DeviceType           ExtendedInformationType `json:"device_type,omitempty"`
	DeviceStatus         DeviceStatus            `json:"device_status,omitempty"`
	DeviceTypeInstance   byte                    `json:"device_type_instance,omitempty"`
	SegmentGroupNumber   uint16                  `json:"segment_group_number,omitempty"`
	BusNumber            byte                    `json:"bus_number,omitempty"`
	DeviceFunctionNumber byte                    `json:"device_function_number,omitempty"`
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
		"\t%s\n",
		o.ReferenceDesignation,
		o.DeviceType,
		o.DeviceStatus,
		o.DeviceTypeInstance,
		o.SlotSegment())
}
