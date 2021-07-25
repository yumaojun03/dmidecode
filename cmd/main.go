package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yumaojun03/dmidecode"
)

var (
	debug bool
	help  bool
	t     string
)

var (
	decoder *dmidecode.Decoder
)

func main() {
	dmi, err := dmidecode.New()
	if err != nil {
		fmt.Printf("init decoder error, %s", err)
		os.Exit(1)
	}
	decoder = dmi

	flag.Parse()
	decoder.Debug(debug)

	if help {
		usage()
		os.Exit(0)
	}

	eps := decoder.EntryPoint()
	fmt.Printf("SMBIOS %d.%d.%d - table: address: %#x, size: %d\n",
		eps.Major, eps.Minor, eps.Revision, eps.Address, eps.Size)

	switch t {
	case "":
		infos, err := decoder.ALL()
		checkDecodeErr(t, err)
		infos.Print()
	case "bios", "0":
		infos, err := decoder.BIOS()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
		}
	case "system", "1":
		infos, err := decoder.System()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
		}
	case "baseboard", "2":
		infos, err := decoder.BaseBoard()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
		}
	case "chassis", "3":
		infos, err := decoder.Chassis()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
		}
	case "4":
		infos, err := decoder.Processor()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
		}
	case "7":
		pcinfos, err := decoder.ProcessorCache()
		checkDecodeErr(t, err)
		for i := range pcinfos {
			fmt.Println(pcinfos[i])
		}
	case "port", "8":
		infos, err := decoder.PortConnector()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
		}
	case "slot", "9":
		infos, err := decoder.Slot()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
		}
	case "10":
		infos, err := decoder.Onboard()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
		}
	case "11":
		oems, err := decoder.OEM()
		checkDecodeErr(t, err)
		for i := range oems {
			fmt.Println(oems[i])
		}
	case "16":
		mainfos, err := decoder.MemoryArray()
		checkDecodeErr(t, err)
		for i := range mainfos {
			fmt.Println(mainfos[i])
		}
	case "17":
		mdinfos, err := decoder.MemoryDevice()
		checkDecodeErr(t, err)
		for i := range mdinfos {
			fmt.Println(mdinfos[i])
		}
	case "battery", "22":
		infos, err := decoder.Battery()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
		}
	case "power", "39":
		infos, err := decoder.PowerSupply()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
		}
	case "41":
		einfos, err := decoder.OnboardExtended()
		checkDecodeErr(t, err)
		for i := range einfos {
			fmt.Println(einfos[i])
		}
	case "onboard":
		infos, err := decoder.Onboard()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
		}
		einfos, err := decoder.OnboardExtended()
		checkDecodeErr(t, err)
		for i := range einfos {
			fmt.Println(einfos[i])
		}
	case "processor":
		infos, err := decoder.Processor()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
		}
		pcinfos, err := decoder.ProcessorCache()
		checkDecodeErr(t, err)
		for i := range pcinfos {
			fmt.Println(pcinfos[i])
		}
	case "memory":
		mainfos, err := decoder.MemoryArray()
		checkDecodeErr(t, err)
		for i := range mainfos {
			fmt.Println(mainfos[i])
		}
		mdinfos, err := decoder.MemoryDevice()
		checkDecodeErr(t, err)
		for i := range mdinfos {
			fmt.Println(mdinfos[i])
		}
	default:
		fmt.Printf("ERR: unknown decode type %s\n", t)
		usage()
		os.Exit(1)
	}
}

func checkDecodeErr(t string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERR: decode %s error, %s", t, err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `dmidecode version: 0.1.3
Usage: dmidecode [-h] [-d] [-t type]

Options:
`)
	flag.PrintDefaults()
}

func init() {
	flag.BoolVar(&help, "h", false, "this help")
	flag.BoolVar(&debug, "d", false, "debug mode")
	flag.StringVar(&t, "t", "", "dmidocode type [bios, system, baseboard, chassis, onboard, port, processor, memory, slot]")
	flag.Usage = usage
}
