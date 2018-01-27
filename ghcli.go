package main

import (
	"os"
)

func main() {
	cli_root := cmd_init()

	if error := cli_root.Parse(); error != nil {
		os.Exit(1)
	}
}
