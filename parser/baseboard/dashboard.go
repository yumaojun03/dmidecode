package baseboard

import (
	"dmidecode/smbios"
	"fmt"
)

// Information 主板信息
type Information struct {
	smbios.Header
	Manufacturer                   string
	ProductName                    string
	Version                        string
	SerialNumber                   string
	AssetTag                       string
	FeatureFlags                   FeatureFlags
	LocationInChassis              string
	ChassisHandle                  uint16
	BoardType                      Type
	NumberOfContainedObjectHandles byte
	ContainedObjectHandles         []byte
}

func (b Information) String() string {
	return fmt.Sprintf("Base Board Information\n"+
		"\tManufacturer: %s\n"+
		"\tProduct Name: %s\n"+
		"\tVersion: %s\n"+
		"\tSerial Number: %s\n"+
		"\tAsset Tag: %s\n"+
		"\tFeatures:%s\n"+
		"\tLocation In Chassis: %s\n"+
		"\tType: %s",
		b.Manufacturer,
		b.ProductName,
		b.Version,
		b.SerialNumber,
		b.AssetTag,
		b.FeatureFlags,
		b.LocationInChassis,
		b.BoardType)
}
