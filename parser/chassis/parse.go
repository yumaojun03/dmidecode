package chassis

import "dmidecode/smbios"

// Parse 解析底座信息
func Parse(s *smbios.Structure) (*Information, error) {
	data := s.Formatted

	info := &Information{
		Manufacturer:                 s.GetString(0),
		Type:                         Type(data[0x01] & 127),
		Lock:                         Lock(data[0x01] >> 7),
		Version:                      s.GetString(1),
		SerialNumber:                 s.GetString(2),
		AssetTag:                     s.GetString(3),
		BootUpState:                  State(data[0x05]),
		PowerSupplyState:             State(data[0x06]),
		ThermalState:                 State(data[0x07]),
		SecurityStatus:               SecurityStatus(data[0x08]),
		OEMdefined:                   smbios.U16(data[0x09 : 0x09+4]),
		Height:                       Height(data[0xD]),
		NumberOfPowerCords:           data[0xE],
		ContainedElementCount:        data[0xF],
		ContainedElementRecordLength: data[0x10],
		// TODO: 7.4.4
		//ci.ContainedElements:
		SKUNumber: s.GetString(4),
	}

	return info, nil
}
