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

// Bug represents a individual bug entry
// It has an ID, which must be unique within a Buglist,
// a title which briefly describes a problem
// and a string status, usually "Open" or "Closed".
// While ID is a string, as it really could be anything,
// right now it will always be a number.
type Bug struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// Buglist contains a list of the above bugs called "Bugs". It also includes
// an unsigned 64 bit integer. This integer is used to track
// which id to assign to the next new bug. It get incremented as
// new bugs are added to the Buglist
type Buglist struct {
	Bugs   []Bug
	NextID uint64
}

// OpenBugList opens the file at the given path, unmarshal the json data and
// returns the Buglist object. If it can't open the file because the file doesn't
// exist, it'll create an empty Buglist and return it. If it can't open the file
// or unmarshal the data for any other reason, it'll return an error
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

// Write is given a Buglist and a file path, marshals the Buglist to json, and
// writes it to the file path. It includes a newline at the end of the file for
// posix compatibility.
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

// GetBug, when given a Buglist and a bug id, will return the bug with that id
// in the Buglist. If a bug with that id doesn't exist in the given Buglist,
// GetBug returns an error saying "bug not found"
func (buglist *Buglist) GetBug(id string) (Bug, error) {

	for _, bug := range buglist.Bugs {
		if bug.ID == id {
			return bug, nil
		}
	}

	return Bug{}, fmt.Errorf("bug not found %q", id)
}

// All returns the entire given buglist as a slice of bugs
func (buglist *Buglist) All() []Bug {
	return buglist.Bugs
}

// CreateBug takes in the title of a new bug, assigns this new bug an id 
// from the given Buglist's NexID, and assigns it a status of "Open".
// It adds this bug to the given Buglist, and increments the NextID
// It checks that the NextID integer hasn't overflowed. It will return an error 
// if it has
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

// UpdateBugStatus is given a bug ID and a new status. It finds that bug by ID
// in the given Buglist and changes the bug's status to the given new status.
// If it can't find the bug with the given ID, it'll return an error.
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
