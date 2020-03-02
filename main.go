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
	helpers.Generate("cmd", "main", gcgm.New())
	helpers.Generate("config", "config", gcgc.New())

	fmt.Println(successMessage)
}
