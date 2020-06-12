package onboard

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

// ParseType41 解析
func ParseType41(s *smbios.Structure) (*ExtendedInformation, error) {
	if s.Type() != 41 {
		return nil, fmt.Errorf("ParseType41 only parse type 41 data, but now: %d", s.Type())
	}

	data := s.Formatted

	info := &ExtendedInformation{
		ReferenceDesignation: s.GetString(0x0),
		DeviceType:           ExtendedInformationType(data[0x01] & 127),
		DeviceStatus:         DeviceStatus(data[0x01] >> 7),
		DeviceTypeInstance:   data[0x02],
		SegmentGroupNumber:   smbios.U16(data[0x03:0x05]),
		BusNumber:            data[0x05],
		DeviceFunctionNumber: data[0x06],
	}

	if len(s.Formatted) >= 0x06 {
		info.DeviceFunctionNumber = data[0x06]
	}

	return info, nil
}

// ParseType10 解析
func ParseType10(s *smbios.Structure) (*Information, error) {
	if s.Type() != 10 {
		return nil, fmt.Errorf("ParseType10 only parse type 10 data, but now: %d", s.Type())
	}

	data := s.Formatted

	var devices []Device
	for i := 0x0; i <= len(data)-1; i += 2 {
		d := Device{
			DeviceStatus: DeviceStatus(data[i] >> 7),
			DeviceType:   ExtendedInformationType(data[i] & 127),
			Description:  s.GetString(i + 1),
		}
		devices = append(devices, d)
	}

	info := &Information{
		Devices: devices,
	}

	return info, nil
}
