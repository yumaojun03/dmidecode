package bios_test

import (
	"dmidecode/parser/bios"
	"dmidecode/smbios"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	s = &smbios.Structure{
		Header: smbios.Header{
			Type:   0,
			Length: 24,
			Handle: 0,
		},
		Formatted: []byte{0x1, 0x2, 0x0, 0xe8, 0x3, 0x0, 0x80, 0x0, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0x4, 0x7, 0xff, 0xff},
		Strings:   []string{"Xen", "4.7.2-2.2", "05/11/2017"},
	}
)

func TestParse(t *testing.T) {
	bios, err := bios.Parse(s)
	if assert.NoError(t, err) {
		t.Log(bios)
	}
}
