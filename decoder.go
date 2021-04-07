package dmidecode

import (
	"fmt"

	"github.com/yumaojun03/dmidecode/parser/baseboard"
	"github.com/yumaojun03/dmidecode/parser/battery"
	"github.com/yumaojun03/dmidecode/parser/bios"
	"github.com/yumaojun03/dmidecode/parser/chassis"
	"github.com/yumaojun03/dmidecode/parser/memory"
	"github.com/yumaojun03/dmidecode/parser/oem"
	"github.com/yumaojun03/dmidecode/parser/onboard"
	"github.com/yumaojun03/dmidecode/parser/port"
	"github.com/yumaojun03/dmidecode/parser/processor"
	"github.com/yumaojun03/dmidecode/parser/slot"
	"github.com/yumaojun03/dmidecode/parser/system"
	"github.com/yumaojun03/dmidecode/smbios"
)

// New 实例化
func New() (*Decoder, error) {
	eps, ss, err := smbios.ReadStructures()
	if err != nil {
		return nil, err
	}

	d := new(Decoder)
	d.eps = eps
	d.total = len(ss)

	for i := range ss {
		switch smbios.StructureType(ss[i].Header.Type) {
		case smbios.BIOS:
			d.bios = append(d.bios, ss[i])
		case smbios.System:
			d.system = append(d.system, ss[i])
		case smbios.BaseBoard:
			d.baseBoard = append(d.baseBoard, ss[i])
		case smbios.Chassis:
			d.chassis = append(d.chassis, ss[i])
		case smbios.Processor:
			d.processor = append(d.processor, ss[i])
		case smbios.OnBoardDevices:
			d.onBoardDevices = append(d.onBoardDevices, ss[i])
		case smbios.OEMStrings:
			d.omeStrings = append(d.omeStrings, ss[i])
		case smbios.OnBoardDevicesExtendedInformation:
			d.onBoardExtendedDevices = append(d.onBoardExtendedDevices, ss[i])
		case smbios.PortConnector:
			d.portConnector = append(d.portConnector, ss[i])
		case smbios.Cache:
			d.cache = append(d.cache, ss[i])
		case smbios.PhysicalMemoryArray:
			d.physicalMemoryArray = append(d.physicalMemoryArray, ss[i])
		case smbios.MemoryDevice:
			d.memoryDevice = append(d.memoryDevice, ss[i])
		case smbios.SystemSlots:
			d.systemSlots = append(d.systemSlots, ss[i])
		case smbios.PortableBattery:
			d.portableBattery = append(d.portableBattery, ss[i])
		default:
		}
	}

	return d, nil
}

// Decoder decoder
type Decoder struct {
	debug bool
	total int

	eps                    *smbios.EntryPoint
	bios                   []*smbios.Structure
	system                 []*smbios.Structure
	baseBoard              []*smbios.Structure
	chassis                []*smbios.Structure
	onBoardDevices         []*smbios.Structure
	omeStrings             []*smbios.Structure
	onBoardExtendedDevices []*smbios.Structure
	portConnector          []*smbios.Structure
	processor              []*smbios.Structure
	cache                  []*smbios.Structure
	physicalMemoryArray    []*smbios.Structure
	memoryDevice           []*smbios.Structure
	systemSlots            []*smbios.Structure
	portableBattery        []*smbios.Structure
}

// Debug 开关Debug
func (d *Decoder) Debug(on bool) {
	d.debug = on
}

// BIOS 解析bios信息
func (d *Decoder) BIOS() ([]*bios.Information, error) {
	infos := make([]*bios.Information, 0, len(d.bios))
	for i := range d.bios {
		d.println(d.bios[i])
		info, err := bios.Parse(d.bios[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// System 解析system信息
func (d *Decoder) System() ([]*system.Information, error) {
	infos := make([]*system.Information, 0, len(d.system))
	for i := range d.system {
		d.println(d.system[i])
		info, err := system.Parse(d.system[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// BaseBoard 解析baseboard信息
func (d *Decoder) BaseBoard() ([]*baseboard.Information, error) {
	infos := make([]*baseboard.Information, 0, len(d.baseBoard))
	for i := range d.baseBoard {
		d.println(d.baseBoard[i])
		info, err := baseboard.Parse(d.baseBoard[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// Chassis 解析chassis信息
func (d *Decoder) Chassis() ([]*chassis.Information, error) {
	infos := make([]*chassis.Information, 0, len(d.chassis))
	for i := range d.chassis {
		d.println(d.chassis[i])
		info, err := chassis.Parse(d.chassis[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// OnboardExtended 解析onboard信息
func (d *Decoder) OnboardExtended() ([]*onboard.ExtendedInformation, error) {
	infos := make([]*onboard.ExtendedInformation, 0, len(d.onBoardExtendedDevices))
	for i := range d.onBoardExtendedDevices {
		d.println(d.onBoardExtendedDevices[i])
		info, err := onboard.ParseType41(d.onBoardExtendedDevices[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// Onboard 解析onboard信息
func (d *Decoder) Onboard() ([]*onboard.Information, error) {
	infos := make([]*onboard.Information, 0, len(d.onBoardDevices))
	for i := range d.onBoardDevices {
		d.println(d.onBoardDevices[i])
		info, err := onboard.ParseType10(d.onBoardDevices[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

func (d *Decoder) OEM() ([]*oem.OEM, error) {
	infos := make([]*oem.OEM, 0, len(d.omeStrings))
	for i := range d.omeStrings {
		d.println(d.omeStrings[i])
		info, err := oem.Parse(d.omeStrings[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// PortConnector 解析port connector信息
func (d *Decoder) PortConnector() ([]*port.Information, error) {
	infos := make([]*port.Information, 0, len(d.portConnector))
	for i := range d.portConnector {
		d.println(d.portConnector[i])
		info, err := port.Parse(d.portConnector[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// Processor 解析processor信息
func (d *Decoder) Processor() ([]*processor.Processor, error) {
	infos := make([]*processor.Processor, 0, len(d.processor))
	for i := range d.processor {
		d.println(d.processor[i])
		info, err := processor.ParseProcessor(d.processor[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// ProcessorCache 解析processor cache信息
func (d *Decoder) ProcessorCache() ([]*processor.Cache, error) {
	infos := make([]*processor.Cache, 0, len(d.cache))
	for i := range d.cache {
		d.println(d.cache[i])
		info, err := processor.ParseCache(d.cache[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// MemoryArray 解析memory array信息
func (d *Decoder) MemoryArray() ([]*memory.PhysicalMemoryArray, error) {
	infos := make([]*memory.PhysicalMemoryArray, 0, len(d.physicalMemoryArray))
	for i := range d.physicalMemoryArray {
		d.println(d.physicalMemoryArray[i])
		info, err := memory.ParseMemoryArray(d.physicalMemoryArray[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// MemoryDevice 解析memory device信息
func (d *Decoder) MemoryDevice() ([]*memory.MemoryDevice, error) {
	infos := make([]*memory.MemoryDevice, 0, len(d.memoryDevice))
	for i := range d.memoryDevice {
		d.println(d.memoryDevice[i])
		info, err := memory.ParseMemoryDevice(d.memoryDevice[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// Slot 解析memory device信息
func (d *Decoder) Slot() ([]*slot.SystemSlot, error) {
	infos := make([]*slot.SystemSlot, 0, len(d.systemSlots))
	for i := range d.systemSlots {
		d.println(d.systemSlots[i])
		info, err := slot.Parse(d.systemSlots[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// Battery 解析battery信息
func (d *Decoder) Battery() ([]*battery.Information, error) {
	infos := make([]*battery.Information, 0, len(d.portableBattery))
	for i := range d.portableBattery {
		d.println(d.portableBattery[i])
		info, err := battery.Parse(d.portableBattery[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}
	return infos, nil
}

// EntryPoint todo
func (d *Decoder) EntryPoint() *smbios.EntryPoint {
	return d.eps
}

func (d *Decoder) println(s *smbios.Structure) {
	if d.debug {
		fmt.Printf("[Debug] *smbios.Structure: %s\n", s)
	}
}

// ALL decode all
func (d *Decoder) ALL() (*InformationSet, error) {
	errs := NewErrorSet()
	sets := NewInformationSet()

	biosInfos, err := d.BIOS()
	errs.checkOrAdd(err)
	sets.addBios(biosInfos)

	systemInfos, err := d.System()
	errs.checkOrAdd(err)
	sets.addSystem(systemInfos)

	bbInfos, err := d.BaseBoard()
	errs.checkOrAdd(err)
	sets.addBaseBoard(bbInfos)

	csInfos, err := d.Chassis()
	errs.checkOrAdd(err)
	sets.addChassis(csInfos)

	oemInfos, err := d.OEM()
	errs.checkOrAdd(err)
	sets.addOEM(oemInfos)

	obInfos, err := d.OnboardExtended()
	errs.checkOrAdd(err)
	sets.addOnboard(obInfos)

	pcInfos, err := d.PortConnector()
	errs.checkOrAdd(err)
	sets.addPortConnector(pcInfos)

	processorInfos, err := d.Processor()
	errs.checkOrAdd(err)
	sets.addProcessor(processorInfos)

	pcacheInfos, err := d.ProcessorCache()
	errs.checkOrAdd(err)
	sets.addCache(pcacheInfos)

	maInfos, err := d.MemoryArray()
	errs.checkOrAdd(err)
	sets.addMemoryArray(maInfos)

	mdInfos, err := d.MemoryDevice()
	errs.checkOrAdd(err)
	sets.addMemoryDevice(mdInfos)

	slotsInfos, err := d.Slot()
	errs.checkOrAdd(err)
	sets.addSlot(slotsInfos)

	batteryInfos, err := d.Battery()
	errs.checkOrAdd(err)
	sets.addBattery(batteryInfos)

	return sets, errs.Error()
}
