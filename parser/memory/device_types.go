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
	}
	return types[m-1]
}

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
