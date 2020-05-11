package slot

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

// SystemSlot todo
type SystemSlot struct {
	smbios.Header
	Designation          string             `json:"designation,omitempty"`
	Type                 Type               `json:"type,omitempty"`
	DataBusWidth         DataBusWidth       `json:"data_bus_width,omitempty"`
	CurrentUsage         Usage              `json:"current_usage,omitempty"`
	Length               Length             `json:"length,omitempty"`
	ID                   ID                 `json:"id,omitempty"`
	Characteristics1     Characteristics1   `json:"characteristics_1,omitempty"`
	Characteristics2     Characteristics2   `json:"characteristics_2,omitempty"`
	SegmentGroupNumber   SegmengGroupNumber `json:"segment_group_number,omitempty"`
	BusNumber            Number             `json:"bus_number,omitempty"`
	DeviceFunctionNumber Number             `json:"device_function_number,omitempty"`
}

func (s SystemSlot) String() string {
	return fmt.Sprintf("System Slot %d\n"+
		"\tSlot Designation: %s\n"+
		"\tSlot Type: %s\n"+
		"\tSlot Data Bus Width: %s\n"+
		"\tCurrent Usage: %s\n"+
		"\tSlot Length: %s\n"+
		"\tSlot ID: %d\n"+
		"\tSlot Characteristics1: %s\n"+
		"\tSlot Characteristics2: %s\n"+
		"\tSegment Group Number: %x\n"+
		"\tBus Number: %x\n"+
		"\tDevice/Function Number: %x",
		s.ID,
		s.Designation,
		s.Type,
		s.DataBusWidth,
		s.CurrentUsage,
		s.Length,
		s.ID,
		s.Characteristics1,
		s.Characteristics2,
		s.SegmentGroupNumber,
		s.BusNumber,
		s.DeviceFunctionNumber)
}
