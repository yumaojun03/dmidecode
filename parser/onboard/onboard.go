package onboard

import (
	"fmt"
	"strings"

	"github.com/yumaojun03/dmidecode/smbios"
)

// Information 设备信息 (smbios < 2.6)
// 参考: https://www.dmtf.org/sites/default/files/standards/documents/DSP0134_3.1.1.pdf
//      7.11 On Board Devices Information (Type 10, Obsolete)
type Information struct {
	smbios.Header
	Devices []Device `json:"devices,omitempty"`
}

// Device 设备
type Device struct {
	Description  string                  `json:"description,omitempty"`
	DeviceStatus DeviceStatus            `json:"device_status,omitempty"`
	DeviceType   ExtendedInformationType `json:"device_type,omitempty"`
}

func (i Information) String() string {
	str := []string{}
	for index, d := range i.Devices {
		deviceStr := fmt.Sprintf("On Board Device %d Information\n"+
			"\tDevice Type: %s\n"+
			"\tDevice Status: %s\n"+
			"\tDescription: %s\n",
			index+1,
			d.DeviceType,
			d.DeviceStatus,
			d.Description)
		str = append(str, deviceStr)
	}

	return strings.Join(str, "")
}

// ExtendedInformation 扩展设备信息 (smbios >= 2.6)
// 参考文档: https://www.dmtf.org/sites/default/files/standards/documents/DSP0134_3.1.1.pdf
// 7.42 Onboard Devices Extended Information (Type 41)
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
