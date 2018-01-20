package main

import(
    "github.com/leehuk/golang-clicommand"
)

func cmd_init_auth(parent *clicommand.Command) {
    cli_auth := clicommand.New("auth", "Manage OAuth Access", parent, nil)
    clicommand.New("create", "Create OAuth Token", cli_auth, command_auth_create)
    clicommand.New("create", "Get OAuth Token Details", cli_auth, command_auth_get)
}

func command_auth_create(params []string, args map[string]string) error {
    return nil
}

func command_auth_get(params []string, args map[string]string) error {
    return nil
}
