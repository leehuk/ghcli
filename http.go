package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	}

	httpreq, err := http.NewRequest(method, "https://api.github.com"+api, datab)
	if err != nil {
		return  nil, err
	}

	if data != nil {
		httpreq.Header.Set("Content-Type", "application/json")
	}

	return httpreq, nil
}

func ghHttpExecRequest(httpclient *http.Client, httpreq *http.Request) (interface{}, error) {
	httpresp, err := httpclient.Do(httpreq)
	if err != nil {
		return nil, err
	}

	defer httpresp.Body.Close()

	httpbody, err := ioutil.ReadAll(httpresp.Body)
	if err != nil {
		return nil, err
	}

	var resdata interface{}

	err = json.Unmarshal(httpbody, &resdata)
	if err != nil {
		return nil, err
	}

	return resdata, nil
}

func ghhttpToken(method string, api string, data map[string]interface{}, options map[string]string) (interface{}, error) {
	httpclient, err := ghHttpNewClient()
	if err != nil {
		return nil, err
	}

	httpreq, err := ghHttpNewRequest(method, api, data)
	if err != nil {
		return nil, err
	}

	httpreq.Header.Set("Authorization", fmt.Sprintf("token %s", options["apitoken"]))

	return ghHttpExecRequest(httpclient, httpreq)
}

func ghhttpBasic(method string, api string, data map[string]interface{}, options map[string]string) (interface{}, error) {
	httpclient, err := ghHttpNewClient()
	if err != nil {
		return nil, err
	}

	httpreq, err := ghHttpNewRequest(method, api, data)
	if err != nil {
		return nil, err
	}

	httpreq.SetBasicAuth(options["username"], options["password"])

	return ghHttpExecRequest(httpclient, httpreq)

}
