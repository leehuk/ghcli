package main

import (
	"errors"
	"os"

	"github.com/leehuk/golang-clicommand"
)

func cmd_init() *clicommand.Command {
	// root object
	cliRoot := clicommand.NewCommand("ghcli", "CLI tool for accessing the github.com API", nil)

	// global callbacks
	cliRoot.BindCallback(cmd_cb_validate_creds)

	// global parameters
	cliRoot.NewArg("oj", "Output in JSON", false)
	cliRoot.NewArg("os", "Output in simple parseable form", false)

	cliRoot.NewArg("username", "Username for github.com, or use ENV GHAPI_USERNAME", true)
	cliRoot.NewArg("password", "Password for github.com, or use ENV GHAPI_PASSWORD", true)
	cliRoot.NewArg("mfatoken", "MFA Token (e.g. Auth App) for github.com, or use ENV GHAPI_MFATOKEN", true)
	cliRoot.NewArg("apitoken", "API Token for github.com, or use ENV GHAPI_APITOKEN", true)

	// ghcli auth
	cliAuth := cliRoot.NewCommand("auth", "Manage OAuth Access", nil)

	// ghcli auth create
	cliAuth.NewCommand("create", "Create OAuth Token", command_auth_create)

	// ghcli auth get
	cliAuth.NewCommand("get", "Get OAuth Token Details", command_auth_get)

	return cliRoot
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
