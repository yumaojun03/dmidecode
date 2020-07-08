package smbios_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yumaojun03/dmidecode/smbios"
)

var (
	s = &smbios.Structure{
		Header: smbios.Header{
			Type:   9,
			Length: 17,
			Handle: 2304,
		},
		Formatted: []byte{0x1, 0xb1, 0xd, 0x3, 0x4, 0x1, 0x0, 0x4, 0x1, 0xff, 0xff, 0xff, 0xff},
		Strings:   []string{"PCIe Slot 1"},
	}
)

func TestRead(t *testing.T) {
	_, ss, err := smbios.ReadStructures()
	t.Log(ss, err)

}

func TestParsePanice(t *testing.T) {
	should := assert.New(t)

	err := mockParse(s)
	should.Equal(err.Error(), "parse structure (Header: Type: 9, Length: 17, Handle: 2304, Data: [1 177 13 3 4 1 0 4 1 255 255 255 255] Strings: [PCIe Slot 1])  pannic: parse panic test")
}

func mockParse(*smbios.Structure) (err error) {
	defer smbios.ParseRecovery(s, &err)
	panic("parse panic test")
}
