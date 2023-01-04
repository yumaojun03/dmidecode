package supply

import "github.com/yumaojun03/dmidecode/smbios"

// Parse 解析电源信息
func Parse(s *smbios.Structure) (info *SystemPowerSupply, err error) {
	defer smbios.ParseRecovery(s, &err)

	info = &SystemPowerSupply{
		PowerUnitGroup:   uint32(s.GetByte(0x00)),
		Location:         s.GetString(0x01),
		DeviceName:       s.GetString(0x02),
		Manufacturer:     s.GetString(0x03),
		SerialNumber:     s.GetString(0x04),
		AssetTag:         s.GetString(0x05),
		ModelPartNumber:  s.GetString(0x06),
		MaxPowerCapacity: s.GetByte(0x08),
		Type:             Characteristics(s.U16(0x0a, 0x0c)).Type().String(),
		Status:           Characteristics(s.U16(0x0a, 0x0c)).Status().String(),
		HotReplaceable:   Characteristics(s.U16(0x0a, 0x0c)).HotReplaceable(),
		Present:          Characteristics(s.U16(0x0a, 0x0c)).Present(),
		Unplugged:        Characteristics(s.U16(0x0a, 0x0c)).Unplugged(),
	}

	return info, nil
}
