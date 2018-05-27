package main

import (
	"fmt"
	"os"

	"github.com/leehuk/go-clicommand"
)

var (
	// maintain a global pointer to auth parent menu as it has a special exemption from
	// apitoken requirements, and we dont want to use reflection to compare pointers as
	// its potentially unstable
	cliAuthPtr *clicommand.Command
)

func cmd_init() *clicommand.Command {
	// ghcli
	cliRoot := clicommand.NewCommand("ghcli", "CLI tool for accessing the github.com API", nil)
	cliRoot.BindCallbackPre(cmd_cb_env_translate)
	cliRoot.NewOption("apitoken", "API Token for github.com, or use ENV GHAPI_APITOKEN", true).SetRequired()

	cliRoot.NewOption("ob", "Output in beautified json", false)

	// ghcli auth
	cliAuth := cliRoot.NewCommand("auth", "Manage OAuth Access", nil)
	// The auth api only supports basic auth using username+password, so carve out
	// its exception to apitoken requirements
	cliAuthPtr = cliAuth
	cliAuth.BindCallbackPre(cmd_cb_env_translate_auth)
	cliAuth.BindCallback(cmd_cb_validate_creds_auth)

	cliAuthOptionNote := clicommand.NewOption("note", "Description of oauth token purpose", true).SetRequired()
	cliAuthOptionScopes := clicommand.NewOption("scopes", "Comma separated list of scopes", true)

	// ghcli auth create
	cliAuthCreate := cliAuth.NewCommand("create", "Create OAuth Token", command_auth_create)
	cliAuthCreate.BindOption(cliAuthOptionNote, cliAuthOptionScopes)
	cliAuthCreate.NewOption("username", "Username for github.com, or use ENV GHAPI_USERNAME", true).SetRequired()
	cliAuthCreate.NewOption("password", "Password for github.com, or use ENV GHAPI_PASSWORD", true).SetRequired()
	cliAuthCreate.NewOption("mfatoken", "MFA Token (e.g. Auth App) for github.com, or use ENV GHAPI_MFATOKEN", true)
	cliAuthCreate.NewOption("os", "Output the API token only", false)

	// ghcli auth get
	cliAuth.NewCommand("get", "Get OAuth Token Details", command_auth_get)

	// ghcli auth list
	cliAuth.NewCommand("list", "List OAuth Tokens", command_auth_list)

	return cliRoot
}

func cmd_cb_env_translate(data *clicommand.Data) error {
	for k, t := range map[string]string{"GHAPI_APITOKEN": "apitoken"} {
		if v := os.Getenv(k); v != "" {
			data.Options[t] = v
		}
	}

	return nil
}

func cmd_cb_env_translate_auth(data *clicommand.Data) error {
	for k, t := range map[string]string{"GHAPI_USERNAME": "username", "GHAPI_PASSWORD": "password", "GHAPI_MFATOKEN": "mfatoken"} {
		if v := os.Getenv(k); v != "" {
			data.Options[t] = v
		}
	}

	if arg := data.Cmd.GetOption("apitoken", true); arg != nil {
		for _, cmd := range arg.GetParents() {
			cmd.UnbindOption(arg)
		}
	}

	return nil
}

func cmd_cb_validate_creds_auth(data *clicommand.Data) error {

	_, tokok := data.Options["apitoken"]
	_, usrok := data.Options["username"]
	_, pwdok := data.Options["password"]

	if tokok && (usrok || pwdok) {
		return fmt.Errorf("Conflicting options: apitoken && username,password")
	} else if !tokok && (!usrok || !pwdok) {
		return fmt.Errorf("Required option missing: apitoken || username,password")
	}

	return nil
}
