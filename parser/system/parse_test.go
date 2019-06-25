package system_test

import (
	"dmidecode/parser/system"
	"dmidecode/smbios"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	s = &smbios.Structure{
		Header: smbios.Header{
			Type:   1,
			Length: 27,
			Handle: 256,
		},
		Formatted: []byte{0x1, 0x2, 0x3, 0x4, 0xb2, 0xdd, 0x7d, 0x4, 0xa3, 0xeb, 0x21, 0xe7, 0x82, 0x22, 0xa6, 0x37, 0x62, 0x9b, 0x1b, 0xfa, 0x06, 0x0, 0x0},
		Strings:   []string{"Xen", "HVM domU", "4.7.2-2.2", "b2dd7d04-a3eb-21e7-8222-a637629b1bfa"},
	}
)

func TestParse(t *testing.T) {
	bios, err := system.Parse(s)
	if assert.NoError(t, err) {
		t.Log(bios)
	}
}
