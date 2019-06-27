package slot

import "dmidecode/smbios"

// Parse 解析slot信息
func Parse(s *smbios.Structure) (*SystemSlot, error) {
	data := s.Formatted
	info := &SystemSlot{
		Designation:          s.GetString(0x0),
		Type:                 Type(data[0x01]),
		DataBusWidth:         DataBusWidth(data[0x02]),
		CurrentUsage:         Usage(data[0x03]),
		Length:               Length(data[0x04]),
		ID:                   ID(smbios.U16(data[0x05:0x07])),
		Characteristics1:     Characteristics1(data[0x07]),
		Characteristics2:     Characteristics2(data[0x08]),
		SegmentGroupNumber:   SegmengGroupNumber(smbios.U16(data[0x09:0x0B])),
		BusNumber:            Number(data[0x0B]),
		DeviceFunctionNumber: Number(data[0x0C]),
	}

	return info, nil

}
