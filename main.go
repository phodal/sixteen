package main

import (
	"errors"
	//"sixteen/cmd"
	"fmt"
	"github.com/manifoldco/promptui"
	"strconv"
)

func main() {
	prompt := promptui.Select{
		Label: "Refactoring",
		Items: []string{
			"list",
			"step",
			"switch",
			"delete",
			"create",
		},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Println(result)
	switch result {
	case "list":
		fmt.Println("list")
	default:
		//listLog
		validate()
	}
}

func validate() {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Number",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}
