package main

import(
    "errors"
    "os"

    "github.com/leehuk/golang-clicommand"
)

func cmd_init() *clicommand.CLICommand {
    cli_root := clicommand.New("ghcli", "CLI tool for accessing the github.com API")

    cli_root.AddCallback(cmd_cb_validate_creds)

    cli_root.AddArg("oj", "Output in JSON", false)
    cli_root.AddArg("os", "Output in simple parseable form", false)

    cli_root.AddArg("username", "Username for github.com, or use ENV GHAPI_USERNAME", true)
    cli_root.AddArg("password", "Password for github.com, or use ENV GHAPI_PASSWORD", true)
//    cli_root.AddArgNote("password", "Only one of password or apitoken must be specified")
    cli_root.AddArg("mfatoken", "MFA Token (e.g. Auth App) for github.com, or use ENV GHAPI_MFATOKEN", true)
    cli_root.AddArg("apitoken", "API Token for github.com, or use ENV GHAPI_APITOKEN", true)

    cli_auth := cli_root.AddMenu("auth", "Manage OAuth Access", nil)
    cli_auth.AddMenu("create", "Create OAuth Token", command_auth_create)
    cli_auth.AddMenu("get", "Get OAuth Token Details", command_auth_get)

    return cli_root
}

func cmd_cb_validate_creds(data *clicommand.CLICommandData) error {
    for _, k := range []string{"GHAPI_USERNAME","GHAPI_PASSWORD","GHAPI_MFATOKEN","GHAPI_APITOKEN"} {
        if v := os.Getenv(k); v != "" {
            data.Options[k] = v
        }
    }

    if _, ok := data.Options["username"]; !ok {
        return errors.New("Required option missing: username")
    }

    _, haspwd := data.Options["password"]
    _, hastok := data.Options["apitoken"]

    if haspwd && hastok {
        return errors.New("Only one of password or apitoken may be specified, both provided")
    } else if !haspwd && !hastok {
        return errors.New("Required option missing: password|apitoken")
    }

    return nil
}
