package baseboard_test

import (
	"dmidecode/parser/baseboard"
	"dmidecode/smbios"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	s = &smbios.Structure{
		Header: smbios.Header{
			Type:   2,
			Length: 8,
			Handle: 512,
		},
		Formatted: []byte{0x1, 0x2, 0x3, 0x4},
		Strings:   []string{"Dell Inc.", "0CNCJW", "A05", ".HVSRF52.CN7475154A0700."},
	}
)

func TestParse(t *testing.T) {
	bios, err := baseboard.Parse(s)
	if assert.NoError(t, err) {
		t.Log(bios)
	}
}
