package main

import(
    "github.com/leehuk/golang-clicommand"
)

func cmd_init() *clicommand.Command {
    cli_root := clicommand.New("gh", "CLI tool for accessing the github.com API", nil, nil)
    clicommand.NewArg(cli_root, "oj", false, "Output in JSON")
    clicommand.NewArg(cli_root, "os", false, "Output in simple parseable form")

    cli_auth := clicommand.New("auth", "Manage OAuth Access", cli_root, nil)
    clicommand.New("create", "Create OAuth Token", cli_auth, command_auth_create)
    clicommand.New("create", "Get OAuth Token Details", cli_auth, command_auth_get)

    return cli_root
}
