package main

import (
	"fmt"

	"dmidecode/smbios"
)

func main() {
	ss, err := smbios.ReadStructures()
	if err != nil {
		panic(err)
	}

	for _, s := range ss {
		fmt.Println(s)
		for i, v := range s.Strings {
			fmt.Println(i, " ", v)
		}
	}
}
