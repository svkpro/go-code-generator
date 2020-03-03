package {{.PackageName}}

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Settings struct {
	HttpPort string `json:"{{.HttpPort}}" env:"{{.HttpPort}}"`
}

func New() *Settings {
	wd, err := os.Getwd()

	tj, err := filepath.Abs(wd + "/config/config.json")
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

	t := &Settings{}
	err = json.Unmarshal(data, t)
	if err != nil {
		panic(err)
	}

	return t
}