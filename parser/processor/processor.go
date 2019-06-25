package processor

import (
	"dmidecode/smbios"
	"fmt"
)

type Information struct {
	smbios.Header
	SocketDesignation string
	ProcessorType     ProcessorType
	Family            ProcessorFamily
	Manufacturer      string
	ID                ProcessorID
	Version           string
	Voltage           ProcessorVoltage
	ExternalClock     uint16
	MaxSpeed          uint16
	CurrentSpeed      uint16
	Status            ProcessorStatus
	Upgrade           ProcessorUpgrade
	L1CacheHandle     uint16
	L2CacheHandle     uint16
	L3CacheHandle     uint16
	SerialNumber      string
	AssetTag          string
	PartNumber        string
	CoreCount         byte
	CoreEnabled       byte
	ThreadCount       byte
	Characteristics   ProcessorCharacteristics
	Family2           ProcessorFamily
}

func (p Information) String() string {
	return fmt.Sprintf("Processor Information\n"+
		"\tSocket Designation: %s\n"+
		"\tProcessor Type: %s\n"+
		"\tFamily: %s\n"+
		"\tManufacturer: %s\n"+
		"\tID: %d\n"+
		"\tVersion: %s\n"+
		"\tVoltage: %s\n"+
		"\tExternal Clock: %d\n"+
		"\tMax Speed: %d\n"+
		"\tCurrent Speed: %d\n"+
		"\tStatus: %s\n"+
		"\tUpgrade: %s\n"+
		"\tL1 Cache Handle: %d\n"+
		"\tL2 Cache Handle: %d\n"+
		"\tL3 Cache Handle: %d\n"+
		"\tSerial Number: %s\n"+
		"\tAsset Tag: %s\n"+
		"\tPart Number: %s\n"+
		"\tCore Count: %d\n"+
		"\tCore Enabled: %d\n"+
		"\tThread Count: %d\n"+
		"\tCharacteristics: %s\n"+
		"\tFamily2: %s",
		p.SocketDesignation,
		p.ProcessorType,
		p.Family,
		p.Manufacturer,
		p.ID,
		p.Version,
		p.Voltage,
		p.ExternalClock,
		p.MaxSpeed,
		p.CurrentSpeed,
		p.Status,
		p.Upgrade,
		p.L1CacheHandle,
		p.L2CacheHandle,
		p.L3CacheHandle,
		p.SerialNumber,
		p.AssetTag,
		p.PartNumber,
		p.CoreCount,
		p.CoreEnabled,
		p.ThreadCount,
		p.Characteristics,
		p.Family2)
}
