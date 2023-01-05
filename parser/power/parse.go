package power

import (
	"strings"

	"github.com/yumaojun03/dmidecode/smbios"
)

func Parse(s *smbios.Structure) (info *Information, err error) {
	defer smbios.ParseRecovery(s, &err)
	info = &Information{
		Header:           s.Header,
		PowerUnitGroup:   s.GetByte(0x0),
		Location:         strings.TrimRight(s.GetString(0x1), " "),
		DeviceName:       strings.TrimRight(s.GetString(0x2), " "),
		Manufacturer:     strings.TrimRight(s.GetString(0x3), " "),
		SerialNumber:     strings.TrimRight(s.GetString(0x4), " "),
		AssetTagNumber:   s.GetString(0x5),
		ModelPartNumber:  s.GetString(0x6),
		RevisionLevel:    s.GetString(0x7),
		MaxPowerCapacity: s.U16(0x08, 0x0A),
		Characteristics:  SupplyCharacteristics(s.U16(0x0A, 0x0C)),
	}
	return info, err
}
