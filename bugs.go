package bugs

import (
	"os"
	"strings"
)

func Get() ([]string, error) {
	file, err := os.ReadFile("bugs.txt")
	if err != nil {
		return nil, err
	}

	var bugString = string(file)
	bugs := strings.Split(bugString, "\n")

	return bugs, nil
}

func Create(title string) {
	//bugs = append(bugs, title)
}
