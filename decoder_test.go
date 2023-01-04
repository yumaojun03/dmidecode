package dmidecode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yumaojun03/dmidecode"
)

func TestSystemPowerSupply(t *testing.T) {
	dmi, err := dmidecode.New()
	assert.Equal(t, err, nil)

	info, err := dmi.SystemPowerSupply()
	assert.Equal(t, err, nil)

	fmt.Println(info[0].String())
}
