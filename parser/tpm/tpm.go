package tpm

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

// Information tpm信息
type Information struct {
	smbios.Header
	VendorID string     `json:"vendor_id,omitempty"`
	SpecificationVersion  string     `json:"specification_version,omitempty"`
	FirmwareRevision      string     `json:"firmware_revision,omitempty"`
	Description string     `json:"description,omitempty"`
	Characteristics         string     `json:"characteristics,omitempty"`
	OEMSpecificInformation string `json:"oem_specific_information,omitempty"`
}

func (s Information) String() string {
	return fmt.Sprintf("Handle %x, DMI type %d, %d bytes\n"+
	"\tTPM Device\n"+
		"\tVendor ID: %s\n"+
		"\tSpecification Version: %s\n"+
		"\tFirmware Revision: %s\n"+
		"\tDescription: %s\n"+
		"\tCharacteristics: \n%s\n"+
		"\tOEM-specific Information: %s\n",
		s.Header.Handle,
		s.Header.Type,
		s.Header.Length,
		s.VendorID,
		s.SpecificationVersion,
		s.FirmwareRevision,
		s.Description,
		s.Characteristics,
		s.OEMSpecificInformation,
		)
}