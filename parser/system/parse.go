package system

import (
	"bytes"
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

// Parse 解析smbios struct数据
func Parse(s *smbios.Structure) (info *Information, err error) {
	defer smbios.ParseRecovery(s, &err)

	info = &Information{
		Header:       s.Header,
		Manufacturer: s.GetString(0x0),
		ProductName:  s.GetString(0x1),
		Version:      s.GetString(0x2),
		SerialNumber: s.GetString(0x3),
		UUID:         uuid(s.GetBytes(0x04, 0x14), s.GetString(2)),
		WakeUpType:   WakeUpType(s.GetByte(0x14)),
		SKUNumber:    s.GetString(0x15),
		Family:       s.GetString(0x16),
	}

	return info, nil
}

// UUID 主板uuid
// If the value is all FFh, the ID is not currently present in the system,
// but it can be set. If the value is all 00h, the ID is not present in the system.
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
