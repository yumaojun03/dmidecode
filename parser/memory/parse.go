package memory

import (
	"github.com/yumaojun03/dmidecode/smbios"
)

// ParseMemoryDevice 解析
func ParseMemoryDevice(s *smbios.Structure) (info *MemoryDevice, err error) {
	defer smbios.ParseRecovery(s, &err)

	info = &MemoryDevice{
		PhysicalMemoryArrayHandle:  s.U16(0x00, 0x02),
		ErrorInformationHandle:     s.U16(0x02, 0x04),
		TotalWidth:                 s.U16(0x04, 0x06),
		DataWidth:                  s.U16(0x06, 0x08),
		Size:                       s.U16(0x08, 0x0a),
		FormFactor:                 MemoryDeviceFormFactor(s.GetByte(0x0a)),
		DeviceSet:                  s.GetByte(0x0b),
		DeviceLocator:              s.GetString(0xc),
		BankLocator:                s.GetString(0xd),
		Type:                       MemoryDeviceType(s.GetByte(0xe)),
		TypeDetail:                 MemoryDeviceTypeDetail(s.U16(0xf, 0x11)),
		Speed:                      s.U16(0x11, 0x13),
		Manufacturer:               s.GetString(0x13),
		SerialNumber:               s.GetString(0x14),
		AssetTag:                   s.GetString(0x15),
		PartNumber:                 s.GetString(0x16),
		Attributes:                 s.GetByte(0x17),
		ExtendedSize:               s.U32(0x18, 0x1c),
		ConfiguredMemoryClockSpeed: s.U16(0x1c, 0x1e),
		MinimumVoltage:             s.U16(0x1e, 0x20),
		MaximumVoltage:             s.U16(0x20, 0x22),
		ConfiguredVoltage:          s.U16(0x22, 0x24),
	}

	return info, nil
}

// ParseMemoryArray todo
func ParseMemoryArray(s *smbios.Structure) (info *PhysicalMemoryArray, err error) {
	defer smbios.ParseRecovery(s, &err)

	info = &PhysicalMemoryArray{
		Header:                  s.Header,
		Location:                PhysicalMemoryArrayLocation(s.GetByte(0x00)),
		Use:                     PhysicalMemoryArrayUse(s.GetByte(0x01)),
		ErrorCorrection:         PhysicalMemoryArrayErrorCorrection(s.GetByte(0x02)),
		MaximumCapacity:         s.U32(0x03, 0x07),
		ErrorInformationHandle:  s.U16(0x07, 0x09),
		NumberOfMemoryDevices:   s.U16(0x09, 0xb),
		ExtendedMaximumCapacity: s.U64(0xb, 0x13),
	}

	return info, nil
}
