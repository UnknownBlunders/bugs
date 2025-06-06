package bugs

import (
	"encoding/json"
	"errors"
	"os"
)

type Bug struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

func Get(filename string) ([]Bug, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var bugString = string(file)
	var bugs []Bug

	err = json.Unmarshal([]byte(bugString), &bugs)
	if err != nil {
		return nil, err
	}

	return bugs, nil
	// return []string{"1", "2", "3"}, nil
}

func Create(filename string, title string, status string) error {

	buglist, err := Get(filename)

	// Returns the error only if one exists
	// and the error isn't just that the file doesn't exist
	// If the file doesn't exist, we'll just create it later
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	newBug := Bug{Title: title, Status: status}

	buglist = append(buglist, newBug)

	data, err := json.Marshal(buglist)

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0777)
	if err != nil {
		return err
	}

	return nil
}
