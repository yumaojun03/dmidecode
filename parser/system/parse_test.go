package system_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yumaojun03/dmidecode/parser/system"
	"github.com/yumaojun03/dmidecode/smbios"
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

func TestParseUUIDNotPresent(t *testing.T) {
	s := &smbios.Structure{
		Header: smbios.Header{
			Type:   1,
			Length: 27,
			Handle: 256,
		},
		Formatted: []byte{0x1, 0x2, 0x3, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x06, 0x0, 0x0},
		Strings:   []string{"Xen", "HVM domU", "4.7.2-2.2", "b2dd7d04-a3eb-21e7-8222-a637629b1bfa"},
	}
	bios, err := system.Parse(s)
	if assert.NoError(t, err) {
		t.Log(bios)
	}
	if !assert.Equal(t, bios.UUID, "Not present") {
		t.Log("all 0 in uuid field should get not present")
		t.FailNow()
	}
}

func TestParseUUIDNotSettable(t *testing.T) {
	s := &smbios.Structure{
		Header: smbios.Header{
			Type:   1,
			Length: 27,
			Handle: 256,
		},
		Formatted: []byte{0x1, 0x2, 0x3, 0x4, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x06, 0x0, 0x0},
		Strings:   []string{"Xen", "HVM domU", "4.7.2-2.2", "b2dd7d04-a3eb-21e7-8222-a637629b1bfa"},
	}
	bios, err := system.Parse(s)
	if assert.NoError(t, err) {
		t.Log(bios)
	}
	if !assert.Equal(t, bios.UUID, "Settable") {
		t.Log("all 0xff in uuid field should be Not present but settable")
		t.FailNow()
	}
}

func TestParseUUIDWith0xff(t *testing.T) {
	s := &smbios.Structure{
		Header: smbios.Header{
			Type:   1,
			Length: 27,
			Handle: 256,
		},
		Formatted: []byte{0x1, 0x2, 0x3, 0x4, 0xff, 0xfe, 0xfd, 0xfc, 0xfb, 0xfa, 0xf9, 0xf8, 0xf7, 0xf6, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x06, 0x0, 0x0},
		Strings:   []string{"Xen", "HVM domU", "4.7.2-2.2", "b2dd7d04-a3eb-21e7-8222-a637629b1bfa"},
	}
	bios, err := system.Parse(s)
	if assert.NoError(t, err) {
		t.Log(bios)
	}
	if !assert.Equal(t, bios.UUID, "FCFDFEFF-FAFB-F8F9-F7F6-FFFFFFFFFFFF") {
		t.FailNow()
	}
}
