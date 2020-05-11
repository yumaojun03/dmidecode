package processor

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/smbios"
)

type Cache struct {
	smbios.Header
	SocketDesignation   string                   `json:"socket_designation,omitempty"`
	Configuration       CacheConfiguration       `json:"configuration,omitempty"`
	MaximumCacheSize    CacheSize                `json:"maximum_cache_size,omitempty"`
	InstalledSize       CacheSize                `json:"installed_size,omitempty"`
	SupportedSRAMType   CacheSRAMType            `json:"supported_sram_type,omitempty"`
	CurrentSRAMType     CacheSRAMType            `json:"current_sram_type,omitempty"`
	CacheSpeed          CacheSpeed               `json:"cache_speed,omitempty"`
	ErrorCorrectionType CacheErrorCorrectionType `json:"error_correction_type,omitempty"`
	SystemCacheType     CacheSystemCacheType     `json:"system_cache_type,omitempty"`
	Associativity       CacheAssociativity       `json:"associativity,omitempty"`
}

func (c Cache) String() string {
	return fmt.Sprintf("Cache Information\n"+
		"\tSocket Designation: %s\n"+
		"\tConfiguration: %s\n"+
		"\tMaximum Cache Size: %s\n"+
		"\tInstalled Size: %s\n"+
		"\tSupportedSRAM Type: %s\n"+
		"\tCurrentSRAM Type: %s\n"+
		"\tCache Speed: %d\n"+
		"\tError Correction Type: %s\n"+
		"\tSystem Cache Type: %s\n"+
		"\tAssociativity: %s",
		c.SocketDesignation,
		c.Configuration,
		c.MaximumCacheSize,
		c.InstalledSize,
		c.SupportedSRAMType,
		c.CurrentSRAMType,
		c.CacheSpeed,
		c.ErrorCorrectionType,
		c.SystemCacheType,
		c.Associativity)
}
