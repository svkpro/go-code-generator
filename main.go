package main

import (
	"fmt"
	"go-code-generator/helpers"
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
	wd, err := os.Getwd()
	die(err)
	mwd := fmt.Sprintf(codePath, wd, "cmd")
	helpers.MakeDir(mwd)
	gf, err := os.Create(fmt.Sprintf("%s/main.go", mwd))
	die(err)
	defer gf.Close()

	tplContent, err := ioutil.ReadFile(fmt.Sprintf(tplPath, "main/main.tpl"))
	die(err)

	var tpl = template.Must(template.New(fmt.Sprintf(codePath, "", "main")).Parse(string(tplContent)))

	tpl.Execute(gf, gcgm.New())

	fmt.Println(successMessage)
}
