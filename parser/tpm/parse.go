package tpm

import "github.com/yumaojun03/dmidecode/smbios"

// Parse 解析smbios struct数据
func Parse(s *smbios.Structure) (info *Information, err error) {
	defer smbios.ParseRecovery(s, &err)
	info = &Information{
		Header:       s.Header,
		VendorID:        string(s.GetBytes(0x00, 0x04)),
		SpecificationVersion: SpecificationVersion{s.GetByte(0x04), s.GetByte(0x05)}.String(),
		FirmwareRevision: FirmwareRevision{s.U32(0x06, 0x0A), s.U32(0x0A, 0x0E)}.String(),
		Description: s.GetString(0x0E),
		Characteristics: TpmDeviceCharacteristics{s.U64(0x0F, 0x17)}.String(),
		OEMSpecificInformation: OEMSpecificInformation{s.U32(0x17, 0x1B)}.String(),
	}
	return info, nil
}