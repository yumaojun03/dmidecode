package bios

import (
	"github.com/yumaojun03/dmidecode/smbios"
)

// Parse 参考 https://www.dmtf.org/sites/default/files/standards/documents/DSP0134_3.1.1.pdf
func Parse(s *smbios.Structure) (info *Information, err error) {
	defer smbios.ParseRecovery(s, &err)

	sas := s.U16(0x02, 0x04)

	info = &Information{
		Header:                 s.Header,
		Vendor:                 s.GetString(0x0),
		BIOSVersion:            s.GetString(0x1),
		ReleaseDate:            s.GetString(0x4),
		StartingAddressSegment: sas,
		RomSize:                RomSize(64 * (s.GetByte(0x05) + 1)),
		RuntimeSize:            RuntimeSize((uint(0x10000) - uint(sas)) << 4),
		Characteristics:        Characteristics(s.U64(0x06, 0x08)),
		CharacteristicsExt1:    Ext1(s.GetByte(0x08)),
		CharacteristicsExt2:    Ext2(s.GetByte(0x09)),
	}

	return info, nil
}
