package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"sixteen/utils"

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
		listTasks()
	case "create":
		createNew()
	default:
		validate()
	}
}

const task_path = "docs/refactoring/"

func listTasks() []string {
	files, err := ioutil.ReadDir(task_path)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []string
	for _, f := range files {
		task, _ := ParseTask(task_path + f.Name())
		tasks = append(tasks, task)
	}
	return tasks
}

func ParseTask(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		return "", nil
	}

	fmt.Println(data[0])

	return "", nil
}

func createNew() {
	prompt := promptui.Prompt{
		Label: "title",
	}

	title, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	buildRefactoringFile(title)
}

func buildRefactoringFile(title string) {
	_ = os.MkdirAll("docs", os.ModePerm)
	_ = os.MkdirAll(task_path, os.ModePerm)

	fileName := utils.BuildFileName(utils.GenerateId(), title)
	_ = ioutil.WriteFile(task_path+"/"+fileName, []byte("# "+title+"\n\n"+" - [ ] todo"), 0644)
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
