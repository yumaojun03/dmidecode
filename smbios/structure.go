package smbios

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/digitalocean/go-smbios/smbios"
)

// A Header is a Structure's header.
type Header struct {
	Type   uint8
	Length uint8
	Handle uint16
}

func (h *Header) String() string {
	return fmt.Sprintf("Type: %d, Length: %d, Handle: %d",
		h.Type, h.Length, h.Handle)
}

// A Structure is an SMBIOS structure.
type Structure struct {
	Header    Header
	Formatted []byte
	Strings   []string

	formattedCount *int
}

// IsOverFlow 判断是否越界
func (s *Structure) IsOverFlow(length int) bool {
	return s.DataLength() >= length
}

// DataLength data长度
func (s *Structure) DataLength() int {
	return len(s.Formatted)
}

func (s *Structure) String() string {
	return fmt.Sprintf("Header: %s, Data: %v Strings: %v",
		s.Header.String(), s.Formatted, s.Strings)
}

// Type 协议Header
func (s *Structure) Type() uint8 {
	return s.Header.Type
}

// ReadStructures 读取smbios结构数据
func ReadStructures() ([]*Structure, error) {
	// Find SMBIOS data in operating system-specific location.
	rc, ep, err := smbios.Stream()
	if err != nil {
		return nil, fmt.Errorf("failed to open stream: %v", err)
	}
	// Be sure to close the stream!
	defer rc.Close()

	// Decode SMBIOS structures from the stream.
	d := smbios.NewDecoder(rc)
	ss, err := d.Decode()
	if err != nil {
		return nil, fmt.Errorf("failed to decode structures: %v", err)
	}

	// Determine SMBIOS version and table location from entry point.
	major, minor, rev := ep.Version()
	addr, size := ep.Table()

	fmt.Printf("SMBIOS %d.%d.%d - table: address: %#x, size: %d\n",
		major, minor, rev, addr, size)

	data := convertNative(ss)
	return data, nil
}

func convertNative(ss []*smbios.Structure) []*Structure {
	data := make([]*Structure, 0, len(ss))
	for _, s := range ss {
		ns := Structure{
			Header: Header{
				Type:   s.Header.Type,
				Length: s.Header.Length,
				Handle: s.Header.Handle,
			},
			Formatted: s.Formatted,
			Strings:   s.Strings,
		}

		data = append(data, &ns)
	}

	return data
}

// GetString todo
func (s *Structure) GetString(offset int) string {
	if offset > s.FormattedCount()-1 {
		return "Unknown"
	}
	index := s.Formatted[offset]

	if index == 0 {
		return "Unknown"
	}

	return s.Strings[index-1]
}

// FormattedCount 格式化数据队列长度, 避免索引出界
func (s *Structure) FormattedCount() int {
	if s.formattedCount == nil {
		c := len(s.Formatted)
		s.formattedCount = &c
	}

	return *s.formattedCount
}

// U16 bytes to uint16
func U16(data []byte) uint16 {
	var u uint16
	binary.Read(bytes.NewBuffer(data[0:2]), binary.LittleEndian, &u)
	return u
}

// U32 bytes to uint32
func U32(data []byte) uint32 {
	var u uint32
	binary.Read(bytes.NewBuffer(data[0:4]), binary.LittleEndian, &u)
	return u
}

// U64 bytes to uint64
func U64(data []byte) uint64 {
	var u uint64
	binary.Read(bytes.NewBuffer(data[0:8]), binary.LittleEndian, &u)
	return u
}
