package smbios

import (
	"fmt"
)

// ParseRecovery todo
func ParseRecovery(s *Structure, err *error) {
	if rc := recover(); rc != nil {
		*err = fmt.Errorf("parse structure (%s)  pannic: %v", s, rc)
	}
}
