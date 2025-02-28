package provider

type StringFlagValue struct {
	FlagKey   string
	FlagValue string
}

type BoolFlagValue struct {
	FlagKey   string
	FlagValue bool
}

type FloatFlagValue struct {
	FlagKey   string
	FlagValue float64
}

type IntFlagValue struct {
	FlagKey   string
	FlagValue int64
}
