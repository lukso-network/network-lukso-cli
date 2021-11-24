package main

import (
	"fmt"
	"reflect"

	"lukso-cli/config"
	"lukso-cli/flags"
)

var LuksoSettings config.LuksoValues

func main() {
	flags.InitFlags()

	if flags.FlagValues.Config != "" {
		println("Config loaded")
		config.LoadConfig(flags.FlagValues.Config)
	}

	config.LoadDefaults()

	// Build Settings

	v := reflect.ValueOf(LuksoSettings)
	typeOfS := v.Type()

	// Awful but works
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())

		r := reflect.ValueOf(&flags.FlagValues)
		f := reflect.Indirect(r).FieldByName(typeOfS.Field(i).Name)
		println(f.String())

		if f.String() != "" {
			println(flags.FlagValues.Network)
		} else if config.ConfigValues.Network != "" {
			println(config.ConfigValues.Network)
		} else {
			println(config.DefaultValues.Network)
		}

	}

}
