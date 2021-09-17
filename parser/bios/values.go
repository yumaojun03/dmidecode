package bios

import (
	"fmt"
)

// RuntimeSize todo
type RuntimeSize uint

func (b RuntimeSize) String() string {
	if (b & 0x3FF) > 0 {
		return fmt.Sprintf("%d Bytes", b)
	}
	return fmt.Sprintf("%d kB", b>>10)
}

// RomSize todo
type RomSize uint

func (b RomSize) String() string {
	return fmt.Sprintf("%d kB", b)
}
