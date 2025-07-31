package bugs_test

import (
	"slices"
	"testing"

	"github.com/unkownblunders/bugs"
)

// The get function should return all the bugs from our persistent storage
func TestGetBugs(t *testing.T) {
	bugList, err := bugs.Get("testdata/test-bugs.txt")

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	assertTestBugList(t, bugList)

}

func TestWriteBugs(t *testing.T) {
	// Get the testdata buglist, write it to a new location
	// Get the buglist from the new location
	// assert that it should be an identical copy

	buglist, err := bugs.Get("testdata/test-bugs.txt")

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

	copyBugList, err := bugs.Get(newfile)

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	assertTestBugList(t, copyBugList)

}

func TestGetBugFindsBugByTitle(t *testing.T) {
	buglist, _ := bugs.Get("testdata/test-bugs.txt")

	want := bugs.Bug{
		Title:  "Adding bugs is broken",
		Status: "Closed",
	}
	bug, err := buglist.GetBug("Adding bugs is broken")

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
	buglist, _ := bugs.Get("testdata/test-bugs.txt")

	_, err := buglist.GetBug("nonexistant bug")

	if err == nil {
		t.Error("Want error for nonexistant bug, got nil")
	}
}

func TestCreateBugs(t *testing.T) {
	// If the file is empty:
	var newList bugs.Buglist
	newBug := bugs.Bug{Title: "Creating Bugs is broken", Status: "Closed"}

	newList.Create(newBug.Title, newBug.Status)

	if !slices.Equal(newList.All(), []bugs.Bug{newBug}) {
		t.Errorf("Found mismatch: got %#v, expected %#v", newList.All(), []bugs.Bug{newBug})
	}

}

func TestUpdateStatusBugs(t *testing.T) {

	buglist, _ := bugs.Get("testdata/test-bugs.txt")

	bugUpdate := bugs.Bug{"new bug!", "Closed"}

	err := buglist.UpdateStatus(bugUpdate.Title, bugUpdate.Status)

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	updatedBug, _ := buglist.GetBug("new bug!")

	if updatedBug.Status != bugUpdate.Status {
		t.Errorf("Status update failed. got %#v, expected %#v", updatedBug.Status, bugUpdate.Status)
	}
}

// Doesn't start with "Test", therefore not a test
func assertTestBugList(t *testing.T, buglist *bugs.Buglist) {
	want := []bugs.Bug{
		{
			Title:  "Adding bugs is broken",
			Status: "Closed",
		},
		{
			Title:  "new bug!",
			Status: "Open",
		},
		{
			Title:  "Another bug :(",
			Status: "Open",
		},
	}

	if !slices.Equal(buglist.All(), want) {
		t.Errorf("Found mismatch: got %#v, expected %#v", buglist.All(), want)
	}
}
