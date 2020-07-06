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
	case "bios":
		infos, err := decoder.BIOS()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
		}
	case "system":
		infos, err := decoder.System()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
		}
	case "baseboard":
		infos, err := decoder.BaseBoard()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
		}
	case "chassis":
		infos, err := decoder.Chassis()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
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
	case "port":
		infos, err := decoder.PortConnector()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
		}
	case "processor":
		infos, err := decoder.Processor()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
		}
		pcinfos, err := decoder.ProcessorCache()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(pcinfos[i])
		}
	case "memory":
		infos, err := decoder.MemoryArray()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
		}
		mdinfos, err := decoder.MemoryDevice()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(mdinfos[i])
		}
	case "slot":
		infos, err := decoder.Slot()
		checkDecodeErr(t, err)
		for i := range infos {
			fmt.Println(infos[i])
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
	fmt.Fprintf(os.Stderr, `dmidecode version: 0.1.2
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
