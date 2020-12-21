package battery

type DeviceChemistry byte

const (
	Other DeviceChemistry = 1 << iota
	Unknown
	LeadAcid
	NickelCadmium
	NickelMetalHydride
	LithiumIon
	ZincAir
	LithiumPolymer
)

func (d DeviceChemistry) String() string {
	factors := [...]string{
		"Other",
		"Unknown",
		"Lead Acid",
		"Nickel Cadmium",
		"Nickel metal hydride",
		"Lithium-ion",
		"Zinc air",
		"Lithium Polymer",
	}
	return factors[d]
}
