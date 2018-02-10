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


func (j *jsonData) Get(search string) interface{} {
	return j.data.Path(search).Data()
}
