package battery

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

type Information struct {
	smbios.Header
	Location                 string          `json:"location,omitempty"`
	Manufacturer             string          `json:"manufacturer,omitempty"`
	ManufacturerDate         string          `json:"manufacturer_date,omitempty"`
	SerialNumber             string          `json:"serial_number,omitempty"`
	DeviceName               string          `json:"device_name,omitempty"`
	DeviceChemistry          DeviceChemistry `json:"device_chemistry,omitempty"`
	DesignCapacity           uint16          `json:"design_chemistry,omitempty"`
	DesignVoltage            uint16          `json:"design_voltage,omitempty"`
	SBDSVersionNumber        string          `json:"sbds_version_number,omitempty"`
	MaximumError             uint16          `json:"maximum_error,omitempty"`
	SBDSSerialNumber         uint16          `json:"sbds_serial_number,omitempty"`
	SBDSManufactureDate      uint16          `json:"sbds_manufacture_date,omitempty"`
	SBDSDeviceChemistry      string          `json:"sbds_device_chemistry,omitempty"`
	DesignCapacityMultiplier uint16          `json:"design_capacity_multiplier,omitempty"`
	OEMSpecific              uint16          `json:"oem_specific,omitempty"`
}

func (b Information) String() string {
	s := fmt.Sprintf("Handle %x, DMI type %d, %d bytes\n"+
		"Portable Battery\n"+
		"\tLocation: %s\n"+
		"\tManufacturer: %s\n"+
		"\tManufacturer Date: %s\n"+
		"\tSerial Number: %s\n"+
		"\tDevice Name: %s\n"+
		"\tDevice Chemistry: %s\n"+
		"\tDesign Capacity: %d\n"+
		"\tDesign Voltage: %d\n"+
		"\tSBDS Version Number: %s\n"+
		"\tMaximum Error: %d\n"+
		"\tSBDS Serial Number: %d\n"+
		"\tSBDS Manufacture Date: %d\n"+
		"\tSBDS Device Chemistry: %s\n"+
		"\tDesign Capacity Multiplier: %d\n"+
		"\tOEM-specific Information: %d\n",
		b.Header.Handle,
		b.Header.Type,
		b.Header.Length,
		b.Location,
		b.Manufacturer,
		b.ManufacturerDate,
		b.SerialNumber,
		b.DeviceName,
		b.DeviceChemistry,
		b.DesignCapacity,
		b.DesignVoltage,
		b.SBDSVersionNumber,
		b.MaximumError,
		b.SBDSSerialNumber,
		b.SBDSManufactureDate,
		b.SBDSDeviceChemistry,
		b.DesignCapacityMultiplier,
		b.OEMSpecific)

	return s
}
