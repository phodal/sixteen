package utils

import "github.com/teris-io/shortid"

const ID_LENGTH = 9

func GenerateId() string {
	str, _ := shortid.Generate()
	return str
}