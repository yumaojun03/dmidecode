package bios

import (
	"fmt"

	"dmidecode/smbios"
)

// Information bios信息
type Information struct {
	smbios.Header
	Vendor                                 string
	BIOSVersion                            string
	StartingAddressSegment                 uint16
	ReleaseDate                            string
	RomSize                                RomSize
	RuntimeSize                            RuntimeSize
	Characteristics                        Characteristics
	CharacteristicsExt1                    Ext1
	CharacteristicsExt2                    Ext2
	SystemBIOSMajorRelease                 byte
	SystemBIOSMinorRelease                 byte
	EmbeddedControllerFirmwareMajorRelease byte
	EmbeddedControllerFirmawreMinorRelease byte
}

func (b Information) String() string {
	s := fmt.Sprintf("Handle %x, DMI type %d, %d bytes\n"+
		"\tBIOS Information\n"+
		"\tVendor: %s\n"+
		"\tVersion: %s\n"+
		"\tRelease Date: %s\n"+
		"\tAddress: 0x%4X0\n"+
		"\tRuntime Size: %s\n"+
		"\tROM Size: %s\n"+
		"\tCharacteristics:%s",
		b.Header.Handle,
		b.Header.Type,
		b.Header.Length,
		b.Vendor,
		b.BIOSVersion,
		b.ReleaseDate,
		b.StartingAddressSegment,
		b.RuntimeSize,
		b.RomSize,
		b.Characteristics)

	if b.CharacteristicsExt1 != 0 {
		s += b.CharacteristicsExt1.String()
	}
	if b.CharacteristicsExt2 != 0 {
		s += b.CharacteristicsExt2.String()
	}
	return s
}
