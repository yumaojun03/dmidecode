package processor_test

import (
	"testing"

	"github.com/yumaojun03/dmidecode/parser/processor"
	"github.com/yumaojun03/dmidecode/smbios"

	"github.com/stretchr/testify/assert"
)

var (
	ps = &smbios.Structure{
		Header: smbios.Header{
			Type:   4,
			Length: 42,
			Handle: 1024,
		},
		Formatted: []byte{0x1, 0x3, 0xb3, 0x2, 0xf2, 0x6, 0x3, 0x0, 0xff, 0xfb,
			0xeb, 0xbf, 0x3, 0x8d, 0x80, 0x25, 0xa0, 0xf, 0xfc, 0x8,
			0x41, 0x2b, 0x0, 0x7, 0x1, 0x7, 0x2, 0x7, 0x0, 0x0,
			0x0, 0xa, 0xa, 0x14, 0xfc, 0x0, 0xb3, 0x0},
		Strings: []string{"CPU1", "Intel", "Intel(R) Xeon(R) CPU E5-2650 v3 @ 2.30GHz"},
	}

	cs = &smbios.Structure{
		Header: smbios.Header{
			Type:   7,
			Length: 19,
			Handle: 1796,
		},
		Formatted: []byte{0x0, 0x81, 0x1, 0x0, 0xa, 0x0, 0xa, 0x2, 0x0, 0x2, 0x0, 0x0, 0x5, 0x5, 0x7},
		Strings:   []string{},
	}
)

func TestParseProcessor(t *testing.T) {
	bios, err := processor.ParseProcessor(ps)
	if assert.NoError(t, err) {
		t.Log(bios)
	}
}

func TestParseCache(t *testing.T) {
	bios, err := processor.ParseCache(cs)
	if assert.NoError(t, err) {
		t.Log(bios)
	}
}
