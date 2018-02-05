package main

import (
	"github.com/leehuk/go-clicommand"
)

func command_auth_create(params *clicommand.Data) error {
	var postdata = make(map[string]interface{})

	postdata["note"] = params.Options["note"]

	if data, err := ghhttpBasic("POST", "/authorizations", postdata, params.Options); err == nil {
		dataj, _ := json.Marshal(data)
		fmt.Printf("%s\n", dataj)
	} else {
		return err
	}

	return nil
}

func command_auth_get(params *clicommand.Data) error {
	return nil
}

func command_auth_list(params *clicommand.Data) error {
	var postdata = make(map[string]interface{})

	if data, err := ghhttpBasic("GET", "/authorizations", postdata, params.Options); err == nil {
		dataj, _ := json.Marshal(data)
		fmt.Printf("%s\n", dataj)
	} else {
		return err
	}

	return nil
}
