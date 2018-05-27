package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/leehuk/go-clicommand"
)

func ghHttpNewClient() (*http.Client, error) {
	return &http.Client{}, nil
}

func ghHttpNewRequest(method string, api string, data map[string]interface{}) (*http.Request, error) {
	var datab *bytes.Buffer

	if data != nil {
		dataJson, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		datab = bytes.NewBuffer(dataJson)
	} else {
		datab = new(bytes.Buffer)
	}

	httpreq, err := http.NewRequest(method, "https://api.github.com"+api, datab)
	if err != nil {
		return nil, err
	}

	if data != nil {
		httpreq.Header.Set("Content-Type", "application/json")
	}

	return httpreq, nil
}

func ghHttpExecRequest(httpclient *http.Client, httpreq *http.Request) (map[string]interface{}, error) {
	httpresp, err := httpclient.Do(httpreq)
	if err != nil {
		return nil, err
	}

	defer httpresp.Body.Close()

	httpbody, err := ioutil.ReadAll(httpresp.Body)
	if err != nil {
		return nil, err
	}

	var jdata map[string]interface{}
	json.Unmarshal(httpbody, &jdata)

	if httpresp.StatusCode >= 400 {
		if message, ok := jdata["message"].(string); ok {
			return nil, fmt.Errorf("%d %s", httpresp.StatusCode, message)
		}

		return nil, fmt.Errorf("%s", httpresp.Status)
	}

	return jdata, nil
}

func ghHttp(method string, api string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	httpclient, err := ghHttpNewClient()
	if err != nil {
		return nil, err
	}

	httpreq, err := ghHttpNewRequest(method, api, data)
	if err != nil {
		return nil, err
	}

	_, userok := options["username"]
	_, passok := options["password"]
	_, mfaok := options["mfatoken"]

	if userok && passok {
		httpreq.SetBasicAuth(options["username"], options["password"])

		if mfaok {
			httpreq.Header.Set("X-GitHub-OTP", fmt.Sprintf("%v", options["mfatoken"]))
		}
	} else {
		httpreq.Header.Set("Authorization", fmt.Sprintf("token %s", options["apitoken"]))
	}

	return ghHttpExecRequest(httpclient, httpreq)
}

func ghPrint(data map[string]interface{}, params *clicommand.Data) {
	if _, ok := params.Options["ob"]; ok {
		dataj, _ := json.MarshalIndent(data, "", "  ")
		fmt.Printf("%s\n", dataj)
	} else {
		dataj, _ := json.Marshal(data)
		fmt.Printf("%s\n", dataj)
	}
}
