package bugs

import (
	"encoding/json"
	"errors"
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
	bugs   []Bug
	nextID int
}

type saveBugList struct {
	Bugs   []Bug
	NextID int
}

func NewBuglist() *Buglist {
	return &Buglist{
		bugs:   []Bug{},
		nextID: 0,
	}
}

func OpenBugList(filename string) (*Buglist, error) {
	data, err := os.ReadFile(filename)
	if errors.Is(err, os.ErrNotExist) {
		return NewBuglist(), nil
	}
	if err != nil {
		return nil, err
	}

	var savedBugList saveBugList
	err = json.Unmarshal(data, &savedBugList)
	if err != nil {
		return nil, err
	}

	return &Buglist{
			bugs:   savedBugList.Bugs,
			nextID: savedBugList.NextID,
		},
		nil
}

func (buglist *Buglist) Write(filename string) error {

	data, err := json.Marshal(saveBugList{
		Bugs:   buglist.bugs,
		NextID: buglist.nextID,
	})

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0777)
	if err != nil {
		return err
	}

	return nil
}

func (buglist *Buglist) GetBug(id string) (Bug, error) {

	for _, bug := range buglist.bugs {
		if bug.ID == id {
			return bug, nil
		}
	}

	return Bug{}, errors.New("bug not found")
}

func (buglist *Buglist) All() []Bug {
	return buglist.bugs
}

func (buglist *Buglist) CreateBug(title string) (id string) {
	id = strconv.Itoa(buglist.nextID)

	buglist.nextID++
	newBug := Bug{ID: id, Title: title, Status: "Open"}

	buglist.bugs = append(buglist.bugs, newBug)

	return id
}

func (buglist *Buglist) UpdateBugStatus(id string, status string) error {

	// Get and update bug
	for index, bug := range buglist.bugs {
		if bug.ID == id {
			buglist.bugs[index].Status = status
			break
		}
	}

	return nil
}
