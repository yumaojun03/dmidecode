<<<<<<< HEAD
package port

import (
	"github.com/yumaojun03/dmidecode/smbios"
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
=======
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
>>>>>>> fd60aa940a50299abe4322cdf257a0501b993f04
