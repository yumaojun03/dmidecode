package memory

import (
	"dmidecode/smbios"
	"fmt"
)

// PhysicalMemoryArray todo
type PhysicalMemoryArray struct {
	smbios.Header
	Location                PhysicalMemoryArrayLocation
	Use                     PhysicalMemoryArrayUse
	ErrorCorrection         PhysicalMemoryArrayErrorCorrection
	MaximumCapacity         uint32
	ErrorInformationHandle  uint16
	NumberOfMemoryDevices   uint16
	ExtendedMaximumCapacity uint64
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
