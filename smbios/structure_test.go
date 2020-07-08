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

func TestTypes(t *testing.T) {
	should := assert.New(t)

	should.Equal("System Boot", smbios.SystemBoot.String())
}

func TestGetByte(t *testing.T) {
	should := assert.New(t)

	should.Equal(s.GetByte(0x0b), uint8(0xff))
}

func TestGetBytes(t *testing.T) {
	should := assert.New(t)

	should.Equal(s.GetBytes(0x09, 0x0b), []uint8([]byte{0xff, 0xff}))
}

func TestString(t *testing.T) {
	should := assert.New(t)

	should.Equal(s.GetString(0x0), "PCIe Slot 1")
}

func TestU16(t *testing.T) {
	should := assert.New(t)

	should.Equal(s.U16(0x05, 0x07), uint16(0x1))
}

func TestU32(t *testing.T) {
	should := assert.New(t)

	should.Equal(s.U32(0x05, 0x09), uint32(0x1040001))
}

func TestU64(t *testing.T) {
	should := assert.New(t)

	should.Equal(s.U64(0x05, 0x0d), uint64(0x0))
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
