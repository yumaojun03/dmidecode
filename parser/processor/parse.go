package processor

import (
	"github.com/yumaojun03/dmidecode/smbios"
)

// ParseProcessor 处理器
// 7.5.3 Processor ID field format
// 		The Processor ID field contains processor-specific information that describes the processor’s features.
// 7.5.3.1 x86-class CPUs
// 		For x86 class CPUs, the field’s format depends on the processor’s support of the CPUID instruction. If the
// 		instruction is supported, the Processor ID field contains two DWORD-formatted values. The first (offsets
// 		08h-0Bh) is the EAX value returned by a CPUID instruction with input EAX set to 1; the second (offsets
// 		0Ch-0Fh) is the EDX value returned by that instruction.
// 		Otherwise, only the first two bytes of the Processor ID field are significant (all others are set to 0) and
// 		contain (in WORD-format) the contents of the DX register at CPU reset.
// 7.5.3.2 ARM32-class CPUs
// 		For ARM32-class CPUs, the Processor ID field contains two DWORD-formatted values. The first (offsets
// 		08h-0Bh) is the contents of the Main ID Register (MIDR); the second (offsets 0Ch-0Fh) is zero.
// 7.5.3.3 ARM64-class CPUs
// 		For ARM64-class CPUs, the Processor ID field contains two DWORD-formatted values. Th
func ParseProcessor(s *smbios.Structure) (info *Processor, err error) {
	defer smbios.ParseRecovery(s, &err)

	info = &Processor{
		SocketDesignation: s.GetString(0x0),
		ProcessorType:     ProcessorType(s.GetByte(0x01)),
		Family:            ProcessorFamily(s.GetByte(0x02)),
		Manufacturer:      s.GetString(0x3),
		ID:                ProcessorID(s.GetBytes(0x4, 0xc)),
		Version:           s.GetString(0xc),
		Voltage:           ProcessorVoltage(s.GetByte(0xD)),
		ExternalClock:     s.U16(0xE, 0x10),
		MaxSpeed:          s.U16(0x10, 0x12),
		CurrentSpeed:      s.U16(0x12, 0x14),
		Status:            ProcessorStatus(s.GetByte(0x14)),
		Upgrade:           ProcessorUpgrade(s.GetByte(0x15)),
		L1CacheHandle:     s.U16(0x16, 0x18),
		L2CacheHandle:     s.U16(0x18, 0x1E),
		L3CacheHandle:     s.U16(0x1E, 0x20),
		SerialNumber:      s.GetString(0x1c),
		AssetTag:          s.GetString(0x1d),
		PartNumber:        s.GetString(0x1e),
		CoreCount:         s.GetByte(0x1F),
		CoreEnabled:       s.GetByte(0x20),
		ThreadCount:       s.GetByte(0x21),
		Characteristics:   ProcessorCharacteristics(s.U16(0x22, 0x24)),
		Family2:           ProcessorFamily(s.GetByte(0x24)),
	}

	return info, nil
}

// ParseCache 缓存信息
func ParseCache(s *smbios.Structure) (info *Cache, err error) {
	defer smbios.ParseRecovery(s, &err)

	info = &Cache{
		SocketDesignation:   s.GetString(0x0),
		Configuration:       NewCacheConfiguration(s.U16(0x01, 0x03)),
		MaximumCacheSize:    NewCacheSize(s.U16(0x03, 0x05)),
		InstalledSize:       NewCacheSize(s.U16(0x05, 0x07)),
		SupportedSRAMType:   CacheSRAMType(s.U16(0x07, 0x09)),
		CurrentSRAMType:     CacheSRAMType(s.U16(0x09, 0x0b)),
		CacheSpeed:          CacheSpeed(s.GetByte(0x0b)),
		ErrorCorrectionType: CacheErrorCorrectionType(s.GetByte(0xc)),
		SystemCacheType:     CacheSystemCacheType(s.GetByte(0xd)),
		Associativity:       CacheAssociativity(s.GetByte(0xe)),
	}

	return info, nil
}
