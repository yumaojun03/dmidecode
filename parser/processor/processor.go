package processor

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

// Processor CPU信息
type Processor struct {
	smbios.Header
	SocketDesignation string                   `json:"socket_designation,omitempty"`
	ProcessorType     ProcessorType            `json:"processor_type,omitempty"`
	Family            ProcessorFamily          `json:"family,omitempty"`
	Manufacturer      string                   `json:"manufacturer,omitempty"`
	ID                ProcessorID              `json:"id,omitempty"`
	Version           string                   `json:"version,omitempty"`
	Voltage           ProcessorVoltage         `json:"voltage,omitempty"`
	ExternalClock     uint16                   `json:"external_clock,omitempty"`
	MaxSpeed          uint16                   `json:"max_speed,omitempty"`
	CurrentSpeed      uint16                   `json:"current_speed,omitempty"`
	Status            ProcessorStatus          `json:"status,omitempty"`
	Upgrade           ProcessorUpgrade         `json:"upgrade,omitempty"`
	L1CacheHandle     uint16                   `json:"l_1_cache_handle,omitempty"`
	L2CacheHandle     uint16                   `json:"l_2_cache_handle,omitempty"`
	L3CacheHandle     uint16                   `json:"l_3_cache_handle,omitempty"`
	SerialNumber      string                   `json:"serial_number,omitempty"`
	AssetTag          string                   `json:"asset_tag,omitempty"`
	PartNumber        string                   `json:"part_number,omitempty"`
	CoreCount         byte                     `json:"core_count,omitempty"`
	CoreEnabled       byte                     `json:"core_enabled,omitempty"`
	ThreadCount       byte                     `json:"thread_count,omitempty"`
	Characteristics   ProcessorCharacteristics `json:"characteristics,omitempty"`
	Family2           ProcessorFamily          `json:"family_2,omitempty"`
}

func (p Processor) String() string {
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
		"\tL1 Cache Handle: %x\n"+
		"\tL2 Cache Handle: %x\n"+
		"\tL3 Cache Handle: %x\n"+
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
