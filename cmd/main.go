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
	case "bios":
		infos, err := decoder.BIOS()
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
	flag.StringVar(&t, "t", "", "dmidocode type")
	flag.Usage = usage
}
