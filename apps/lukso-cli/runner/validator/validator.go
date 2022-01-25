package validator

import "lukso-cli/config"

func CheckValidatorRequirements(LuksoSettings config.LuksoValues) {
	println("checking")
}

func Start(LuksoSettings *config.LuksoValues) {
	println("Starting Validator")
}

func Stop(LuksoSettings *config.LuksoValues) {
	println("Stopping Validator")
}
