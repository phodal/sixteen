package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"sixteen/domain"
	. "sixteen/domain"
	"sixteen/utils"
)

func main() {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			var command = c.Args().Get(0)
			if command == "" {
				res, done := runPrompt()
				if done {
					return nil
				}
				command = res
			}

			executeCommand(command)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func runPrompt() (string, bool) {
	prompt := promptui.Select{
		Label: "Refactoring",
		Items: []string{
			"list",
			"step",
			"switch",
			"delete",
			"commit",
			"create",
		},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", true
	}
	return result, false
}

func executeCommand(result string) {
	switch result {
	case "list":
		tasks := domain.GetTasks()
		fmt.Println(tasks)
	case "create":
		createNew()
	case "step":
		tasks := domain.GetTasks()
		index := selectTask(tasks)
		name := getStepName(tasks[index])
		fmt.Println(name)
	case "commit":
		doCommit()
	default:
		fmt.Printf("command %s not found----", result)
	}
}

func doCommit() {
	prompt := promptui.Prompt{
		Label: "Commit Message",
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	tasks := domain.GetTasks()
	task := tasks[selectTask(tasks)]

	utils.CommitByMessage("refactoring: " + result + "-" + task.Id)
}

func selectTask(tasks []TaskModel) int {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Id | cyan }}-{{ .Title | red }} {{ if eq .Done false }} ⌛ {{end}}",
		Inactive: "  {{ .Id | cyan }}-{{ .Title | red }} {{ if eq .Done false }} ⌛ {{end}}",
		Selected: "\U0001F336 {{ .Id | red | cyan }}-{{ if eq .Done false }} ⌛ {{end}}",
	}

	prompt := promptui.Select{
		Label:     "Refactoring",
		Templates: templates,
		Items:     tasks,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return 0
	}

	return i
}

func getStepName(model TaskModel) string {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Content | red }} {{ if eq .Done true }} 👍 {{end}}",
		Inactive: "  {{ .Content | red }} {{ if eq .Done true }} 👍 {{end}}",
		Selected: "\U0001F336 {{ if eq .Done true }} 👍 {{end}}",
	}

	prompt := promptui.Select{
		Label:     model.Title,
		Templates: templates,
		Items:     model.Todos,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
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
	_ = os.MkdirAll(TASK_PATH, os.ModePerm)

	fileName := utils.BuildFileName(utils.GenerateId(), title)
	_ = ioutil.WriteFile(TASK_PATH+"/"+fileName, []byte("# "+title+"\n\n"+" - [ ] todo"), 0644)
}
