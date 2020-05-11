package port

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

// Information port信息
type Information struct {
	smbios.Header
	InternalReferenceDesignator string        `json:"internal_reference_designator,omitempty"`
	InternalConnectorType       ConnectorType `json:"internal_connector_type,omitempty"`
	ExternalReferenceDesignator string        `json:"external_reference_designator,omitempty"`
	ExternalConnectorType       ConnectorType `json:"external_connector_type,omitempty"`
	Type                        PortType      `json:"type,omitempty"`
}

func (p Information) String() string {
	return fmt.Sprintf("Port Information\n"+
		"\tInternal Reference Designator: %s\n"+
		"\tInternal Connector Type: %s\n"+
		"\tExternal Reference Designator: %s\n"+
		"\tExternal Connector Type: %s\n"+
		"\tType: %s",
		p.InternalReferenceDesignator,
		p.InternalConnectorType,
		p.ExternalReferenceDesignator,
		p.ExternalConnectorType,
		p.Type)
}
