package bugs

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	StatusClosed = "Closed"
	StatusOpen   = "Open"
)

type Bug struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type Buglist struct {
	Bugs   []Bug
	NextID uint64
}

func OpenBugList(path string) (*Buglist, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return &Buglist{}, nil
	}
	if err != nil {
		return nil, err
	}

	var buglist Buglist
	err = json.Unmarshal(data, &buglist)
	if err != nil {
		return nil, err
	}

	return &buglist, nil
}

func (buglist *Buglist) Write(path string) error {
	data, err := json.Marshal(buglist)
	if err != nil {
		return err
	}
	data = append(data, byte('\n'))
	err = os.WriteFile(path, data, 0o777)
	if err != nil {
		return err
	}
	return nil
}

func (buglist *Buglist) GetBug(id string) (Bug, error) {

	for _, bug := range buglist.Bugs {
		if bug.ID == id {
			return bug, nil
		}
	}

	return Bug{}, fmt.Errorf("bug not found %q", id)
}

func (buglist *Buglist) All() []Bug {
	return buglist.Bugs
}

func (buglist *Buglist) CreateBug(title string) (id string, err error) {
	id = strconv.FormatUint(buglist.NextID, 10)

	buglist.NextID++
	if buglist.NextID == 0 {
		return "", errors.New("NextID must have overflowed, it's now 0")
	}
	newBug := Bug{ID: id, Title: title, Status: "Open"}

	buglist.Bugs = append(buglist.Bugs, newBug)

	return id, nil
}

func (buglist *Buglist) UpdateBugStatus(id string, status string) error {

	// Get and update bug
	for index, bug := range buglist.Bugs {
		if bug.ID == id {
			buglist.Bugs[index].Status = status
			return nil
		}
	}

	return fmt.Errorf("Bug not found: %q", id)
}
