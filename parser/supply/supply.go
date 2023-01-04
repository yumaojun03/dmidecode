package supply

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

// SystemPowerSupply 电源设备
type SystemPowerSupply struct {
	smbios.Header
	PowerUnitGroup             uint32
	Location                   string
	DeviceName                 string
	Manufacturer               string
	SerialNumber               string
	AssetTag                   string
	ModelPartNumber            string
	MaxPowerCapacity           uint8 // 最大功率数
	PowerSupplyCharacteristics uint16
	Type                       string // 电源类型
	Status                     string
	HotReplaceable             uint8 // 是否支持热拔插
	Present                    uint8 // 在位状态
	Unplugged                  uint8 // 是否可插拔
}

func (s SystemPowerSupply) String() string {
	return fmt.Sprintf("System Power Supply\n"+
		"\tPower Unit Group: %d\n"+
		"\tLocation: %s\n"+
		"\tDevice Name: %s\n"+
		"\tManufacturer: %s\n"+
		"\tSerial Number: %s\n"+
		"\tAsset Tag Number: %s\n"+
		"\tModel Part Number: %s\n"+
		"\tStatus: %s\n"+
		"\tType: %s\n"+
		"\tHot Replaceable: %d\n"+
		"\tPresent: %d\n"+
		"\tUnplugged: %d\n",
		s.PowerUnitGroup,
		s.Location,
		s.DeviceName,
		s.Manufacturer,
		s.SerialNumber,
		s.AssetTag,
		s.ModelPartNumber,
		s.Status,
		s.Type,
		s.HotReplaceable,
		s.Present,
		s.Unplugged)
}
