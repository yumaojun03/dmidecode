package tools

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func bcd(data []byte) int64 {
	var b int64
	l := len(data)
	if l > 8 {
		panic("bcd: Out of range")
	}
	// Number of 4-bits
	nb := int64(l * 2)
	for i := int64(0); i < nb; i++ {
		var shift uint64
		if i%2 == 0 {
			shift = 0
		} else {
			shift = 4
		}
		b += int64((data[i/2]>>shift)&0x0F) * int64(math.Pow10(int(i)))
	}
	return b
}

func U16(data []byte) uint16 {
	var u uint16
	binary.Read(bytes.NewBuffer(data[0:2]), binary.LittleEndian, &u)
	return u
}

func U32(data []byte) uint32 {
	var u uint32
	binary.Read(bytes.NewBuffer(data[0:4]), binary.LittleEndian, &u)
	return u
}

func u64(data []byte) uint64 {
	var u uint64
	binary.Read(bytes.NewBuffer(data[0:8]), binary.LittleEndian, &u)
	return u
}

func uuid(data []byte, ver string) string {
	if bytes.Index(data, []byte{0x00}) != -1 {
		return "Not present"
	}

	if bytes.Index(data, []byte{0xFF}) != -1 {
		return "Not settable"
	}

	if ver > "2.6" {
		return fmt.Sprintf("%02X%02X%02X%02X-%02X%02X-%02X%02X-%02X%02X-%02X%02X%02X%02X%02X%02X",
			data[3], data[2], data[1], data[0], data[5], data[4], data[7], data[6],
			data[8], data[9], data[10], data[11], data[12], data[13], data[14], data[15])
	}
	return fmt.Sprintf("%02X%02X%02X%02X-%02X%02X-%02X%02X-%02X%02X-%02X%02X%02X%02X%02X%02X",
		data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7],
		data[8], data[9], data[10], data[11], data[12], data[13], data[14], data[15])
}
