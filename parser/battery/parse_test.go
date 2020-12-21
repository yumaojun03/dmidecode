package battery_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yumaojun03/dmidecode/parser/battery"
	"github.com/yumaojun03/dmidecode/smbios"
)

var (
	s = &smbios.Structure{
		Header: smbios.Header{
			Type:   22,
			Length: 26,
			Handle: 44,
		},
		Formatted: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00},
		Strings:   []string{"Fake", "-Virtual Battery 0-", "08/08/2010", "Battery 0", "CRB Battery 0", "LithiumPolymer"},
	}
)

func TestParse(t *testing.T) {
	should := assert.New(t)

	battery, err := battery.Parse(s)
	if assert.NoError(t, err) {
		should.Equal("Fake", battery.Location)
		should.Equal("-Virtual Battery 0-", battery.Manufacturer)
		should.Equal("08/08/2010", battery.ManufacturerDate)
		should.Equal("Battery 0", battery.SerialNumber)
		should.Equal("CRB Battery 0", battery.DeviceName)
		should.Equal("Lithium Polymer", battery.DeviceChemistry.String())
		t.Log(battery)
	}
}
