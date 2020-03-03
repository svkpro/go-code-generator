package main

import (
	"fmt"
	"go-code-generator/helpers"
	gcgc "go-code-generator/templates/config"
	gcgd "go-code-generator/templates/domain"
	gcgm "go-code-generator/templates/cmd"
	"os"
)

const (
	successMessage = "The code has been generated successfully!"
)

func main() {
	helpers.GenerateGo("cmd", "main", gcgm.NewGoTpl())
	helpers.GenerateGo("config", "config", gcgc.NewGoTpl())
	helpers.GenerateJson("config", "config", gcgc.NewJsonTpl())
	helpers.GenerateGo("domain", "service", gcgd.NewGoTpl())

	wd, err := os.Getwd()
	helpers.Die(err)
	cwd := fmt.Sprintf("%s/code", wd)

	files, err := helpers.TakeFiles(cwd)
	helpers.Die(err)
	err = helpers.ZipFiles(fmt.Sprintf("%s/code.zip", cwd), files)
	helpers.Die(err)

	fmt.Println(successMessage)
}
