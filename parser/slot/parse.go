package slot

import (
	"github.com/yumaojun03/dmidecode/smbios"
)

// Parse 解析slot信息
func Parse(s *smbios.Structure) (*SystemSlot, error) {
	info := &SystemSlot{
		Designation:          s.GetString(0x0),
		Type:                 Type(s.GetByte(0x01)),
		DataBusWidth:         DataBusWidth(s.GetByte(0x02)),
		CurrentUsage:         Usage(s.GetByte(0x03)),
		Length:               Length(s.GetByte(0x04)),
		ID:                   ID(s.U16(0x05, 0x07)),
		Characteristics1:     Characteristics1(s.GetByte(0x07)),
		Characteristics2:     Characteristics2(s.GetByte(0x08)),
		SegmentGroupNumber:   SegmengGroupNumber(s.U16(0x09, 0x0b)),
		BusNumber:            Number(s.GetByte(0x0b)),
		DeviceFunctionNumber: Number(s.GetByte(0x0c)),
	}

	return info, nil

}
