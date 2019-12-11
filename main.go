package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sixteen/domain"
	"sixteen/utils"
	"strconv"
	"strings"
)

type TaskModel struct {
	Id    string
	Title string
	Todos []TodoModel
}

type TodoModel struct {
	Done    bool
	Content string
}

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

	switch result {
	case "list":
		tasks := listTasks()
		fmt.Println(tasks)
	case "create":
		createNew()
	case "step":
		tasks := listTasks()
		listSteps(tasks)
	default:
		validate()
	}
}

func listSteps([]TaskModel) {

}

const TASK_PATH = "docs/refactoring/"

func listTasks() []TaskModel {
	files, err := ioutil.ReadDir(TASK_PATH)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []TaskModel
	for _, f := range files {
		task, _ := ParseTask(TASK_PATH + f.Name())
		tasks = append(tasks, *task)
	}
	return tasks
}

func ParseTask(filePath string) (*TaskModel, error) {
	id := domain.GetTaskIdFromFilePath(filePath)
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	var todos []TodoModel
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
		var todoCompile = regexp.MustCompile(`\s-\s\[[ |x]\]\s(.*)`)

		for _, match := range todoCompile.FindAllString(scanner.Text(), -1) {
			content := todoCompile.ReplaceAllString(match, `$1`)
			var hasDone = false
			if (strings.Contains(match, " - [x]")) {
				hasDone =true
			}
			todo := &TodoModel{Content: content, Done: hasDone}
			todos = append(todos, *todo)
		}
	}

	file.Close()

	task := &TaskModel{
		Id:    id,
		Title: strings.ReplaceAll(txtlines[0], "# ", ""),
		Todos: todos,
	}

	return task, nil
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
