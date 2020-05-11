package memory

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

// ParseMemoryDevice 解析
func ParseMemoryDevice(s *smbios.Structure) (*MemoryDevice, error) {
	data := s.Formatted

	fmt.Println(smbios.U16(data[0xF:0x11]))
	md := &MemoryDevice{
		PhysicalMemoryArrayHandle:  smbios.U16(data[0x00:0x02]),
		ErrorInformationHandle:     smbios.U16(data[0x02:0x04]),
		TotalWidth:                 smbios.U16(data[0x04:0x06]),
		DataWidth:                  smbios.U16(data[0x06:0x08]),
		Size:                       smbios.U16(data[0x08:0x0A]),
		FormFactor:                 MemoryDeviceFormFactor(data[0x0A]),
		DeviceSet:                  data[0x0B],
		DeviceLocator:              s.GetString(0xc),
		BankLocator:                s.GetString(0xd),
		Type:                       MemoryDeviceType(data[0xE]),
		TypeDetail:                 MemoryDeviceTypeDetail(smbios.U16(data[0xF:0x11])),
		Speed:                      smbios.U16(data[0x11:0x13]),
		Manufacturer:               s.GetString(0x13),
		SerialNumber:               s.GetString(0x14),
		PartNumber:                 s.GetString(0x16),
		Attributes:                 data[0x17],
		ExtendedSize:               smbios.U32(data[0x18:0x1c]),
		ConfiguredMemoryClockSpeed: smbios.U16(data[0x1c:0x1e]),
		MinimumVoltage:             smbios.U16(data[0x1e:0x20]),
		MaximumVoltage:             smbios.U16(data[0x20:0x22]),
		ConfiguredVoltage:          smbios.U16(data[0x22:0x24]),
	}

	return md, nil
}

// ParseMemoryArray todo
func ParseMemoryArray(s *smbios.Structure) (*PhysicalMemoryArray, error) {
	data := s.Formatted

	mr := &PhysicalMemoryArray{
		Header:                  s.Header,
		Location:                PhysicalMemoryArrayLocation(data[0x00]),
		Use:                     PhysicalMemoryArrayUse(data[0x01]),
		ErrorCorrection:         PhysicalMemoryArrayErrorCorrection(data[0x02]),
		MaximumCapacity:         smbios.U32(data[0x03:0x07]),
		ErrorInformationHandle:  smbios.U16(data[0x07:0x09]),
		NumberOfMemoryDevices:   smbios.U16(data[0x09:0xB]),
		ExtendedMaximumCapacity: smbios.U64(data[0x0B:]),
	}

	return mr, nil
}
