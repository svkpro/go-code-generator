package main

import (
	"fmt"
	gcg "go-code-generator/templates/main"
	"io/ioutil"
	"os"
	"text/template"
)

const (
	tplPath        = "templates/%s"
	codePath       = "code/%s"
	successMessage = "The code has been generated successfully!"
)

func main() {
	gf, err := os.Create(fmt.Sprintf(codePath, "main.go"))
	die(err)
	defer gf.Close()

	tplContent, err := ioutil.ReadFile(fmt.Sprintf(tplPath, "main/main.tpl"))
	die(err)

	var tpl = template.Must(template.New(fmt.Sprintf(codePath, "main")).Parse(string(tplContent)))

	tpl.Execute(gf, gcg.New("main", "ServiceName"))

	fmt.Println(successMessage)
}
