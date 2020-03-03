package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

const (
	tplPath  = "templates/%s"
	codePath = "%s/code/%s"
)

func MakeDir(dn string) {
	_, err := os.Stat(dn)
	if err != nil {
		os.Mkdir(dn, 0777)
	} else {
		os.Remove(dn + "/*")
	}
}

func GenerateGo(dn string, pn string, gt interface{}) {
	wd, err := os.Getwd()
	Die(err)
	cwd := fmt.Sprintf(codePath, wd, dn)
	MakeDir(cwd)

	gf, err := os.Create(fmt.Sprintf("%s/%s.go", cwd, pn))
	Die(err)
	defer gf.Close()

	tplContent, err := ioutil.ReadFile(fmt.Sprintf(tplPath, fmt.Sprintf("%s/%s.tpl", dn, pn)))
	Die(err)

	var tpl = template.Must(template.New(fmt.Sprintf(codePath, "", pn)).Parse(string(tplContent)))
	tpl.Execute(gf, gt)
}

func GenerateJson(dn string, pn string, jt interface{}) {
	wd, err := os.Getwd()
	Die(err)
	cwd := fmt.Sprintf(codePath, wd, dn)
	MakeDir(cwd)

	rankingsJson, err := json.Marshal(jt)
	Die(err)

	err = ioutil.WriteFile(fmt.Sprintf("%s/%s.json", cwd, pn), rankingsJson, 0775)
}
