package power

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

type SupplyCharacteristics uint16

func IntRange(start, end uint16) []uint16 {
	ret := make([]uint16, 0)
	for i := start; i <= end; i++ {
		ret = append(ret, i)
	}
	return ret
}

func Range(a ...uint16) []uint16 {
	ret := make([]uint16, 0)
	ret = append(ret, a...)

	return ret
}

const (
	outOfSpec = "<OUT OF SPEC>"
)

var (
	SupplyCharacteristicsHotReplaceable             = Range(0x00)
	SupplyCharacteristicsPresent                    = Range(0x01)
	SupplyCharacteristicsUnpluggedWall              = Range(0x02)
	SupplyCharacteristicsInputVoltageRangeSwitching = IntRange(3, 6)
	SupplyCharacteristicsStatus                     = IntRange(7, 9)
	SupplyCharacteristicsType                       = IntRange(10, 13)
)

type Information struct {
	smbios.Header
	PowerUnitGroup          byte                  `json:"power_unit_group"`
	Location                string                `json:"location,omitempty"`
	DeviceName              string                `json:"device_name,omitempty"`
	Manufacturer            string                `json:"manufacturer,omitempty"`
	SerialNumber            string                `json:"serial_number,omitempty"`
	AssetTagNumber          string                `json:"asset_tag_number"`
	ModelPartNumber         string                `json:"model_part_number"`
	RevisionLevel           string                `json:"revision_level"`
	MaxPowerCapacity        uint16                `json:"max_power_capacity,omitempty"`
	Characteristics         SupplyCharacteristics `json:"power_supply_characteristics"`
	InputVoltageProbeHandle uint16                `json:"input_voltage_probe_handle"`
	CoolingDeviceHandle     uint16                `json:"cooling_device_handle"`
	InputCurrentProbeHandle uint16                `json:"input_current_probe_handle"`
}

func (s SupplyCharacteristics) Is(obj []uint16) bool {
	for _, o := range obj {
		if o == uint16(s) {
			return true
		}
	}
	return false
}

func (s SupplyCharacteristics) PowerSupplyHotReplaceable() string {
	if s&(1<<0) > 0 {
		return "Yes"
	}
	return "No"
}

func (s SupplyCharacteristics) PowerSupplyUnpluggedWall() string {
	if s&(1<<2) > 0 {
		return "No"
	}
	return "Yes"
}

func (s SupplyCharacteristics) PowerSupplyStatus() string {
	chars := []string{
		"Other",        // 001b
		"Unknown",      //010b
		"Ok",           //011b
		"Non-critical", //100b
		"Critical",     //101b
	}
	if s&(1<<1) > 0 {
		t := (s >> 7) & 0x07
		if t >= 0x01 && t <= 0x05 {
			return fmt.Sprintf("Status: Present, %s", chars[t-0x01])
		}
		return outOfSpec
	} else {
		return fmt.Sprintf("Status: Not Present")
	}
}

func (s SupplyCharacteristics) PowerSupplyType() string {
	chars := []string{
		"Other",     // 0001b
		"Unknown",   // 0010b
		"Linear",    //0011b
		"Switching", //0100b
		"Battery",   //0101b
		"UPS",       //0110b
		"Converter", //0111b
		"Regulator", // 1000b
	}

	t := (s >> 10) & (0x0F)
	if t >= 0x01 && t <= 0x08 {
		return chars[t-0x01]
	}
	return outOfSpec
}

func (s SupplyCharacteristics) PowerSupplyRangeSwitching() string {
	chars := []string{
		"Other",          // 0001b
		"Unknown",        // 0010b
		"Manual",         // 0011b
		"Auto-switch",    // 0100b
		"Wide range",     // 0101b
		"Not applicable", // 0110b
	}
	t := (s >> 3) & (0x0F)
	if t >= 0x01 && t <= 0x06 {
		return chars[t-0x01]
	}
	return outOfSpec
}

func (s SupplyCharacteristics) String() string {
	return fmt.Sprintf("%d", s)
}

func (p *Information) String() string {
	return fmt.Sprintf(
		"Power Supply Information:\n"+
			"\tPower Unit Group: %d\n"+
			"\tLocation: %s\n"+
			"\tDeviceName: %s\n"+
			"\tManufacturer: %s\n"+
			"\tSerial Number: %s\n"+
			"\tAsset Tag Number: %s\n"+
			"\tModel Part Number: %s\n"+
			"\tRevision Level: %s\n"+
			"\tMax Power Capacity: %dW\n"+
			"\tInput Voltage Range Switching: %s\n"+
			"\tPlugged: %s\n"+
			"\tHot Replaceable: %s\n"+
			"\tType: %s\n"+
			"\t%s\n",
		int(p.PowerUnitGroup),
		p.Location,
		p.DeviceName,
		p.Manufacturer,
		p.SerialNumber,
		p.AssetTagNumber,
		p.ModelPartNumber,
		p.RevisionLevel,
		p.MaxPowerCapacity,
		p.Characteristics.PowerSupplyRangeSwitching(),
		p.Characteristics.PowerSupplyUnpluggedWall(),
		p.Characteristics.PowerSupplyHotReplaceable(),
		p.Characteristics.PowerSupplyType(),
		p.Characteristics.PowerSupplyStatus(),
	)
}
