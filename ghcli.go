package main

import (
	"fmt"
	"os"
)

func main() {
	cli_root := cmd_init()

	if error := cli_root.Parse(); error != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", error)
		os.Exit(1)
	}
}
