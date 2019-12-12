package system

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

// Information 系统信息
type Information struct {
	smbios.Header
	Manufacturer string     `json:"manufacturer,omitempty"`
	ProductName  string     `json:"product_name,omitempty"`
	Version      string     `json:"version,omitempty"`
	SerialNumber string     `json:"serial_number,omitempty"`
	UUID         string     `json:"uuid,omitempty"`
	WakeUpType   WakeUpType `json:"wake_up_type,omitempty"`
	SKUNumber    string     `json:"sku_number,omitempty"`
	Family       string     `json:"family,omitempty"`
}

func (s Information) String() string {
	return fmt.Sprintf("Handle %x, DMI type %d, %d bytes\n"+
		"\tSystem Information\n"+
		"\tManufacturer: %s\n"+
		"\tProduct Name: %s\n"+
		"\tVersion: %s\n"+
		"\tSerial Number: %s\n"+
		"\tUUID: %s\n"+
		"\tWake-up Type: %s\n"+
		"\tSKU Number: %s\n"+
		"\tFamily: %s",
		s.Header.Handle,
		s.Header.Type,
		s.Header.Length,
		s.Manufacturer,
		s.ProductName,
		s.Version,
		s.SerialNumber,
		s.UUID,
		s.WakeUpType,
		s.SKUNumber,
		s.Family)
}
