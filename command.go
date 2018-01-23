package main

import(
    "github.com/leehuk/golang-clicommand"
)

func cmd_init() *clicommand.CLICommand {
    cli_root := clicommand.New("ghcli", "CLI tool for accessing the github.com API")

    cli_root.AddArg("oj", "Output in JSON", false)
    cli_root.AddArg("os", "Output in simple parseable form", false)

    cli_root.AddArg("username", "Username for github.com, or use ENV GHAPI_USERNAME", true)
    cli_root.AddArg("password", "Password for github.com, or use ENV GHAPI_PASSWORD", true)
    cli_root.AddArg("mfatoken", "MFA Token for github.com, or use ENV GHAPI_MFATOKEN", true)
    cli_root.AddArg("apitoken", "API Token for github.com, or use ENV GHAPI_APITOKEN", true)

    cli_auth := cli_root.AddMenu("auth", "Manage OAuth Access", nil)
    cli_auth.AddMenu("create", "Create OAuth Token", command_auth_create)
    cli_auth.AddMenu("get", "Get OAuth Token Details", command_auth_get)

    return cli_root
}
