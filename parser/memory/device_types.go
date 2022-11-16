package memory

import (
	"strings"
)

type MemoryDeviceFormFactor byte

const (
	MemoryDeviceFormFactorOther MemoryDeviceFormFactor = 1 + iota
	MemoryDeviceFormFactorUnknown
	MemoryDeviceFormFactorSIMM
	MemoryDeviceFormFactorSIP
	MemoryDeviceFormFactorChip
	MemoryDeviceFormFactorDIP
	MemoryDeviceFormFactorZIP
	MemoryDeviceFormFactorProprietaryCard
	MemoryDeviceFormFactorDIMM
	MemoryDeviceFormFactorTSOP
	MemoryDeviceFormFactorRowofchips
	MemoryDeviceFormFactorRIMM
	MemoryDeviceFormFactorSODIMM
	MemoryDeviceFormFactorSRIMM
	MemoryDeviceFormFactorFB_DIMM
)

func (m MemoryDeviceFormFactor) String() string {
	factors := [...]string{
		"Other",
		"Unknown",
		"SIMM",
		"SIP",
		"Chip",
		"DIP",
		"ZIP",
		"Proprietary Card",
		"DIMM",
		"TSOP",
		"Row of chips",
		"RIMM",
		"SODIMM",
		"SRIMM",
		"FB-DIMM",
	}
	return factors[m-1]
}

type MemoryDeviceType byte

const (
	MemoryDeviceTypeOther MemoryDeviceType = 1 + iota
	MemoryDeviceTypeUnknown
	MemoryDeviceTypeDRAM
	MemoryDeviceTypeEDRAM
	MemoryDeviceTypeVRAM
	MemoryDeviceTypeSRAM
	MemoryDeviceTypeRAM
	MemoryDeviceTypeROM
	MemoryDeviceTypeFLASH
	MemoryDeviceTypeEEPROM
	MemoryDeviceTypeFEPROM
	MemoryDeviceTypeEPROM
	MemoryDeviceTypeCDRAM
	MemoryDeviceType3DRAM
	MemoryDeviceTypeSDRAM
	MemoryDeviceTypeSGRAM
	MemoryDeviceTypeRDRAM
	MemoryDeviceTypeDDR
	MemoryDeviceTypeDDR2
	MemoryDeviceTypeDDR2FB_DIMM
	_
	_
	_
	MemoryDeviceTypeDDR3
	MemoryDeviceTypeFBD2
	MemoryDeviceTypeDDR4
	MemoryDeviceTypeLPDDR
	MemoryDeviceTypeLPDDR2
	MemoryDeviceTypeLPDDR3
	MemoryDeviceTypeLPDDR4
)

func (m MemoryDeviceType) String() string {
	types := [...]string{
		"",
		"Other",
		"Unknown",
		"DRAM",
		"EDRAM",
		"VRAM",
		"SRAM",
		"RAM",
		"ROM",
		"FLASH",
		"EEPROM",
		"FEPROM",
		"EPROM",
		"CDRAM",
		"3DRAM",
		"SDRAM",
		"SGRAM",
		"RDRAM",
		"DDR",
		"DDR2",
		"DDR2 FB-DIMM",
		"Reserved1",
		"Reserved2",
		"Reserved3",
		"DDR3",
		"FBD2",
		"DDR4",
		"LPDDR",
		"LPDDR2",
		"LPDDR3",
		"LPDDR4",
		"Logical non-volatile device",
		"HBM (High Bandwidth Memory)",
		"HBM2 (High Bandwidth Memory Generation 2)",
		"DDR5",
		"LPDDR5",
		"HBM3 (High Bandwidth Memory Generation 3)",
	}
	if len(types) <= int(m) {
		return "Unknown"
	}
	return types[m]
}

// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37..32567, 32568..65535}
// Values={Unknown, Other, DRAM, Synchronous DRAM, Cache DRAM, EDO, EDRAM, VRAM, SRAM, RAM, ROM, Flash, EEPROM, FEPROM, EPROM, CDRAM, 3DRAM, SDRAM, SGRAM, RDRAM, DDR, DDR-2, BRAM, FB-DIMM, DDR3, FBD2, DDR4, LPDDR, LPDDR2, LPDDR3, LPDDR4, Logical non-volatile device, HBM (High Bandwidth Memory), HBM2 (High Bandwidth Memory Generation 2), DDR5, LPDDR5, HBM3 (High Bandwidth Memory Generation 3), DMTF Reserved, Vendor Reserved}
type MemoryDeviceTypeDetail uint16

const (
	MemoryDeviceTypeDetailReserved MemoryDeviceTypeDetail = 1 + iota
	MemoryDeviceTypeDetailOther
	MemoryDeviceTypeDetailUnknown
	MemoryDeviceTypeDetailFast_paged
	MemoryDeviceTypeDetailStaticcolumn
	MemoryDeviceTypeDetailPseudo_static
	MemoryDeviceTypeDetailRAMBUS
	MemoryDeviceTypeDetailSynchronous
	MemoryDeviceTypeDetailCMOS
	MemoryDeviceTypeDetailEDO
	MemoryDeviceTypeDetailWindowDRAM
	MemoryDeviceTypeDetailCacheDRAM
	MemoryDeviceTypeDetailNon_volatile
	MemoryDeviceTypeDetailRegisteredBuffered
	MemoryDeviceTypeDetailUnbufferedUnregistered
	MemoryDeviceTypeDetailLRDIMM
)

func (m MemoryDeviceTypeDetail) String() string {
	details := [...]string{
		"Reserved",
		"Other",
		"Unknown",
		"Fast-paged",
		"Static column",
		"Pseudo-static",
		"RAMBUS",
		"Synchronous",
		"CMOS",
		"EDO",
		"Window DRAM",
		"Cache DRAM",
		"Non-volatile",
		"Registered (Buffered)",
		"Unbuffered (Unregistered)",
		"LRDIMM",
	}

	d := []string{}
	for i := 1; i <= 16; i++ {
		if m&(1<<uint(i)) != 0 {
			d = append(d, details[i])
		}
	}

	return strings.Join(d, ",")
}
