package bugs_test

import (
	"testing"

	"github.com/unkownblunders/bugs"
)

// The get function should return all the bugs from our persistent storage
func TestGetBugs(t *testing.T) {
	var bugList []bugs.Bug
	bugList, err := bugs.Get("testdata/test-bugs.txt")

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

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	assertBugList(t, bugList, want)

}

func TestCreateBugs(t *testing.T) {
	newBug := bugs.Bug{Title: "Creating Bugs is broken", Status: "Closed"}

	filename := t.TempDir() + "/buglist.txt"

	err := bugs.Create(filename, newBug.Title, newBug.Status)

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	bugList, err := bugs.Get(filename)

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	assertBugList(t, bugList, []bugs.Bug{newBug})
}

// Doesn't start with "Test", therefore not a test
func assertBugList(t *testing.T, bugList []bugs.Bug, expectedBugs []bugs.Bug) {

	for index, bug := range bugList {
		if bug != expectedBugs[index] {
			// continues to test, even if it errors
			t.Errorf("Found mismatch: got %q, expected %q", bug, expectedBugs[index])
		}
	}
}
