package main

import(
    "github.com/leehuk/golang-clicommand"
)

func cmd_init() *clicommand.CLICommand {
    cmd_root := clicommand.New("gh", "CLI tool for accessing the github.com API")

    cmd_root.AddArg("oj", "Output in JSON", false)
    cmd_root.AddArg("os", "Output in simple parseable form", false)

    cmd_auth := cmd_root.AddMenu("auth", "Manage OAuth Access", nil)
    cmd_auth.AddMenu("create", "Create OAuth Token", command_auth_create)
    cmd_auth.AddMenu("create", "Get OAuth Token Details", command_auth_get)

    return cmd_root
}
