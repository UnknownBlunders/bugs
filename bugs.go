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

type Buglist struct {
	bugs []Bug
}

func Get(filename string) (*Buglist, error) {
	var buglist Buglist

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var bugString = string(file)

	err = json.Unmarshal([]byte(bugString), &buglist.bugs)
	if err != nil {
		return nil, err
	}

	return &buglist, nil
}

func (buglist *Buglist) Write(filename string) error {

	data, err := json.Marshal(buglist.bugs)

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0777)
	if err != nil {
		return err
	}

	return nil
}

func (buglist *Buglist) GetBug(title string) (Bug, error) {

	for _, bug := range buglist.bugs {
		if bug.Title == title {
			return bug, nil
		}
	}

	return Bug{}, errors.New("bug not found")
}

func (buglist *Buglist) All() []Bug {
	return buglist.bugs
}

func (buglist *Buglist) Create(title string, status string) {

	newBug := Bug{Title: title, Status: status}

	buglist.bugs = append(buglist.bugs, newBug)
}

func (buglist *Buglist) UpdateStatus(title string, status string) error {

	// Get and update bug
	for index, bug := range buglist.bugs {
		if bug.Title == title {
			buglist.bugs[index].Status = status
			break
		}
	}

	return nil
}
