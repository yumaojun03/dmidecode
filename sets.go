package dmidecode

import (
	"errors"
	"strings"

	"github.com/yumaojun03/dmidecode/parser/baseboard"
	"github.com/yumaojun03/dmidecode/parser/bios"
	"github.com/yumaojun03/dmidecode/parser/chassis"
	"github.com/yumaojun03/dmidecode/parser/memory"
	"github.com/yumaojun03/dmidecode/parser/onboard"
	"github.com/yumaojun03/dmidecode/parser/port"
	"github.com/yumaojun03/dmidecode/parser/processor"
	"github.com/yumaojun03/dmidecode/parser/slot"
	"github.com/yumaojun03/dmidecode/parser/system"
)

// NewInformationSet todo
func NewInformationSet() *InformationSet {
	return &InformationSet{
		bios:         []*bios.Information{},
		system:       []*system.Information{},
		baseboard:    []*baseboard.Information{},
		chassis:      []*chassis.Information{},
		onboard:      []*onboard.ExtendedInformation{},
		portConnetor: []*port.Information{},
		processor:    []*processor.Processor{},
		cache:        []*processor.Cache{},
		memoryArray:  []*memory.PhysicalMemoryArray{},
		memoryDevice: []*memory.MemoryDevice{},
		slot:         []*slot.SystemSlot{},
	}
}

// InformationSet 集合
type InformationSet struct {
	bios         []*bios.Information
	system       []*system.Information
	baseboard    []*baseboard.Information
	chassis      []*chassis.Information
	onboard      []*onboard.ExtendedInformation
	portConnetor []*port.Information
	processor    []*processor.Processor
	cache        []*processor.Cache
	memoryArray  []*memory.PhysicalMemoryArray
	memoryDevice []*memory.MemoryDevice
	slot         []*slot.SystemSlot
}

func (s *InformationSet) addBios(infos []*bios.Information) {
	s.bios = infos
}

func (s *InformationSet) addSystem(infos []*system.Information) {
	s.system = infos
}

func (s *InformationSet) addBaseBoard(infos []*baseboard.Information) {
	s.baseboard = infos
}

// NewErrorSet todo
func NewErrorSet() *ErrorSet {
	return &ErrorSet{
		errs: []error{},
	}
}

// ErrorSet errorset
type ErrorSet struct {
	errs []error
}

func (s *ErrorSet) checkOrAdd(err error) {
	if err != nil {
		s.errs = append(s.errs, err)
	}
}

// HasError todo
func (s *ErrorSet) Error() error {
	if len(s.errs) > 0 {
		errMSG := make([]string, 0, len(s.errs))
		for i := range s.errs {
			errMSG = append(errMSG, s.errs[i].Error())
		}
		errStrs := strings.Join(errMSG, ",")
		return errors.New(errStrs)
	}

	return nil
}
