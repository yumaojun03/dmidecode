package processor

import (
	"fmt"
	"strings"
)

type ProcessorType byte

const (
	ProcessorTypeOther ProcessorType = 1 + iota
	ProcessorTypeUnknown
	ProcessorTypeCentralProcessor
	ProcessorTypeMathProcessor
	ProcessorTypeDSPProcessor
	ProcessorTypeVideoProcessor
)

func (p ProcessorType) String() string {
	types := [...]string{
		"Other",
		"Unknown",
		"CentralProcessor",
		"MathProcessor",
		"DSPProcessor",
		"VideoProcessor",
	}
	return types[p-1]
}

type ProcessorFamily uint16

const (
	ProcessorOther ProcessorFamily = 1 + iota
	ProcessorUnknown
	ProcessorProcessorFamily8086
	ProcessorProcessorFamily80286
	ProcessorIntel386TMprocessor
	ProcessorIntel486TMprocessor
	ProcessorProcessorFamily8087
	ProcessorProcessorFamily80287
	ProcessorProcessorFamily80387
	ProcessorProcessorFamily80487
	ProcessorIntelPentiumprocessor
	ProcessorPentiumProprocessor
	ProcessorPentiumIIprocessor
	ProcessorPentiumprocessorwithMMXTMtechnology
	ProcessorIntelCeleronprocessor
	ProcessorPentiumIIXeonTMprocessor
	ProcessorPentiumIIIprocessor
	ProcessorM1Family
	ProcessorM2Family
	ProcessorIntelCeleronMprocessor
	ProcessorIntelPentium4HTprocessor
	_
	_
	ProcessorAMDDuronTMProcessorFamily
	ProcessorK5Family
	ProcessorK6Family
	ProcessorK6_2
	ProcessorK6_3
	ProcessorAMDAthlonTMProcessorFamily
	ProcessorAMD29000Family
	ProcessorK6_2Plus
	ProcessorPowerPCFamily
	ProcessorPowerPC601
	ProcessorPowerPC603
	ProcessorPowerPC603Plus
	ProcessorPowerPC604
	ProcessorPowerPC620
	ProcessorPowerPCx704
	ProcessorPowerPC750
	ProcessorIntelCoreTMDuoprocessor
	ProcessorIntelCoreTMDuomobileprocessor
	ProcessorIntelCoreTMSolomobileprocessor
	ProcessorIntelAtomTMprocessor
	_
	_
	_
	_
	ProcessorAlphaFamily
	ProcessorAlpha21064
	ProcessorAlpha21066
	ProcessorAlpha21164
	ProcessorAlpha21164PC
	ProcessorAlpha21164a
	ProcessorAlpha21264
	ProcessorAlpha21364
	ProcessorAMDTurionTMIIUltraDual_CoreMobileMProcessorFamily
	ProcessorAMDTurionTMIIDual_CoreMobileMProcessorFamily
	ProcessorAMDAthlonTMIIDual_CoreMProcessorFamily
	ProcessorAMDOpteronTM6100SeriesProcessor
	ProcessorAMDOpteronTM4100SeriesProcessor
	ProcessorAMDOpteronTM6200SeriesProcessor
	ProcessorAMDOpteronTM4200SeriesProcessor
	ProcessorAMDFXTMSeriesProcessor
	ProcessorMIPSFamily
	ProcessorMIPSR4000
	ProcessorMIPSR4200
	ProcessorMIPSR4400
	ProcessorMIPSR4600
	ProcessorMIPSR10000
	ProcessorAMDC_SeriesProcessor
	ProcessorAMDE_SeriesProcessor
	ProcessorAMDA_SeriesProcessor
	ProcessorAMDG_SeriesProcessor
	ProcessorAMDZ_SeriesProcessor
	ProcessorAMDR_SeriesProcessor
	ProcessorAMDOpteronTM4300SeriesProcessor
	ProcessorAMDOpteronTM6300SeriesProcessor
	ProcessorAMDOpteronTM3300SeriesProcessor
	ProcessorAMDFireProTMSeriesProcessor
	ProcessorSPARCFamily
	ProcessorSuperSPARC
	ProcessormicroSPARCII
	ProcessormicroSPARCIIep
	ProcessorUltraSPARC
	ProcessorUltraSPARCII
	ProcessorUltraSPARCIii
	ProcessorUltraSPARCIII
	ProcessorUltraSPARCIIIi
	_
	_
	_
	_
	_
	_
	_
	Processor68040Family
	Processor68xxx
	ProcessorProcessorFamily68000
	ProcessorProcessorFamily68010
	ProcessorProcessorFamily68020
	ProcessorProcessorFamily68030
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	ProcessorHobbitFamily
	_
	_
	_
	_
	_
	_
	_
	ProcessorCrusoeTMTM5000Family
	ProcessorCrusoeTMTM3000Family
	ProcessorEfficeonTMTM8000Family
	_
	_
	_
	_
	_
	ProcessorWeitek
	_
	ProcessorItaniumTMprocessor
	ProcessorAMDAthlonTM64ProcessorFamily
	ProcessorAMDOpteronTMProcessorFamily
	ProcessorAMDSempronTMProcessorFamily
	ProcessorAMDTurionTM64MobileTechnology
	ProcessorDual_CoreAMDOpteronTMProcessorFamily
	ProcessorAMDAthlonTM64X2Dual_CoreProcessorFamily
	ProcessorAMDTurionTM64X2MobileTechnology
	ProcessorQuad_CoreAMDOpteronTMProcessorFamily
	ProcessorThird_GenerationAMDOpteronTMProcessorFamily
	ProcessorAMDPhenomTMFXQuad_CoreProcessorFamily
	ProcessorAMDPhenomTMX4Quad_CoreProcessorFamily
	ProcessorAMDPhenomTMX2Dual_CoreProcessorFamily
	ProcessorAMDAthlonTMX2Dual_CoreProcessorFamily
	ProcessorPA_RISCFamily
	ProcessorPA_RISC8500
	ProcessorPA_RISC8000
	ProcessorPA_RISC7300LC
	ProcessorPA_RISC7200
	ProcessorPA_RISC7100LC
	ProcessorPA_RISC7100
	_
	_
	_
	_
	_
	_
	_
	_
	_
	ProcessorV30Family
	ProcessorQuad_CoreIntelXeonprocessor3200Series
	ProcessorDual_CoreIntelXeonprocessor3000Series
	ProcessorQuad_CoreIntelXeonprocessor5300Series
	ProcessorDual_CoreIntelXeonprocessor5100Series
	ProcessorDual_CoreIntelXeonprocessor5000Series
	ProcessorDual_CoreIntelXeonprocessorLV
	ProcessorDual_CoreIntelXeonprocessorULV
	ProcessorDual_CoreIntelXeonprocessor7100Series
	ProcessorQuad_CoreIntelXeonprocessor5400Series
	ProcessorQuad_CoreIntelXeonprocessor
	ProcessorDual_CoreIntelXeonprocessor5200Series
	ProcessorDual_CoreIntelXeonprocessor7200Series
	ProcessorQuad_CoreIntelXeonprocessor7300Series
	ProcessorQuad_CoreIntelXeonprocessor7400Series
	ProcessorMulti_CoreIntelXeonprocessor7400Series
	ProcessorPentiumIIIXeonTMprocessor
	ProcessorPentiumIIIProcessorwithIntelSpeedStepTMTechnology
	ProcessorPentium4Processor
	ProcessorIntelXeonprocessor
	ProcessorAS400Family
	ProcessorIntelXeonTMprocessorMP
	ProcessorAMDAthlonTMXPProcessorFamily
	ProcessorAMDAthlonTMMPProcessorFamily
	ProcessorIntelItanium2processor
	ProcessorIntelPentiumMprocessor
	ProcessorIntelCeleronDprocessor
	ProcessorIntelPentiumDprocessor
	ProcessorIntelPentiumProcessorExtremeEdition
	ProcessorIntelCoreTMSoloProcessor
	ProcessorReserved
	ProcessorIntelCoreTM2DuoProcessor
	ProcessorIntelCoreTM2Soloprocessor
	ProcessorIntelCoreTM2Extremeprocessor
	ProcessorIntelCoreTM2Quadprocessor
	ProcessorIntelCoreTM2Extrememobileprocessor
	ProcessorIntelCoreTM2Duomobileprocessor
	ProcessorIntelCoreTM2Solomobileprocessor
	ProcessorIntelCoreTMi7processor
	ProcessorDual_CoreIntelCeleronprocessor
	ProcessorIBM390Family
	ProcessorG4
	ProcessorG5
	ProcessorESA390G6
	ProcessorzArchitecturebase
	ProcessorIntelCoreTMi5processor
	ProcessorIntelCoreTMi3processor
	_
	_
	_
	ProcessorVIAC7TM_MProcessorFamily
	ProcessorVIAC7TM_DProcessorFamily
	ProcessorVIAC7TMProcessorFamily
	ProcessorVIAEdenTMProcessorFamily
	ProcessorMulti_CoreIntelXeonprocessor
	ProcessorDual_CoreIntelXeonprocessor3xxxSeries
	ProcessorQuad_CoreIntelXeonprocessor3xxxSeries
	ProcessorVIANanoTMProcessorFamily
	ProcessorDual_CoreIntelXeonprocessor5xxxSeries
	ProcessorQuad_CoreIntelXeonprocessor5xxxSeries
	_
	ProcessorDual_CoreIntelXeonprocessor7xxxSeries
	ProcessorQuad_CoreIntelXeonprocessor7xxxSeries
	ProcessorMulti_CoreIntelXeonprocessor7xxxSeries
	ProcessorMulti_CoreIntelXeonprocessor3400Series
	_
	_
	_
	ProcessorAMDOpteronTM3000SeriesProcessor
	ProcessorAMDSempronTMIIProcessor
	ProcessorEmbeddedAMDOpteronTMQuad_CoreProcessorFamily
	ProcessorAMDPhenomTMTriple_CoreProcessorFamily
	ProcessorAMDTurionTMUltraDual_CoreMobileProcessorFamily
	ProcessorAMDTurionTMDual_CoreMobileProcessorFamily
	ProcessorAMDAthlonTMDual_CoreProcessorFamily
	ProcessorAMDSempronTMSIProcessorFamily
	ProcessorAMDPhenomTMIIProcessorFamily
	ProcessorAMDAthlonTMIIProcessorFamily
	ProcessorSix_CoreAMDOpteronTMProcessorFamily
	ProcessorAMDSempronTMMProcessorFamily
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	Processori860
	Processori960
	_
	_
	ProcessorIndicatortoobtaintheprocessorfamilyfromtheProcessorFamily2field
	_
	_
	_
	_
	_
	_
	ProcessorSH_3
	ProcessorSH_4
	ProcessorARM
	ProcessorStrongARM
	Processor6x86
	ProcessorMediaGX
	ProcessorMII
	ProcessorWinChip
	ProcessorDSP
	ProcessorVideoProcessor
)

func (p ProcessorFamily) String() string {
	families := [...]string{
		"Other",
		"Unknown",
		"8086",
		"80286",
		"Intel386TM processor",
		"Intel486TM processor",
		"8087",
		"80287",
		"80387",
		"80487",
		"Intel® Pentium® processor",
		"Pentium® Pro processor",
		"Pentium® II processor",
		"Pentium® processor with MMXTM technology",
		"Intel® Celeron® processor",
		"Pentium® II XeonTM processor",
		"Pentium® III processor",
		"M1 Family",
		"M2 Family",
		"Intel® Celeron® M processor",
		"Intel® Pentium® 4 HT processor",
		"Available for assignment",
		"Available for assignment",
		"AMD DuronTM Processor Family",
		"K5 Family",
		"K6 Family",
		"K6-2",
		"K6-3",
		"AMD AthlonTM Processor Family",
		"AMD29000 Family",
		"K6-2+",
		"Power PC Family",
		"Power PC 601",
		"Power PC 603",
		"Power PC 603+",
		"Power PC 604",
		"Power PC 620",
		"Power PC x704",
		"Power PC 750",
		"Intel® CoreTM Duo processor",
		"Intel® CoreTM Duo mobile processor",
		"Intel® CoreTM Solo mobile processor",
		"Intel® AtomTM processor",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Alpha Family",
		"Alpha 21064",
		"Alpha 21066",
		"Alpha 21164",
		"Alpha 21164PC",
		"Alpha 21164a",
		"Alpha 21264",
		"Alpha 21364",
		"AMD TurionTM II Ultra Dual-Core Mobile M Processor Family",
		"AMD TurionTM II Dual-Core Mobile M Processor Family",
		"AMD AthlonTM II Dual-Core M Processor Family",
		"AMD OpteronTM 6100 Series Processor",
		"AMD OpteronTM 4100 Series Processor",
		"AMD OpteronTM 6200 Series Processor",
		"AMD OpteronTM 4200 Series Processor",
		"AMD FXTM Series Processor",
		"MIPS Family",
		"MIPS R4000",
		"MIPS R4200",
		"MIPS R4400",
		"MIPS R4600",
		"MIPS R10000",
		"AMD C-Series Processor",
		"AMD E-Series Processor",
		"AMD A-Series Processor",
		"AMD G-Series Processor",
		"AMD Z-Series Processor",
		"AMD R-Series Processor",
		"AMD OpteronTM 4300 Series Processor",
		"AMD OpteronTM 6300 Series Processor",
		"AMD OpteronTM 3300 Series Processor",
		"AMD FireProTM Series Processor",
		"SPARC Family",
		"SuperSPARC",
		"microSPARC II",
		"microSPARC IIep",
		"UltraSPARC",
		"UltraSPARC II",
		"UltraSPARC Iii",
		"UltraSPARC III",
		"UltraSPARC IIIi",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"68040 Family",
		"68xxx",
		"68000",
		"68010",
		"68020",
		"68030",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Hobbit Family",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"CrusoeTM TM5000 Family",
		"CrusoeTM TM3000 Family",
		"EfficeonTM TM8000 Family",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Weitek",
		"Available for assignment",
		"ItaniumTM processor",
		"AMD AthlonTM 64 Processor Family",
		"AMD OpteronTM Processor Family",
		"AMD SempronTM Processor Family",
		"AMD TurionTM 64 Mobile Technology",
		"Dual-Core AMD OpteronTM Processor Family",
		"AMD AthlonTM 64 X2 Dual-Core Processor Family",
		"AMD TurionTM 64 X2 Mobile Technology",
		"Quad-Core AMD OpteronTM Processor Family",
		"Third-Generation AMD OpteronTM Processor Family",
		"AMD PhenomTM FX Quad-Core Processor Family",
		"AMD PhenomTM X4 Quad-Core Processor Family",
		"AMD PhenomTM X2 Dual-Core Processor Family",
		"AMD AthlonTM X2 Dual-Core Processor Family",
		"PA-RISC Family",
		"PA-RISC 8500",
		"PA-RISC 8000",
		"PA-RISC 7300LC",
		"PA-RISC 7200",
		"PA-RISC 7100LC",
		"PA-RISC 7100",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"V30 Family",
		"Quad-Core Intel® Xeon® processor 3200 Series",
		"Dual-Core Intel® Xeon® processor 3000 Series",
		"Quad-Core Intel® Xeon® processor 5300 Series",
		"Dual-Core Intel® Xeon® processor 5100 Series",
		"Dual-Core Intel® Xeon® processor 5000 Series",
		"Dual-Core Intel® Xeon® processor LV",
		"Dual-Core Intel® Xeon® processor ULV",
		"Dual-Core Intel® Xeon® processor 7100 Series",
		"Quad-Core Intel® Xeon® processor 5400 Series",
		"Quad-Core Intel® Xeon® processor",
		"Dual-Core Intel® Xeon® processor 5200 Series",
		"Dual-Core Intel® Xeon® processor 7200 Series",
		"Quad-Core Intel® Xeon® processor 7300 Series",
		"Quad-Core Intel® Xeon® processor 7400 Series",
		"Multi-Core Intel® Xeon® processor 7400 Series",
		"Pentium® III XeonTM processor",
		"Pentium® III Processor with Intel® SpeedStepTM Technology",
		"Pentium® 4 Processor",
		"Intel® Xeon® processor",
		"AS400 Family",
		"Intel® XeonTM processor MP",
		"AMD AthlonTM XP Processor Family",
		"AMD AthlonTM MP Processor Family",
		"Intel® Itanium® 2 processor",
		"Intel® Pentium® M processor",
		"Intel® Celeron® D processor",
		"Intel® Pentium® D processor",
		"Intel® Pentium® Processor Extreme Edition",
		"Intel® CoreTM Solo Processor",
		"Reserved",
		"Intel® CoreTM 2 Duo Processor",
		"Intel® CoreTM 2 Solo processor",
		"Intel® CoreTM 2 Extreme processor",
		"Intel® CoreTM 2 Quad processor",
		"Intel® CoreTM 2 Extreme mobile processor",
		"Intel® CoreTM 2 Duo mobile processor",
		"Intel® CoreTM 2 Solo mobile processor",
		"Intel® CoreTM i7 processor",
		"Dual-Core Intel® Celeron® processor",
		"IBM390 Family",
		"G4",
		"G5",
		"ESA/390 G6",
		"z/Architecture base",
		"Intel® CoreTM i5 processor",
		"Intel® CoreTM i3 processor",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"VIA C7TM-M Processor Family",
		"VIA C7TM-D Processor Family",
		"VIA C7TM Processor Family",
		"VIA EdenTM Processor Family",
		"Multi-Core Intel® Xeon® processor",
		"Dual-Core Intel® Xeon® processor 3xxx Series",
		"Quad-Core Intel® Xeon® processor 3xxx Series",
		"VIA NanoTM Processor Family",
		"Dual-Core Intel® Xeon® processor 5xxx Series",
		"Quad-Core Intel® Xeon® processor 5xxx Series",
		"Available for assignment",
		"Dual-Core Intel® Xeon® processor 7xxx Series",
		"Quad-Core Intel® Xeon® processor 7xxx Series",
		"Multi-Core Intel® Xeon® processor 7xxx Series",
		"Multi-Core Intel® Xeon® processor 3400 Series",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"AMD OpteronTM 3000 Series Processor",
		"AMD SempronTM II Processor",
		"Embedded AMD OpteronTM Quad-Core Processor Family",
		"AMD PhenomTM Triple-Core Processor Family",
		"AMD TurionTM Ultra Dual-Core Mobile Processor Family",
		"AMD TurionTM Dual-Core Mobile Processor Family",
		"AMD AthlonTM Dual-Core Processor Family",
		"AMD SempronTM SI Processor Family",
		"AMD PhenomTM II Processor Family",
		"AMD AthlonTM II Processor Family",
		"Six-Core AMD OpteronTM Processor Family",
		"AMD SempronTM M Processor Family",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"i860",
		"i960",
		"Available for assignment",
		"Available for assignment",
		"Indicator to obtain the processor family from the Processor Family 2 field",
		"Reserved ￼",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"Available for assignment",
		"SH-3",
		"SH-4",
		"ARM",
		"StrongARM",
		"6x86",
		"MediaGX",
		"MII",
		"WinChip",
		"DSP",
		"Video Processor",
		"Available for assignment",
		"Reserved",
	}
	return families[p]
}

type ProcessorID []byte

func (p ProcessorID) String() string {
	t := []string{}
	for i := range p {
		t = append(t, fmt.Sprintf("%02X", p[i]))
	}

	return strings.Join(t, " ")
}

type ProcessorVoltage byte

const (
	ProcessorVoltage5V ProcessorVoltage = 1 << iota
	ProcessorVoltage3dot3V
	ProcessorVoltage2dot9V
	ProcessorVoltageReserved
	_
	_
	_
	ProcessorVoltageLegacy
)

func (p ProcessorVoltage) String() string {
	voltages := [...]string{
		"5V",
		"3.3V",
		"2.9V",
	}
	if p&ProcessorVoltageLegacy == 0 {
		return voltages[p]
	}
	return fmt.Sprintf("%.1f", float64(p-0x80)/10)
}

type ProcessorStatus byte

const (
	ProcessorStatusUnknown ProcessorStatus = iota
	ProcessorStatusEnabled
	ProcessorStatusDisabledByUser
	ProcessorStatusDisabledByBIOS
	ProcessorStatusIdle
	_
	_
	ProcessorStatusOther
)

func (p ProcessorStatus) String() string {
	status := [...]string{
		"Unknown",
		"CPU Enabled",
		"Disabled By User through BIOS Setup",
		"Disabled By BIOSa(POST Error)",
		"CPU is Idle, waiting to be enabled",
		"Reserved",
		"Reserved",
		"Other",
	}

	resp := []string{}
	switch uint8(p) & 64 {
	case 0:
		resp = append(resp, "Unpopulated")
	case 64:
		resp = append(resp, "Populated")
	}

	resp = append(resp, status[uint8(p)&7])

	return strings.Join(resp, ",")
}

type ProcessorUpgrade byte

const (
	_ ProcessorUpgrade = iota
	ProcessorUpgradeOther
	ProcessorUpgradeUnknown
	ProcessorUpgradeDaughterBoard
	ProcessorUpgradeZIFSocket
	ProcessorUpgradeReplaceablePiggyBack
	ProcessorUpgradeNone
	ProcessorUpgradeLIFSocket
	ProcessorUpgradeSlot1
	ProcessorUpgradeSlot2
	ProcessorUpgrade370_pinsocket
	ProcessorUpgradeSlotA
	ProcessorUpgradeSlotM
	ProcessorUpgradeSocket423
	ProcessorUpgradeSocketASocket462
	ProcessorUpgradeSocket478
	ProcessorUpgradeSocket754
	ProcessorUpgradeSocket940
	ProcessorUpgradeSocket939
	ProcessorUpgradeSocketmPGA604
	ProcessorUpgradeSocketLGA771
	ProcessorUpgradeSocketLGA775
	ProcessorUpgradeSocketS1
	ProcessorUpgradeSocketAM2
	ProcessorUpgradeSocketF1207
	ProcessorUpgradeSocketLGA1366
	ProcessorUpgradeSocketG34
	ProcessorUpgradeSocketAM3
	ProcessorUpgradeSocketC32
	ProcessorUpgradeSocketLGA1156
	ProcessorUpgradeSocketLGA1567
	ProcessorUpgradeSocketPGA988A
	ProcessorUpgradeSocketBGA1288
	ProcessorUpgradeSocketrPGA988B
	ProcessorUpgradeSocketBGA1023
	ProcessorUpgradeSocketBGA1224
	ProcessorUpgradeSocketLGA1155
	ProcessorUpgradeSocketLGA1356
	ProcessorUpgradeSocketLGA2011
	ProcessorUpgradeSocketFS1
	ProcessorUpgradeSocketFS2
	ProcessorUpgradeSocketFM1
	ProcessorUpgradeSocketFM2
	ProcessorUpgradeSocketLGA2011_3
	ProcessorUpgradeSocketLGA1356_3
)

func (p ProcessorUpgrade) String() string {
	upgrades := [...]string{
		"Other",
		"Unknown",
		"Daughter Board",
		"ZIF Socket",
		"Replaceable Piggy Back",
		"None",
		"LIF Socket",
		"Slot 1",
		"Slot 2",
		"370-pin socket",
		"Slot A",
		"Slot M",
		"Socket 423",
		"Socket A (Socket 462)",
		"Socket 478",
		"Socket 754",
		"Socket 940",
		"Socket 939",
		"Socket mPGA604",
		"Socket LGA771",
		"Socket LGA775",
		"Socket S1",
		"Socket AM2",
		"Socket F (1207)",
		"Socket LGA1366",
		"Socket G34",
		"Socket AM3",
		"Socket C32",
		"Socket LGA1156",
		"Socket LGA1567",
		"Socket PGA988A",
		"Socket BGA1288",
		"Socket rPGA988B",
		"Socket BGA1023",
		"Socket BGA1224",
		"Socket LGA1155",
		"Socket LGA1356",
		"Socket LGA2011",
		"Socket FS1",
		"Socket FS2",
		"Socket FM1",
		"Socket FM2",
		"Socket LGA2011-3",
		"Socket LGA1356-3",
	}
	return upgrades[p]
}

type ProcessorCharacteristics uint16

const (
	ProcessorCharacteristicsReserved ProcessorCharacteristics = 1
	ProcessorCharacteristicsUnknown
	ProcessorCharacteristics64_bitCapable
	ProcessorCharacteristicsMulti_Core
	ProcessorCharacteristicsHardwareThread
	ProcessorCharacteristicsExecuteProtection
	ProcessorCharacteristicsEnhancedVirtualization
	ProcessorCharacteristicsPowerPerformanceControl
)

func (p ProcessorCharacteristics) String() string {
	chars := [...]string{
		"Reserved",
		"Unknown",
		"64-bit Capable",
		"Multi-Core",
		"Hardware Thread",
		"Execute Protection",
		"Enhanced Virtualization",
		"Power/Performance Control",
	}

	resp := []string{}
	for i := 1; i <= 16; i++ {
		if p&(1<<uint(i)) != 0 {
			resp = append(resp, chars[i])
		}
	}

	return strings.Join(resp, ",")
}
