package googlehttpapi

import (
	"fmt"
	"regexp"
)

func Parsing(url string) (string, error) {
	r, err := regexp.Compile("{([A-Za-z0-9_]+)}")
	if err != nil {
		return "", err
	}
	output := r.ReplaceAllStringFunc(url, replaceStringFunc)
	out := fmt.Sprintf("\"%s\"", output)
	return out, nil
}

func replaceStringFunc(variableName string) string {
	return fmt.Sprintf("\"+ arg.%s +\"", variableName[1:len(variableName)-1])
}
