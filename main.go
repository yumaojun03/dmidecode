package main

import (
	"dmidecode/parser/baseboard"
	"dmidecode/parser/bios"
	"dmidecode/parser/chassis"
	"dmidecode/parser/onboard"
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
		}

	}
}
