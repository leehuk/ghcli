package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"github.com/leehuk/go-clicommand"
	"golang.org/x/crypto/ssh/terminal"
)

func cmd_init_auth() {
	// ghcli auth
	cliAuth := cliRoot.NewCommand("auth", "Manage OAuth Access", nil)
	// The auth api only supports basic auth using username+password.
	// We do support loading these from stdin, or environment alternatively though.
	cliAuth.BindCallbackPre(cmd_cb_env_translate_auth)
	cliAuth.BindCallback(cmd_cb_validate_creds_auth)

	cliAuth.NewOption("i", "Interactive prompt for auth credentials", false)
	cliAuth.NewOption("username", "Username for github.com, or use ENV GHAPI_USERNAME.  Defaults to interactive request.", true)
	cliAuth.NewOption("password", "Password for github.com, or use ENV GHAPI_PASSWORD.  Defaults to interactive request.", true)
	cliAuth.NewOption("mfatoken", "MFA Token for github.com, or use ENV GHAPI_MFATOKEN.  Defaults to interactive request.", true)

	cliAuthOptionNote := clicommand.NewOption("note", "Description of oauth token purpose", true).SetRequired()
	cliAuthOptionScopes := clicommand.NewOption("scopes", "Comma separated list of scopes", true)

	// ghcli auth create
	cliAuthCreate := cliAuth.NewCommand("create", "Create OAuth Token", command_auth_create)
	cliAuthCreate.BindOption(cliAuthOptionNote, cliAuthOptionScopes)
	cliAuthCreate.NewOption("os", "Output the API token only", false)

	// ghcli auth get
	cliAuth.NewCommand("get", "Get OAuth Token Details", command_auth_get)

	// ghcli auth list
	cliAuth.NewCommand("list", "List OAuth Tokens", command_auth_list)
}

func cmd_cb_env_translate_auth(data *clicommand.Data) error {
	for k, t := range map[string]string{"GHAPI_USERNAME": "username", "GHAPI_PASSWORD": "password", "GHAPI_MFATOKEN": "mfatoken"} {
		if v := os.Getenv(k); v != "" {
			data.Options[t] = v
		}
	}

	return nil
}

func cmd_cb_validate_creds_auth(data *clicommand.Data) error {
	_, usrok := data.Options["username"]
	_, pwdok := data.Options["password"]
	_, mfaok := data.Options["password"]
	_, intok := data.Options["i"]

	if usrok && pwdok && mfaok {
		return nil
	} else if intok {
		reader := bufio.NewReader(os.Stdin)

		if !usrok {
			fmt.Print("Enter username: ")
			data.Options["username"], _ = reader.ReadString('\n')
		}

		if !pwdok {
			fmt.Print("Enter password: ")
			password, _ := terminal.ReadPassword(int(syscall.Stdin))
			data.Options["password"] = string(password)
			fmt.Print("\n")
		}

		if !mfaok {
			fmt.Print("Enter mfatoken: ")
			mfatoken, _ := terminal.ReadPassword(int(syscall.Stdin))
			fmt.Print("\n")

			if string(mfatoken) != "" {
				data.Options["mfatoken"] = string(mfatoken)
			}
		}
	} else {
		return fmt.Errorf("Required option missing: username,password")
	}


	return nil
}

func command_auth_create(params *clicommand.Data) error {
	var postdata = make(map[string]interface{})

	postdata["note"] = params.Options["note"]

	data, err := ghHttp("POST", "/authorizations", postdata, params.Options)
	if err != nil {
		return err
	}

	if _, ok := params.Options["os"]; ok {
		token, ok := data["token"].(string)
		if ok {
			fmt.Printf("%s", token)
		} else {
			return fmt.Errorf("Unable to decode token field")
		}
	} else {
		ghPrint(data, params)
	}
	return nil
}

func command_auth_get(params *clicommand.Data) error {
	url := "/authorizations/" + params.Params[0]

	if data, err := ghHttp("GET", url, nil, params.Options); err == nil {
		ghPrint(data, params)
	} else {
		return err
	}

	return nil
}

func command_auth_list(params *clicommand.Data) error {
	var postdata = make(map[string]interface{})

	if data, err := ghHttp("GET", "/authorizations", postdata, params.Options); err == nil {
		ghPrint(data, params)
	} else {
		return err
	}

	return nil
}
