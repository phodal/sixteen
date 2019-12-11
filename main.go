package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/manifoldco/promptui"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"sixteen/utils"
	"strconv"
	"strings"
)

type TaskModel struct {
	Id    string
	Title string
	Todos []string
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
		listSteps()
	default:
		validate()
	}
}

func listSteps() {

}

const task_path = "docs/refactoring/"

func listTasks() []TaskModel {
	files, err := ioutil.ReadDir(task_path)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []TaskModel
	for _, f := range files {
		task, _ := ParseTask(task_path + f.Name())
		tasks = append(tasks, *task)
	}
	return tasks
}

func renderHookDropCodeBlock(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	nodeType := reflect.TypeOf(node).String()
	fmt.Println(nodeType)
	switch nodeType {
	case "*ast.Heading":
		heading := node.(*ast.Heading)
		fmt.Println(heading)
		return ast.GoToNext, true
	case "*ast.ListItem":
		listItem := node.(*ast.ListItem)
		fmt.Println(listItem.Content)
		return ast.GoToNext, true
	default:
		return ast.GoToNext, false
	}
}

func ParseTask(filePath string) (*TaskModel, error) {
	id := getIdFromFileName(filePath)
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, nil
	}

	opts := html.RendererOptions{
		Flags:          html.CommonFlags,
		RenderNodeHook: renderHookDropCodeBlock,
	}

	renderer := html.NewRenderer(opts)
	html := markdown.ToHTML(data, nil, renderer)
	fmt.Println(string(html))

	file.Close()

	task := &TaskModel{
		Id:    id,
		Title: strings.ReplaceAll(txtlines[0], "# ", ""),
		Todos: nil,
	}

	return task, nil
}

func getIdFromFileName(filePath string) string {
	splitPath := strings.Split(filePath, "/")
	taskName := splitPath[len(splitPath)-1]
	id := taskName[0:utils.ID_LENGTH]
	return id
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
