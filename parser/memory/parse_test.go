package memory_test

import (
	"testing"

	"github.com/yumaojun03/dmidecode/parser/memory"
	"github.com/yumaojun03/dmidecode/smbios"

	"github.com/stretchr/testify/assert"
)

var (
	ds = &smbios.Structure{
		Header: smbios.Header{
			Type:   17,
			Length: 40,
			Handle: 4352,
		},
		Formatted: []byte{0x0, 0x10, 0xfe, 0xff, 0x48, 0x0, 0x40, 0x0, 0x0, 0x40,
			0x9, 0x1, 0x1, 0x0, 0x1a, 0x80, 0x20, 0x55, 0x8, 0x2,
			0x3, 0x4, 0x5, 0x2, 0x0, 0x0, 0x0, 0x0, 0x55, 0x8,
			0xb0, 0x4, 0xb0, 0x4, 0xb0, 0x4},
		Strings: []string{"A3", "00AD00B300AD", "23A47647", "00151430", "HMA42GR7MFR4N-TF"},
	}

	as = &smbios.Structure{
		Header: smbios.Header{
			Type:   16,
			Length: 23,
			Handle: 4096,
		},
		Formatted: []byte{0x3, 0x3, 0x6, 0x0, 0x0, 0x0, 0x30, 0xfe, 0xff, 0x18, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		Strings:   []string{},
	}
)

func TestParseDevice(t *testing.T) {
	bios, err := memory.ParseMemoryDevice(ds)
	if assert.NoError(t, err) {
		t.Log(bios)
	}
}

func TestParsePhysicalMemoryArray(t *testing.T) {
	bios, err := memory.ParseMemoryArray(as)
	if assert.NoError(t, err) {
		t.Log(bios)
	}
}
