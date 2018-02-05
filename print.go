package main

import (
	"encoding/json"
	"fmt"

	"github.com/leehuk/go-clicommand"
)

func ghPrint(data interface{}, params *clicommand.Data) {
	if _, ok := params.Options["ob"]; ok {
		dataj, _ := json.MarshalIndent(data, "", "  ")
		fmt.Printf("%s\n", dataj)
	} else {
		dataj, _ := json.Marshal(data)
		fmt.Printf("%s\n", dataj)
	}
}
