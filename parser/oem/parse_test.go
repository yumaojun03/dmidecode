package oem_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yumaojun03/dmidecode/parser/oem"
	"github.com/yumaojun03/dmidecode/smbios"
)

// *smbios.Structure: Header: Type: 11, Length: 5, Handle: 44, Data: [2] Strings: [[MS_VM_CERT/SHA1/27d66596a61c48dd3dc7216fd715126e33f59ae7] Welcome to the Virtual Machine]
var (
	s = &smbios.Structure{
		Header: smbios.Header{
			Type:   11,
			Length: 5,
			Handle: 44,
		},
		Formatted: []byte{0x2},
		Strings:   []string{"[MS_VM_CERT/SHA1/27d66596a61c48dd3dc7216fd715126e33f59ae7]", "Welcome to the Virtual Machine"},
	}
)

func TestParse(t *testing.T) {
	bios, err := oem.Parse(s)
	if assert.NoError(t, err) {
		t.Log(bios)
	}
}
