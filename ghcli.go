package main

import(
    "fmt"
    "github.com/leehuk/golang-clicommand"
    "os"
)

func main() {
    cli_root := cmd_init()

    if error := clicommand.Parse(cli_root); error != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", error)
        os.Exit(1)
    }
}
