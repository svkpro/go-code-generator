package template

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ConfigGoTpl struct {
	PackageName string `json:"packageName"`
	ServiceName string `json:"serviceName"`
	HttpPort    string `json:"httpPort"`
}

type ConfigJsonTpl struct {
	HTTP_PORT string
}

func NewGoTpl() *ConfigGoTpl {
	wd, err := os.Getwd()

	tj, err := filepath.Abs(wd + "/templates/config/config.json")
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

	t := &ConfigGoTpl{}
	err = json.Unmarshal(data, t)
	if err != nil {
		panic(err)
	}

	return t
}

func NewJsonTpl() *ConfigJsonTpl {
	return &ConfigJsonTpl{
		HTTP_PORT: "19001",
	}
}
