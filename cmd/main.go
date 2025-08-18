package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/patyumi/script-hub/internal"
)

func main() {
	fmt.Println("Hello from main!")

	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"Gerador de dados", "Sair"},
	}

	_, result, err := prompt.Run()

	switch result {

	case "Gerador de dados":
		internal.GeradorDados()

	case "Sair":
		fmt.Printf("Até mais!\n")

	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

}
