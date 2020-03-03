package main

import (
	"fmt"
	"go-code-generator/helpers"
	gcgc "go-code-generator/templates/config"
	gcgm "go-code-generator/templates/main"
)

const (
	successMessage = "The code has been generated successfully!"
)

func main() {
	helpers.GenerateGo("cmd", "main", gcgm.NewGoTpl())
	helpers.GenerateGo("config", "config", gcgc.NewGoTpl())
	helpers.GenerateJson("config", "config", gcgc.NewJsonTpl())

	fmt.Println(successMessage)
}
