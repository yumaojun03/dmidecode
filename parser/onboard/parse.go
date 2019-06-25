package onboard

import (
	"dmidecode/smbios"
)

// Parse 解析
func Parse(s *smbios.Structure) (*ExtendedInformation, error) {
	data := s.Formatted

	info := &ExtendedInformation{
		ReferenceDesignation: s.GetString(0),
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
