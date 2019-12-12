package bios

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

// Information bios信息
type Information struct {
	smbios.Header
	Vendor                                 string          `json:"vendor,omitempty"`
	BIOSVersion                            string          `json:"bios_version,omitempty"`
	StartingAddressSegment                 uint16          `json:"starting_address_segment,omitempty"`
	ReleaseDate                            string          `json:"release_date,omitempty"`
	RomSize                                RomSize         `json:"rom_size,omitempty"`
	RuntimeSize                            RuntimeSize     `json:"runtime_size,omitempty"`
	Characteristics                        Characteristics `json:"characteristics,omitempty"`
	CharacteristicsExt1                    Ext1            `json:"characteristics_ext_1,omitempty"`
	CharacteristicsExt2                    Ext2            `json:"characteristics_ext_2,omitempty"`
	SystemBIOSMajorRelease                 byte            `json:"system_bios_major_release,omitempty"`
	SystemBIOSMinorRelease                 byte            `json:"system_bios_minor_release,omitempty"`
	EmbeddedControllerFirmwareMajorRelease byte            `json:"firmware_major_release,omitempty"`
	EmbeddedControllerFirmawreMinorRelease byte            `json:"firmawre_minor_release,omitempty"`
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
