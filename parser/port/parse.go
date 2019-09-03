package port

import (
	"dmidecode/smbios"
)

// Parse 解析port信息
func Parse(s *smbios.Structure) (*Information, error) {
	data := s.Formatted
	info := &Information{
		Header:                      s.Header,
		InternalReferenceDesignator: s.GetString(0x0),
		InternalConnectorType:       ConnectorType(data[0x01]),
		ExternalReferenceDesignator: s.GetString(0x2),
		ExternalConnectorType:       ConnectorType(data[0x03]),
		Type:                        PortType(data[0x04]),
	}
	return info, nil
}
