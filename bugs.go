package bugs

import (
	"os"
	"strings"
)

func Get(filename string) ([]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var bugString = string(file)
	bugs := strings.Split(bugString, "\n")

	return bugs, nil
	// return []string{"1", "2", "3"}, nil
}

func Create(filename string, title string) error {

	buglist, err := Get(filename)

	if err != nil {
		return err
	}

	buglist = append(buglist, title)

	// our get function returns a slice of strings
	// Need to convert to one long string, then convert to byte slice
	bugString := strings.Join(buglist, "\n")
	var bugBytes = []byte(bugString)

	err = os.WriteFile(filename, bugBytes, 0777)
	if err != nil {
		return err
	}

	return nil
}
