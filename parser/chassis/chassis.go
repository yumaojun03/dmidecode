package chassis

import (
	"fmt"

	"dmidecode/smbios"
)

// Information 底座信息
type Information struct {
	smbios.Header
	Manufacturer                 string
	Type                         Type
	Lock                         Lock
	Version                      string
	AssetTag                     string
	SerialNumber                 string
	BootUpState                  State
	PowerSupplyState             State
	ThermalState                 State
	SecurityStatus               SecurityStatus
	OEMdefined                   uint16
	Height                       Height
	NumberOfPowerCords           byte
	ContainedElementCount        byte
	ContainedElementRecordLength byte
	ContainedElements            ContainedElements
	SKUNumber                    string
}

func (c Information) String() string {
	return fmt.Sprintf("Chassis Information\n"+
		"\tManufacturer: %s\n"+
		"\tType: %s\n"+
		"\tLock: %s\n"+
		"\tVersion: %s\n"+
		"\tSerial Number: %s\n"+
		"\tAsset Tag: %s\n"+
		"\tBoot-up State: %s\n"+
		"\tPower Supply State: %s\n"+
		"\tThermal State: %s\n"+
		"\tSecurity Status: %s\n"+
		"\tOEM Information: %x\n"+
		"\tHeight: %s\n"+
		"\tNumber Of Power Cords: %d\n"+
		"\tContained Elements: %d\n"+
		"\tSKU Number: %s",
		c.Manufacturer,
		c.Type,
		c.Lock,
		c.Version,
		c.SerialNumber,
		c.AssetTag,
		c.BootUpState,
		c.PowerSupplyState,
		c.ThermalState,
		c.SecurityStatus,
		c.OEMdefined,
		c.Height,
		c.NumberOfPowerCords,
		c.ContainedElements,
		c.SKUNumber)
}
