package supply

type Characteristics uint16

const (
	CharacteristicsHotReplaceable Characteristics = 1 << iota
	CharacteristicsPresent
	CharacteristicsUnplugged
)

func (m Characteristics) HotReplaceable() uint8 {
	return uint8(m&CharacteristicsHotReplaceable) >> 0
}

func (m Characteristics) Present() uint8 {
	return uint8(m&CharacteristicsPresent) >> 1
}

func (m Characteristics) Unplugged() uint8 {
	return uint8(m&CharacteristicsUnplugged) >> 2
}

func (m Characteristics) Type() PowerSupplyType {
	return PowerSupplyType(m & 0x3c00 >> 10)
}

func (m Characteristics) Status() Status {
	return Status(m & 0x01c0 >> 6)
}

type PowerSupplyType uint32

const (
	PowerSupplyTypeOther           PowerSupplyType = 0x01
	PowerSupplyTypeUnknown         PowerSupplyType = 0x02
	PowerSupplyTypeLinear          PowerSupplyType = 0x03
	PowerSupplyTypeSwitching       PowerSupplyType = 0x04
	PowerSupplyTypeBattery         PowerSupplyType = 0x05
	PowerSupplyTypeUPS             PowerSupplyType = 0x06
	PowerSupplyTypeConverter       PowerSupplyType = 0x07
	PowerSupplyTypeRegulator       PowerSupplyType = 0x08
	PowerSupplyTypeAssignmentStart PowerSupplyType = 0x09
	PowerSupplyTypeAssignmentEnd   PowerSupplyType = 0x0f
)

func (p PowerSupplyType) String() string {
	switch p {
	case PowerSupplyTypeOther:
		return "Other"
	case PowerSupplyTypeUnknown:
		return "Unknown"
	case PowerSupplyTypeLinear:
		return "Linear"
	case PowerSupplyTypeSwitching:
		return "Switching"
	case PowerSupplyTypeBattery:
		return "Battery"
	case PowerSupplyTypeUPS:
		return "UPS"
	case PowerSupplyTypeConverter:
		return "Converter"
	case PowerSupplyTypeRegulator:
		return "Regulator"
	default:
		return "Reserved for future assignment"
	}
}

type Status uint32

const (
	StatusOther       Status = 0x01
	StatusUnknown     Status = 0x02
	StatusOK          Status = 0x03
	StatusNonCritical Status = 0x04
	StatusCritical    Status = 0x05
)

func (p Status) String() string {
	switch p {
	case StatusOther:
		return "Other"
	case StatusUnknown:
		return "Unknown"
	case StatusOK:
		return "OK"
	case StatusNonCritical:
		return "Non-critical"
	case StatusCritical:
		return "Critical"
	default:
		return "Unknown"
	}
}
