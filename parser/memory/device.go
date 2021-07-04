package memory

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

// MemoryDevice 内存设备
type MemoryDevice struct {
	smbios.Header
	PhysicalMemoryArrayHandle  uint16
	ErrorInformationHandle     uint16
	TotalWidth                 uint16
	DataWidth                  uint16
	Size                       uint16
	FormFactor                 MemoryDeviceFormFactor
	DeviceSet                  byte
	DeviceLocator              string
	BankLocator                string
	Type                       MemoryDeviceType
	TypeDetail                 MemoryDeviceTypeDetail
	Speed                      uint16
	Manufacturer               string
	SerialNumber               string
	AssetTag                   string
	PartNumber                 string
	Attributes                 byte
	ExtendedSize               uint32
	ConfiguredMemoryClockSpeed uint16
	MinimumVoltage             uint16
	MaximumVoltage             uint16
	ConfiguredVoltage          uint16
}

func (m MemoryDevice) String() string {
	return fmt.Sprintf("Memory Device\n"+
		"\tPhysical Memory Array Handle: %d\n"+
		"\tMemory Error Information Handle: %d\n"+
		"\tTotal Width: %d\n"+
		"\tData Width: %d\n"+
		"\tSize: %d\n"+
		"\tForm Factor: %s\n"+
		"\tDevice Set: %d\n"+
		"\tDevice Locator: %s\n"+
		"\tBank Locator: %s\n"+
		"\tMemory Type: %s\n"+
		"\tType Detail: %s\n"+
		"\tSpeed: %d\n"+
		"\tManufacturer: %s\n"+
		"\tSerial Number: %s\n"+
		"\tAsset Tag: %s\n"+
		"\tPart Number: %s\n"+
		"\tAttributes: %x\n"+
		"\tExtended Size: %d\n"+
		"\tConfigured Memory Clock Speed: %d\n"+
		"\tMinimum voltage: %d\n"+
		"\tMaximum voltage: %d\n"+
		"\tConfigured voltage: %d",
		m.PhysicalMemoryArrayHandle,
		m.ErrorInformationHandle,
		m.TotalWidth,
		m.DataWidth,
		m.Size,
		m.FormFactor,
		m.DeviceSet,
		m.DeviceLocator,
		m.BankLocator,
		m.Type,
		m.TypeDetail,
		m.Speed,
		m.Manufacturer,
		m.SerialNumber,
		m.AssetTag,
		m.PartNumber,
		m.Attributes,
		m.ExtendedSize,
		m.ConfiguredMemoryClockSpeed,
		m.MinimumVoltage,
		m.MaximumVoltage,
		m.ConfiguredVoltage,
	)
}

//The value is in kilobytes
func (m MemoryDevice) ActualSize() uint32 {
	if m.Size == 0 || m.Size == 0xFFFF {
		return 0
	}
	if m.Size == 0x7FFF {
		return m.ExtendedSize * 1024
	}
	if m.Size&0x8000 != 0 {
		return uint32(m.Size &^ 0x8000)
	}
	return uint32(m.Size) * 1024
}
