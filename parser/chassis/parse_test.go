package chassis_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yumaojun03/dmidecode/parser/chassis"
	"github.com/yumaojun03/dmidecode/smbios"
)

var (
	s = &smbios.Structure{
		Header: smbios.Header{
			Type:   3,
			Length: 22,
			Handle: 768,
		},
		Formatted: []byte{0x1, 0x97, 0x0, 0x2, 0x0, 0x3, 0x3, 0x3, 0x2, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0},
		Strings:   []string{"Dell Inc.", "HVSRF52"},
	}
)

func TestParse(t *testing.T) {
	bios, err := chassis.Parse(s)
	if assert.NoError(t, err) {
		t.Log(bios)
	}
}
