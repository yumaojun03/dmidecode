package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yumaojun03/dmidecode"
)

var (
	h bool
	t string
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
	}
}

func checkDecodeErr(t string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "decode %s error, %s", t, err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `dmidecode version: 0.0.1
Usage: dmidecode [-h] [-t type] 

Options:
`)
	flag.PrintDefaults()
}

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&t, "t", "", "dmidocode type [bios, system, baseboard, chassis, onboard, port, processor, memory, slot")
	flag.Usage = usage
}
