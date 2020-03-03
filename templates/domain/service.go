package template

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ServiceTpl struct {
	PackageName string `json:"packageName"`
	ServiceName string `json:"serviceName"`
}

func NewGoTpl() *ServiceTpl {
	wd, err := os.Getwd()

	tj, err := filepath.Abs(wd + "/templates/domain/service.json")
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

	t := &ServiceTpl{}
	err = json.Unmarshal(data, t)
	if err != nil {
		panic(err)
	}

	return t
}
