package main

import(
    "github.com/leehuk/golang-clicommand"
)

func cmd_init() *clicommand.Command {
    cli_root := clicommand.New("gh", "CLI tool for accessing the github.com API", nil, nil)
    clicommand.NewArg(cli_root, "oj", false, "Output in JSON")
    clicommand.NewArg(cli_root, "os", false, "Output in simple parseable form")

    cmd_init_auth(cli_root)

    return cli_root
}
