package chassis

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

// Information 底座信息
type Information struct {
	smbios.Header
	Manufacturer                 string            `json:"manufacturer,omitempty"`
	Type                         Type              `json:"type,omitempty"`
	Lock                         Lock              `json:"lock,omitempty"`
	Version                      string            `json:"version,omitempty"`
	AssetTag                     string            `json:"asset_tag,omitempty"`
	SerialNumber                 string            `json:"serial_number,omitempty"`
	BootUpState                  State             `json:"boot_up_state,omitempty"`
	PowerSupplyState             State             `json:"power_supply_state,omitempty"`
	ThermalState                 State             `json:"thermal_state,omitempty"`
	SecurityStatus               SecurityStatus    `json:"security_status,omitempty"`
	OEMdefined                   uint16            `json:"oe_mdefined,omitempty"`
	Height                       Height            `json:"height,omitempty"`
	NumberOfPowerCords           byte              `json:"number_of_power_cords,omitempty"`
	ContainedElementCount        byte              `json:"contained_element_count,omitempty"`
	ContainedElementRecordLength byte              `json:"contained_element_record_length,omitempty"`
	ContainedElements            ContainedElements `json:"contained_elements,omitempty"`
	SKUNumber                    string            `json:"sku_number,omitempty"`
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
