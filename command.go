package main

import (
	"os"

	"github.com/leehuk/go-clicommand"
)

var (
	cliRoot *clicommand.Command
)

func cmd_init() *clicommand.Command {
	// ghcli
	cliRoot = clicommand.NewCommand("ghcli", "CLI tool for accessing the github.com API", nil)
	cliRoot.BindCallbackPre(cmd_cb_env_translate)
	//cliRoot.NewOption("apitoken", "API Token for github.com, or use ENV GHAPI_APITOKEN", true).SetRequired()

	cliRoot.NewOption("ob", "Output in beautified json", false)

	cmd_init_auth()

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
