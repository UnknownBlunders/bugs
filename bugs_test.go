package bugs_test

import (
	"slices"
	"testing"

	"github.com/unknownblunders/bugs"
)

// The get function should return all the bugs from our persistent storage
func TestOpenBugListReadsBuglistFromExistingFile(t *testing.T) {
	bugList, err := bugs.OpenBugList("testdata/test-bugs.txt")

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	assertTestBugList(t, bugList)

}

func TestOpenBugListCreatesNewEmptyBuglistWhenFileIsMissing(t *testing.T) {
	saveFilePath := t.TempDir() + "/buglist.json"

	buglist, err := bugs.OpenBugList(saveFilePath)
	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	// Check that newly created buglist is empty
	if !slices.Equal(buglist.All(), []bugs.Bug{}) {
		// ends the test
		t.Fatal(err)
	}

}

func TestWriteBugs(t *testing.T) {
	// Get the testdata buglist, write it to a new location
	// Get the buglist from the new location
	// assert that it should be an identical copy

	buglist, err := bugs.OpenBugList("testdata/test-bugs.txt")

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	newfile := t.TempDir() + "/buglist.txt"

	err = buglist.Write(newfile)

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	copyBugList, err := bugs.OpenBugList(newfile)

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	assertTestBugList(t, copyBugList)

}

func TestGetBugFindsBugByTitle(t *testing.T) {
	buglist, _ := bugs.OpenBugList("testdata/test-bugs.txt")

	want := bugs.Bug{
		ID:     "0",
		Title:  "Adding bugs is broken",
		Status: bugs.StatusClosed,
	}
	bug, err := buglist.GetBug("0")

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	if bug != want {
		// %q quoted values like: {"Adding bugs is broken" "Closed"}
		// %v unquoted values like: {Adding bugs is broken Closed}
		// %#v alternate format. For structs, print as Go values like: bugs.Bug{Title:"Adding bugs is broken", Status:"Closed"}
		t.Errorf("Found mismatch: got %#v, expected %#v", bug, want)
	}

}

func TestGetBugErrorsIfBugNotFound(t *testing.T) {
	buglist, _ := bugs.OpenBugList("testdata/test-bugs.txt")

	_, err := buglist.GetBug("nonexistant id")

	if err == nil {
		t.Error("Want error for nonexistant bug, got nil")
	}
}

func TestCreateBugs(t *testing.T) {
	var newList bugs.Buglist
	id := newList.CreateBug("BugA")
	bugA, err := newList.GetBug(id)

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	if bugA.Title != "BugA" {
		t.Errorf("Found mismatch: wrong title")
	}

	id = newList.CreateBug("BugB")
	bugB, err := newList.GetBug(id)

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	if bugB.Title != "BugB" {
		t.Errorf("Found mismatch: wrong title")
	}

}

func TestUpdateStatusBugs(t *testing.T) {
	var newList bugs.Buglist

	id := newList.CreateBug("test bug")

	err := newList.UpdateBugStatus(id, bugs.StatusClosed)

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	updatedBug, err := newList.GetBug(id)

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	if updatedBug.Status != bugs.StatusClosed {
		t.Errorf("Bug status was wrong: %q", updatedBug.Status)
	}
}

// Doesn't start with "Test", therefore not a test
func assertTestBugList(t *testing.T, buglist *bugs.Buglist) {
	want := []bugs.Bug{
		{
			ID:     "0",
			Title:  "Adding bugs is broken",
			Status: bugs.StatusClosed,
		},
		{
			ID:     "1",
			Title:  "new bug!",
			Status: bugs.StatusOpen,
		},
		{
			ID:     "2",
			Title:  "Another bug :(",
			Status: bugs.StatusOpen,
		},
	}

	if !slices.Equal(buglist.All(), want) {
		t.Errorf("Found mismatch: got %#v, expected %#v", buglist.All(), want)
	}
}
