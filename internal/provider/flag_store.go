package provider

var StringFlagValues = make(map[string]StringFlagValue)
var BoolFlagValues = make(map[string]BoolFlagValue)
var FloatFlagValues = make(map[string]FloatFlagValue)
var IntFlagValues = make(map[string]IntFlagValue)

func PopulateFlagValues() {
	StringFlagValues["dataplane_generation"] = StringFlagValue{FlagKey: "dataplane_generation", FlagValue: "metal.v1"}
	BoolFlagValues["grant_soil_access"] = BoolFlagValue{FlagKey: "grant_soil_access", FlagValue: false}
	FloatFlagValues["special_ability_buff_perc"] = FloatFlagValue{FlagKey: "special_ability_buff_perc", FlagValue: 0.23456}
	IntFlagValues["num_of_special_abilities"] = IntFlagValue{FlagKey: "num_of_special_abilities", FlagValue: 12}
}
