package baseboard

import (
	"github.com/yumaojun03/dmidecode/smbios"
)

// Parse 解析
func Parse(s *smbios.Structure) (info *Information, err error) {
	defer smbios.ParseRecovery(s, &err)

	info = &Information{
		Header:            s.Header,
		Manufacturer:      s.GetString(0x0),
		ProductName:       s.GetString(0x1),
		Version:           s.GetString(0x2),
		SerialNumber:      s.GetString(0x3),
		AssetTag:          s.GetString(0x4),
		LocationInChassis: s.GetString(0x6),
		FeatureFlags:      FeatureFlags(s.GetByte(0x05)),
		BoardType:         Type(s.GetByte(0x09)),
	}
	return info, err
}
