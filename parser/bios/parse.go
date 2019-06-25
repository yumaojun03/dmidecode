package bios

import (
	"dmidecode/smbios"
)

// Parse 解析
func Parse(s *smbios.Structure) (*Information, error) {
	data := s.Formatted
	sas := smbios.U16(data[0x02:0x04])

	bi := &Information{
		Header:                 s.Header,
		Vendor:                 s.Strings[0],
		BIOSVersion:            s.Strings[1],
		ReleaseDate:            s.Strings[2],
		StartingAddressSegment: sas,
		RomSize:                RomSize(64 * (data[0x05] + 1)),
		RuntimeSize:            RuntimeSize((uint(0x10000) - uint(sas)) << 4),
		Characteristics:        Characteristics(smbios.U64(data[0x06:0x08])),
	}

	if s.Header.Length >= 0x13 {
		bi.CharacteristicsExt1 = Ext1(data[0x08])
	}
	if s.Header.Length >= 0x14 {
		bi.CharacteristicsExt2 = Ext2(data[0x09])
	}

	return bi, nil
}
