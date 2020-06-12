package onboard_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yumaojun03/dmidecode/parser/onboard"
	"github.com/yumaojun03/dmidecode/smbios"
)

var (
	s41 = &smbios.Structure{
		Header: smbios.Header{
			Type:   41,
			Length: 11,
			Handle: 10498,
		},
		Formatted: []byte{0x1, 0x85, 0x3, 0x0, 0x0, 0x2, 0x0},
		Strings:   []string{"Integrated NIC 3"},
	}

	s10 = &smbios.Structure{
		Header: smbios.Header{
			Type:   10,
			Length: 11,
			Handle: 10498,
		},
		Formatted: []byte{0x85, 0x1, 0x85, 0x2, 0x85, 0x3},
		Strings:   []string{"d1", "d2", "d3"},
	}
)

func TestParseType41(t *testing.T) {
	bios, err := onboard.ParseType41(s41)
	if assert.NoError(t, err) {
		t.Log(bios)
	}
}

func TestParseType10(t *testing.T) {
	bios, err := onboard.ParseType10(s10)
	if assert.NoError(t, err) {
		t.Log(bios)
	}
}
