package main

import (
	"fmt"
	"go-code-generator/helpers"
	gcgc "go-code-generator/templates/config"
	gcgm "go-code-generator/templates/main"
	"io/ioutil"
	"os"
	"text/template"
)

const (
	tplPath        = "templates/%s"
	codePath       = "%s/code/%s"
	successMessage = "The code has been generated successfully!"
)

func main() {
	generate("cmd", "main", gcgm.New())
	generate("config", "config", gcgc.New())

	fmt.Println(successMessage)
}

func generate(dn string, pn string, t interface{}) {
	wd, err := os.Getwd()
	die(err)
	cwd := fmt.Sprintf(codePath, wd, dn)
	helpers.MakeDir(cwd)
	gf, err := os.Create(fmt.Sprintf("%s/%s.go", cwd, pn))
	die(err)
	defer gf.Close()

	tplContent, err := ioutil.ReadFile(fmt.Sprintf(tplPath, fmt.Sprintf("%s/%s.tpl", pn, pn)))
	die(err)

	var tpl = template.Must(template.New(fmt.Sprintf(codePath, "", pn)).Parse(string(tplContent)))
	tpl.Execute(gf, t)
}
