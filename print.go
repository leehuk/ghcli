package main

import (
	"encoding/json"
	"fmt"

	"github.com/leehuk/go-clicommand"
)

func ghPrint(data *jsonData, params *clicommand.Data) {
	if _, ok := params.Options["ob"]; ok {
		dataj, _ := json.MarshalIndent(data.get(), "", "  ")
		fmt.Printf("%s\n", dataj)
	} else {
		fmt.Printf("%s\n", data.getstr())
	}
}
