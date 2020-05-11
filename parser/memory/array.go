package memory

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

// PhysicalMemoryArray todo
type PhysicalMemoryArray struct {
	smbios.Header
	Location                PhysicalMemoryArrayLocation        `json:"location,omitempty"`
	Use                     PhysicalMemoryArrayUse             `json:"use,omitempty"`
	ErrorCorrection         PhysicalMemoryArrayErrorCorrection `json:"error_correction,omitempty"`
	MaximumCapacity         uint32                             `json:"maximum_capacity,omitempty"`
	ErrorInformationHandle  uint16                             `json:"error_information_handle,omitempty"`
	NumberOfMemoryDevices   uint16                             `json:"number_of_memory_devices,omitempty"`
	ExtendedMaximumCapacity uint64                             `json:"extended_maximum_capacity,omitempty"`
}

func (p PhysicalMemoryArray) String() string {
	return fmt.Sprintf("Physcial Memory Array\n"+
		"\tLocation: %s\n"+
		"\tUse: %s\n"+
		"\tMemory Error Correction: %s\n"+
		"\tMaximum Capacity: %d\n"+
		"\tMemory Error Information Handle: %d\n"+
		"\tNumber of Memory Devices: %d\n"+
		"\tExtended Maximum Capacity: %d",
		p.Location,
		p.Use,
		p.ErrorCorrection,
		p.MaximumCapacity,
		p.ErrorInformationHandle,
		p.NumberOfMemoryDevices,
		p.ExtendedMaximumCapacity)
}
