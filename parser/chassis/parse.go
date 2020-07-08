package chassis

import (
	"github.com/yumaojun03/dmidecode/smbios"
)

// Parse 解析底座信息
func Parse(s *smbios.Structure) (info *Information, err error) {
	defer smbios.ParseRecovery(s, &err)

	info = &Information{
		Manufacturer:                 s.GetString(0x0),
		Type:                         Type(s.GetByte(0x01) & 127),
		Lock:                         Lock(s.GetByte(0x01) >> 7),
		Version:                      s.GetString(0x2),
		SerialNumber:                 s.GetString(0x3),
		AssetTag:                     s.GetString(0x4),
		BootUpState:                  State(s.GetByte(0x05)),
		PowerSupplyState:             State(s.GetByte(0x06)),
		ThermalState:                 State(s.GetByte(0x07)),
		SecurityStatus:               SecurityStatus(s.GetByte(0x08)),
		OEMdefined:                   s.U16(0x09, 0xd),
		Height:                       Height(s.GetByte(0xd)),
		NumberOfPowerCords:           s.GetByte(0xe),
		ContainedElementCount:        s.GetByte(0xf),
		ContainedElementRecordLength: s.GetByte(0x10),
		// TODO: 7.4.4
		//ci.ContainedElements:
		SKUNumber: s.GetString(4),
	}

	return info, nil
}
