package baseboard

import "dmidecode/smbios"

// Parse 解析
func Parse(s *smbios.Structure) (*Information, error) {
	data := s.Formatted

	info := &Information{
		Header:            s.Header,
		Manufacturer:      s.GetString(0),
		ProductName:       s.GetString(1),
		Version:           s.GetString(2),
		SerialNumber:      s.GetString(3),
		AssetTag:          s.GetString(4),
		LocationInChassis: s.GetString(5),
	}

	if len(s.Formatted) >= 0x05 {
		info.FeatureFlags = FeatureFlags(data[0x05])
	}

	if len(s.Formatted) >= 0x09 {
		info.BoardType = Type(data[0x09])
	}

	return info, nil
}
