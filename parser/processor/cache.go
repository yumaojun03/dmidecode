package processor

import (
	"dmidecode/smbios"
	"fmt"
)

type Cache struct {
	smbios.Header
	SocketDesignation   string
	Configuration       CacheConfiguration
	MaximumCacheSize    CacheSize
	InstalledSize       CacheSize
	SupportedSRAMType   CacheSRAMType
	CurrentSRAMType     CacheSRAMType
	CacheSpeed          CacheSpeed
	ErrorCorrectionType CacheErrorCorrectionType
	SystemCacheType     CacheSystemCacheType
	Associativity       CacheAssociativity
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
