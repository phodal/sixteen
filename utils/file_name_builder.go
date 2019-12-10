package utils

import "regexp"

func BuildFileName(id string, content string) string {
	return id + "-" + UpdateFileName(content) + ".feature"
}

func UpdateFileName(name string) string {
	rex := regexp.MustCompile(`[ ，,。！？]`)
	return rex.ReplaceAllString(name, "-")
}
