package main

import (
	"github.com/Jeffail/gabs"
)

type jsonData struct {
	data *gabs.Container
}

func jsonParse(datab []byte) (*jsonData, error) {
	data, err := gabs.ParseJSON(datab)
	if err != nil {
		return nil, err
	}

	return &jsonData{data}, nil
}

func (j *jsonData) get() interface{} {
	return j.data.Data()
}

func (j *jsonData) gets(search string) interface{} {
	return j.data.Path(search).Data()
}

func (j *jsonData) getstr() string {
	return j.data.String()
}
