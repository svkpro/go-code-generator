package template

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type MainTpl struct {
	PackageName string `json:"packageName"`
	ServiceName string `json:"serviceName"`
}

func NewGoTpl() *MainTpl {
	wd, err := os.Getwd()

	tj, err := filepath.Abs(wd + "/templates/cmd/main.json")
	if err != nil {
		panic(err)
	}

	jsonConfig, err := os.Open(tj)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := jsonConfig.Close(); err != nil {
			panic(err)
		}
	}()

	data, err := ioutil.ReadAll(jsonConfig)
	if err != nil {
		panic(err)
	}

	t := &MainTpl{}
	err = json.Unmarshal(data, t)
	if err != nil {
		panic(err)
	}

	return t
}
