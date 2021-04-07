package oem

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

// Information 系统信息
type OEM struct {
	smbios.Header
	Strings []string `json:"strings,omitempty"`
}

func (s OEM) String() string {
	var strs = ""
	for _, s := range s.Strings {
		strs += "\t" + s + "\n"
	}
	return fmt.Sprintf("Handle %x, DMI type %d, %d bytes\n"+
		"OEM Strings\n"+strs,
		s.Header.Handle,
		s.Header.Type,
		s.Header.Length)
}
