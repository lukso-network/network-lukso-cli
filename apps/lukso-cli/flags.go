package main

import (
	"flag"
	"fmt"
)


func initFlags() {

	help := map[string]string{
		"help"
		"pandora": "sets pandora tag to be used",
	}


	var nFlag = flag.Int("pandora", 1234, help["pandora"])
	var nFlag = flag.Int("pandora", 1234, help["config"])
	flag.Parse()
	fmt.Println(*nFlag)
}
