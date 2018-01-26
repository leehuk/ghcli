package main

import (
	"errors"
	"os"

	"github.com/leehuk/golang-clicommand"
)

func cmd_init() *clicommand.Command {
	cli_root := clicommand.NewCommand("ghcli", "CLI tool for accessing the github.com API", nil)

	cli_root.BindCallback(cmd_cb_validate_creds)

	cli_root.NewArg("oj", "Output in JSON", false)
	cli_root.NewArg("os", "Output in simple parseable form", false)

	cli_root.NewArg("username", "Username for github.com, or use ENV GHAPI_USERNAME", true)
	cli_root.NewArg("password", "Password for github.com, or use ENV GHAPI_PASSWORD", true)
	cli_root.NewArg("mfatoken", "MFA Token (e.g. Auth App) for github.com, or use ENV GHAPI_MFATOKEN", true)
	cli_root.NewArg("apitoken", "API Token for github.com, or use ENV GHAPI_APITOKEN", true)

	cli_auth := cli_root.NewCommand("auth", "Manage OAuth Access", nil)
	cli_auth.NewCommand("create", "Create OAuth Token", command_auth_create)
	cli_auth.NewCommand("get", "Get OAuth Token Details", command_auth_get)

	return cli_root
}

func cmd_cb_validate_creds(data *clicommand.Data) error {
	for _, k := range []string{"GHAPI_USERNAME", "GHAPI_PASSWORD", "GHAPI_MFATOKEN", "GHAPI_APITOKEN"} {
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
