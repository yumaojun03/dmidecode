package system

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

// Information 系统信息
type Information struct {
	smbios.Header
	Manufacturer string
	ProductName  string
	Version      string
	SerialNumber string
	UUID         string
	WakeUpType   WakeUpType
	SKUNumber    string
	Family       string
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
