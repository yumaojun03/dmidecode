package system

// WakeUpType 唤醒类型
type WakeUpType byte

const (
	Reserved WakeUpType = iota
	Other
	Unknown
	APMTimer
	ModemRing
	LANRemote
	PoAerSwitch
	PCIPME
	ACPowerRestored
)

func (w WakeUpType) String() string {
	types := [...]string{
		"Reserved", /* 0x00 */
		"Other",
		"Unknown",
		"APM Timer",
		"Modem Ring",
		"LAN Remote",
		"Power Switch",
		"PCI PME#",
		"AC Power Restored", /* 0x08 */
	}
	return types[w]
}
