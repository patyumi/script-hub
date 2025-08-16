package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func main() {
	fmt.Println("Hello from main!")

	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"Entrar", "Sair"},
	}

	_, result, err := prompt.Run()

	switch result {

	case "Sair":
		fmt.Printf("You choose %q\n", result)

	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

}
