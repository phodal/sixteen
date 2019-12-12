package domain

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sixteen/utils"
	"strings"
)

const TASK_PATH = "docs/refactoring/"

type TaskModel struct {
	Id    string
	Title string
	Done  bool
	Todos []TodoModel
}

type TodoModel struct {
	Done    bool
	Content string
}

type Task struct {
}

func (t Task) list() {

}

func GetTaskIdFromFilePath(filePath string) string {
	splitPath := strings.Split(filePath, "/")
	taskName := splitPath[len(splitPath)-1]
	id := taskName[0:utils.ID_LENGTH]
	return id
}

func ParseTask(filePath string) (*TaskModel, error) {
	id := GetTaskIdFromFilePath(filePath)
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
			if strings.Contains(match, " - [x]") {
				hasDone = true
			}
			todo := &TodoModel{Content: content, Done: hasDone}
			todos = append(todos, *todo)
		}
	}

	file.Close()

	task := &TaskModel{
		Id:    id,
		Done:  false,
		Title: strings.ReplaceAll(txtlines[0], "# ", ""),
		Todos: todos,
	}

	return task, nil
}

func GetTasks() []TaskModel {
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

func TaskToMap(tasks []TaskModel) map[string]TaskModel {
	var maps = make(map[string]TaskModel)
	for _, task := range tasks {
		maps[task.Id] = task
	}

	return maps
}
