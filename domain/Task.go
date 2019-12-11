package domain

import (
	"sixteen/utils"
	"strings"
)

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

