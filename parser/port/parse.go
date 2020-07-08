package port

import (
	"github.com/yumaojun03/dmidecode/smbios"
)

// Parse 解析port信息
func Parse(s *smbios.Structure) (info *Information, err error) {
	defer smbios.ParseRecovery(s, &err)

	info = &Information{
		Header:                      s.Header,
		InternalReferenceDesignator: s.GetString(0x0),
		InternalConnectorType:       ConnectorType(s.GetByte(0x01)),
		ExternalReferenceDesignator: s.GetString(0x2),
		ExternalConnectorType:       ConnectorType(s.GetByte(0x03)),
		Type:                        PortType(s.GetByte(0x04)),
	}
	return info, nil
}
