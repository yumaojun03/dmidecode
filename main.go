package main

import (
	"dmidecode/parser/baseboard"
	"dmidecode/parser/bios"
	"dmidecode/parser/chassis"
	"dmidecode/parser/memory"
	"dmidecode/parser/onboard"
	"dmidecode/parser/processor"
	"dmidecode/parser/slot"
	"dmidecode/parser/system"
	"dmidecode/smbios"
	"fmt"
)

func main() {
	ss, err := smbios.ReadStructures()
	if err != nil {
		panic(err)
	}

	for _, s := range ss {
		switch smbios.StructureType(s.Header.Type) {
		case smbios.BIOS:
			fmt.Println(s)
			info, err := bios.Parse(s)
			if err != nil {
				panic(err)
			}

			fmt.Println(info)

		case smbios.System:
			fmt.Println(s)
			info, err := system.Parse(s)
			if err != nil {
				panic(err)
			}

			fmt.Println(info)

		case smbios.BaseBoard:
			fmt.Println(s)
			info, err := baseboard.Parse(s)
			if err != nil {
				panic(err)
			}

			fmt.Println(info)

		case smbios.Chassis:
			fmt.Println(s)
			info, err := chassis.Parse(s)
			if err != nil {
				panic(err)
			}

			fmt.Println(info)

		case smbios.OnBoardDevicesExtendedInformation:
			fmt.Println(s)
			info, err := onboard.Parse(s)
			if err != nil {
				panic(err)
			}

			fmt.Println(info)

		case smbios.Processor:
			fmt.Println(s)
			info, err := processor.ParseProcessor(s)
			if err != nil {
				panic(err)
			}

			fmt.Println(info)

		case smbios.Cache:
			fmt.Println(s)
			info, err := processor.ParseCache(s)
			if err != nil {
				panic(err)
			}

			fmt.Println(info)

		case smbios.PhysicalMemoryArray:
			fmt.Println(s)
			info, err := memory.ParseMemoryArray(s)
			if err != nil {
				panic(err)
			}

			fmt.Println(info)

		case smbios.MemoryDevice:
			fmt.Println(s)
			info, err := memory.ParseMemoryDevice(s)
			if err != nil {
				panic(err)
			}

			fmt.Println(info)

		case smbios.SystemSlots:
			fmt.Println(s)
			info, err := slot.Parse(s)
			if err != nil {
				panic(err)
			}

			fmt.Println(info)
		}

	}
}
