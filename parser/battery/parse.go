package battery

import (
	"github.com/yumaojun03/dmidecode/smbios"
)

// Parse 解析電池信息
func Parse(s *smbios.Structure) (info *Information, err error) {
	defer smbios.ParseRecovery(s, &err)

	info = &Information{
		Header:                   s.Header,
		Location:                 s.GetString(0x00),
		Manufacturer:             s.GetString(0x01),
		ManufacturerDate:         s.GetString(0x02),
		SerialNumber:             s.GetString(0x03),
		DeviceName:               s.GetString(0x04),
		DeviceChemistry:          DeviceChemistry(s.GetByte(0x05)),
		DesignCapacity:           s.U16(0x06, 0x08),
		DesignVoltage:            s.U16(0x08, 0x0a),
		SBDSVersionNumber:        s.GetString(0x0a),
		MaximumError:             s.U16(0x0b, 0x0c),
		SBDSManufactureDate:      s.U16(0x0c, 0x0e),
		SBDSDeviceChemistry:      s.GetString(0x0e),
		DesignCapacityMultiplier: s.U16(0x0f, 0x10),
		OEMSpecific:              s.U16(0x10, 0x11),
	}

	return info, nil

}
