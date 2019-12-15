# dmidecode

纯Golang实现的dmidecode, 零依赖, 支持Linux, Unix, Windows


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
	checkError(err)

	for i := range infos {
		fmt.Println(infos[i])
	}
}

```