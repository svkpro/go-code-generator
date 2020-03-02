package helpers

import (
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

func Generate(dn string, pn string, t interface{}) {
	wd, err := os.Getwd()
	die(err)
	cwd := fmt.Sprintf(codePath, wd, dn)
	MakeDir(cwd)

	gf, err := os.Create(fmt.Sprintf("%s/%s.go", cwd, pn))
	die(err)
	defer gf.Close()

	tplContent, err := ioutil.ReadFile(fmt.Sprintf(tplPath, fmt.Sprintf("%s/%s.tpl", pn, pn)))
	die(err)

	var tpl = template.Must(template.New(fmt.Sprintf(codePath, "", pn)).Parse(string(tplContent)))
	tpl.Execute(gf, t)
}
