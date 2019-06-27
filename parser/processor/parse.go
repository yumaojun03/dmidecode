package processor

import (
	"dmidecode/smbios"
)

// ParseProcessor 处理器
func ParseProcessor(s *smbios.Structure) (*Processor, error) {
	data := s.Formatted
	p := &Processor{
		SocketDesignation: s.GetString(0x0),
		ProcessorType:     ProcessorType(data[0x01]),
		Family:            ProcessorFamily(data[0x02]),
		Manufacturer:      s.GetString(0x3),
		// TODO:
		//pi.ProcessorID
		Version:         s.GetString(0xc),
		Voltage:         ProcessorVoltage(data[0xD]),
		ExternalClock:   smbios.U16(data[0xE:0x10]),
		MaxSpeed:        smbios.U16(data[0x10:0x12]),
		CurrentSpeed:    smbios.U16(data[0x12:0x14]),
		Status:          ProcessorStatus(data[0x14]),
		Upgrade:         ProcessorUpgrade(data[0x15]),
		L1CacheHandle:   smbios.U16(data[0x16:0x18]),
		L2CacheHandle:   smbios.U16(data[0x18:0x1E]),
		L3CacheHandle:   smbios.U16(data[0x1E:0x20]),
		SerialNumber:    s.GetString(0x1c),
		AssetTag:        s.GetString(0x1d),
		PartNumber:      s.GetString(0x1e),
		CoreCount:       data[0x1F],
		CoreEnabled:     data[0x20],
		ThreadCount:     data[0x21],
		Characteristics: ProcessorCharacteristics(smbios.U16(data[0x22:0x24])),
		Family2:         ProcessorFamily(data[0x24]),
	}

	return p, nil
}

// ParseCache 缓存信息
func ParseCache(s *smbios.Structure) (*Cache, error) {
	data := s.Formatted

	cache := &Cache{
		SocketDesignation:   s.GetString(0x0),
		Configuration:       NewCacheConfiguration(smbios.U16(data[0x01:0x03])),
		MaximumCacheSize:    NewCacheSize(smbios.U16(data[0x03:0x05])),
		InstalledSize:       NewCacheSize(smbios.U16(data[0x05:0x07])),
		SupportedSRAMType:   CacheSRAMType(smbios.U16(data[0x07:0x09])),
		CurrentSRAMType:     CacheSRAMType(smbios.U16(data[0x09:0x0B])),
		CacheSpeed:          CacheSpeed(data[0x0B]),
		ErrorCorrectionType: CacheErrorCorrectionType(data[0xC]),
		SystemCacheType:     CacheSystemCacheType(data[0xD]),
		Associativity:       CacheAssociativity(data[0xE]),
	}

	return cache, nil
}
