package port_test

import (
	"testing"

	"github.com/yumaojun03/dmidecode/parser/port"
	"github.com/yumaojun03/dmidecode/smbios"

	"github.com/stretchr/testify/assert"
)

var (
	s = &smbios.Structure{
		Header: smbios.Header{
			Type:   8,
			Length: 9,
			Handle: 2048,
		},
		Formatted: []byte{0x0, 0x0, 0x1, 0x12, 0x10},
		Strings:   []string{"Back USB port 1"},
	}
)

func TestParse(t *testing.T) {
	bios, err := port.Parse(s)
	if assert.NoError(t, err) {
		t.Log(bios)
	}
}
