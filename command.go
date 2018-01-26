package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/leehuk/golang-clicommand"
)

var (
	// maintain a global pointer to auth create as it has a special exemption from
	// apitoken requirements, and we dont want to use fragile reflection to compare pointers
	cliAuthCreatePtr *clicommand.Command
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

	// ghcli auth: common arguments
	cliAuthArgNote := clicommand.NewArg("note", "Description of oauth token purpose", true)
	cliAuthArgScopes := clicommand.NewArg("scopes", "Comma separated list of scopes", true)

	// ghcli auth create
	//
	// This has a special exception to apitoken requirements tracked via cliAuthCreatePtr,
	// as it supports password+mfatoken to create api tokens, which also has its own
	// verification callback.
	cliAuthCreate := cliAuth.NewCommand("create", "Create OAuth Token", command_auth_create)
	cliAuthCreatePtr = cliAuthCreate
	cliAuthCreate.BindCallback(cmd_cb_validate_creds_password)
	cliAuthCreate.NewArg("password", "Password for github.com, or use ENV GHAPI_PASSWORD", true)
	cliAuthCreate.NewArg("mfatoken", "MFA Token (e.g. Auth App) for github.com, or use ENV GHAPI_MFATOKEN", true)
	cliAuthCreate.BindArg(cliAuthArgNote, cliAuthArgScopes)

	// ghcli auth get
	cliAuth.NewCommand("get", "Get OAuth Token Details", command_auth_get)

	// ghcli auth list
	cliAuth.NewCommand("list", "List OAuth Tokens", command_auth_list)

	return cliRoot
}

func cmd_cb_validate_creds(data *clicommand.Data) error {
	for k, t := range map[string]string{"GHAPI_USERNAME":"username", "GHAPI_APITOKEN":"apitoken"} {
		if v := os.Getenv(k); v != "" {
			data.Options[t] = v
		}
	}

	if _, ok := data.Options["username"]; !ok {
		return errors.New("Required option missing: username")
	}

	// apitoken is required, except for when calling auth create which can use a password
	if _, ok := data.Options["apitoken"]; !ok && cliAuthCreatePtr != data.Cmd {
		return fmt.Errorf("Required option missing: apitoken")
	}

	return nil
}

func cmd_cb_validate_creds_password(data *clicommand.Data) error {
	for k, t := range map[string]string{"GHAPI_PASSWORD":"password", "GHAPI_MFATOKEN":"mfatoken"} {
		if v := os.Getenv(k); v != "" {
			data.Options[t] = v
		}
	}

	_, tokok := data.Options["apitoken"]
	_, pwdok := data.Options["password"]

	if tokok && pwdok {
		return fmt.Errorf("Conflicting options: apitoken and password")
	} else if !tokok && !pwdok {
		return fmt.Errorf("Required option missing: apitoken|password")
	}

	return nil
}
