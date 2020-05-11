# dmidecode
[![Go Report Card](https://goreportcard.com/badge/github.com/yumaojun03/dmidecode)](https://goreportcard.com/report/github.com/yumaojun03/dmidecode)
[![Release](https://img.shields.io/github/release/yumaojun03/dmidecode.svg?style=flat-square)](https://github.com/yumaojun03/dmidecode/releases)
[![MIT License](https://img.shields.io/github/license/yumaojun03/dmidecode.svg)](https://github.com/yumaojun03/dmidecode/blob/master/LICENSE)

纯Golang实现的dmidecode, 零依赖, 支持Linux, Unix, Windows

功能和命令行的dmidecode工具一样, 使用方式参考: [example](./example/main.go)


## 安装方式

```
$ go get "github.com/yumaojun03/dmidecode"
```

## 使用样例

``` go
package main

import (
	"fmt"
	"os"

	"github.com/yumaojun03/dmidecode"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	dmi, err := dmidecode.New()
	checkError(err)

	infos, err := dmi.BIOS()
	// 支持以下类型的解析
	// dmi.BaseBoard()
	// dmi.Chassis()
	// dmi.MemoryArray()
	// dmi.MemoryDevice()
	// dmi.Onboard()
	// dmi.PortConnector()
	// dmi.Processor()
	// dmi.ProcessorCache()
	// dmi.Slot()
	// dmi.System()
	checkError(err)

	for i := range infos {
		fmt.Println(infos[i])
	}
}
```

## CLI 使用
``` sh
$ go run cmd/main.go -t [bios, system, baseboard, chassis, onboard, port, processor, memory, slot]
```