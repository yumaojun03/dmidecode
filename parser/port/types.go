package port

type ConnectorType byte

const (
	ConnectorTypeNone ConnectorType = iota
	ConnectorTypeCentronics
	ConnectorTypeMiniCentronics
	ConnectorTypeProprietary
	ConnectorTypeDB_25pinmale
	ConnectorTypeDB_25pinfemale
	ConnectorTypeDB_15pinmale
	ConnectorTypeDB_15pinfemale
	ConnectorTypeDB_9pinmale
	ConnectorTypeDB_9pinfemale
	ConnectorTypeRJ_11
	ConnectorTypeRJ_45
	ConnectorType50_pinMiniSCSI
	ConnectorTypeMini_DIN
	ConnectorTypeMicro_DIN
	ConnectorTypePS2
	ConnectorTypeInfrared
	ConnectorTypeHP_HIL
	ConnectorTypeAccessBusUSB
	ConnectorTypeSSASCSI
	ConnectorTypeCircularDIN_8male
	ConnectorTypeCircularDIN_8female
	ConnectorTypeOnBoardIDE
	ConnectorTypeOnBoardFloppy
	ConnectorType9_pinDualInlinepin10cut
	ConnectorType25_pinDualInlinepin26cut
	ConnectorType50_pinDualInline
	ConnectorType68_pinDualInline
	ConnectorTypeOnBoardSoundInputfromCD_ROM
	ConnectorTypeMini_CentronicsType_14
	ConnectorTypeMini_CentronicsType_26
	ConnectorTypeMini_jackheadphones
	ConnectorTypeBNC
	ConnectorType1394
	ConnectorTypeSASSATAPlugReceptacle
	ConnectorTypePC_98
	ConnectorTypePC_98Hireso
	ConnectorTypePC_H98
	ConnectorTypePC_98Note
	ConnectorTypePC_98Full
	ConnectorTypeOther
)

func (p ConnectorType) String() string {
	types := [...]string{
		"None",
		"Centronics",
		"Mini Centronics",
		"Proprietary",
		"DB-25 pin male",
		"DB-25 pin female",
		"DB-15 pin male",
		"DB-15 pin female",
		"DB-9 pin male",
		"DB-9 pin female",
		"RJ-11",
		"RJ-45",
		"50-pin MiniSCSI",
		"Mini-DIN",
		"Micro-DIN",
		"PS/2",
		"Infrared",
		"HP-HIL",
		"Access Bus (USB)",
		"SSA SCSI",
		"Circular DIN-8 male",
		"Circular DIN-8 female",
		"On Board IDE",
		"On Board Floppy",
		"9-pin Dual Inline (pin 10 cut)",
		"25-pin Dual Inline (pin 26 cut)",
		"50-pin Dual Inline",
		"68-pin Dual Inline",
		"On Board Sound Input from CD-ROM",
		"Mini-Centronics Type-14",
		"Mini-Centronics Type-26",
		"Mini-jack (headphones)",
		"BNC",
		"1394",
		"SAS/SATA Plug Receptacle",
		"PC-98",
		"PC-98Hireso",
		"PC-H98",
		"PC-98Note",
		"PC-98Full",
		"Other",
	}
	return types[p]
}

type PortType byte

const (
	PortTypeNone PortType = iota
	PortTypeParallelPortXTATCompatible
	PortTypeParallelPortPS2
	PortTypeParallelPortECP
	PortTypeParallelPortEPP
	PortTypeParallelPortECPEPP
	PortTypeSerialPortXTATCompatible
	PortTypeSerialPort16450Compatible
	PortTypeSerialPort16550Compatible
	PortTypeSerialPort16550ACompatible
	PortTypeSCSIPort
	PortTypeMIDIPort
	PortTypeJoyStickPort
	PortTypeKeyboardPort
	PortTypeMousePort
	PortTypeSSASCSI
	PortTypeUSB
	PortTypeFireWireIEEEP1394
	PortTypePCMCIATypeI2
	PortTypePCMCIATypeII
	PortTypePCMCIATypeIII
	PortTypeCardbus
	PortTypeAccessBusPort
	PortTypeSCSIII
	PortTypeSCSIWide
	PortTypePC_98
	PortTypePC_98_Hireso
	PortTypePC_H98
	PortTypeVideoPort
	PortTypeAudioPort
	PortTypeModemPort
	PortTypeNetworkPort
	PortTypeSATA
	PortTypeSAS
	PortType8251Compatible
	PortType8251FIFOCompatible
	PortTypeOther
)

func (p PortType) String() string {
	types := [...]string{
		"None",
		"Parallel Port XT/AT Compatible",
		"Parallel Port PS/2",
		"Parallel Port ECP",
		"Parallel Port EPP",
		"Parallel Port ECP/EPP",
		"Serial Port XT/AT Compatible",
		"Serial Port 16450 Compatible",
		"Serial Port 16550 Compatible",
		"Serial Port 16550A Compatible",
		"SCSI Port",
		"MIDI Port",
		"Joy Stick Port",
		"Keyboard Port",
		"Mouse Port",
		"SSA SCSI",
		"USB",
		"FireWire (IEEE P1394)",
		"PCMCIA Type I2",
		"PCMCIA Type II",
		"PCMCIA Type III",
		"Cardbus",
		"Access Bus Port",
		"SCSI II",
		"SCSI Wide",
		"PC-98",
		"PC-98-Hireso",
		"PC-H98",
		"Video Port",
		"Audio Port",
		"Modem Port",
		"Network Port",
		"SATA",
		"SAS",
		"8251 Compatible",
		"8251 FIFO Compatible",
		" Other",
	}
	return types[p]
}
