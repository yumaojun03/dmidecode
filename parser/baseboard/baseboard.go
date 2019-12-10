package baseboard

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

// Information 主板信息
type Information struct {
	smbios.Header
	Manufacturer                   string       `json:"manufacturer,omitempty"`
	ProductName                    string       `json:"product_name,omitempty"`
	Version                        string       `json:"version,omitempty"`
	SerialNumber                   string       `json:"serial_number,omitempty"`
	AssetTag                       string       `json:"asset_tag,omitempty"`
	FeatureFlags                   FeatureFlags `json:"feature_flags,omitempty"`
	LocationInChassis              string       `json:"location_in_chassis,omitempty"`
	ChassisHandle                  uint16       `json:"chassis_handle,omitempty"`
	BoardType                      Type         `json:"board_type,omitempty"`
	NumberOfContainedObjectHandles byte         `json:"number_of_contained_object_handles,omitempty"`
	ContainedObjectHandles         []byte       `json:"contained_object_handles,omitempty"`
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
