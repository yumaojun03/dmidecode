package baseboard

import (
	"github.com/yumaojun03/dmidecode/smbios"
)

// Parse 解析
func Parse(s *smbios.Structure) (*Information, error) {
	data := s.Formatted

	info := &Information{
		Header:            s.Header,
		Manufacturer:      s.GetString(0x0),
		ProductName:       s.GetString(0x1),
		Version:           s.GetString(0x2),
		SerialNumber:      s.GetString(0x3),
		AssetTag:          s.GetString(0x4),
		LocationInChassis: s.GetString(0x6),
	}

	if len(s.Formatted) >= 0x05 {
		info.FeatureFlags = FeatureFlags(data[0x05])
	}

	if len(s.Formatted) >= 0x09 {
		info.BoardType = Type(data[0x09])
	}

	return info, nil
}
