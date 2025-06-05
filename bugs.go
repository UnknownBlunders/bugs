package bugs

import (
	"errors"
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

	// Returns the error only if one exists
	// and the error isn't just that the file doesn't exist
	if err != nil && !errors.Is(err, os.ErrNotExist) {
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
