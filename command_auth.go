package main

import (
	"fmt"

	"github.com/leehuk/go-clicommand"
)

func command_auth_create(params *clicommand.Data) error {
	var postdata = make(map[string]interface{})

	postdata["note"] = params.Options["note"]

	if data, err := ghHttp("POST", "/authorizations", postdata, params.Options); err == nil {
		if _, ok := params.Options["os"]; ok {
			token := data.data.Path("token").String()
			if token != "" {
				fmt.Printf("%s", token)
			} else {
				return fmt.Errorf("Unable to decode token field")
			}
		} else {
			ghPrint(data, params)
		}
	} else {
		return err
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
