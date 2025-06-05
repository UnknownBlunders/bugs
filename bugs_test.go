package bugs_test

import (
	"testing"

	"github.com/unkownblunders/bugs"
)

// The get function should return all the bugs from our persistent storage
func TestGetBugs(t *testing.T) {
	var bugList []string
	bugList, err := bugs.Get("testdata/test-bugs.txt")

	want := []string{"Adding bugs is broken", "new bug!", "Another bug :("}

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	assertBugList(t, bugList, want)

}

func TestCreateBugs(t *testing.T) {
	newBug := "Create bug test"

	filename := t.TempDir() + "/buglist.txt"

	err := bugs.Create(filename, newBug)

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	bugList, err := bugs.Get(filename)

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	assertBugList(t, bugList, []string{newBug})
}

// Doesn't start with "Test", therefore not a test
func assertBugList(t *testing.T, bugList []string, expectedBugs []string) {

	for index, bug := range bugList {
		if bug != expectedBugs[index] {
			// continues to test, even if it errors
			t.Errorf("Found mismatch: got %q, expected %q", bug, expectedBugs[index])
		}
	}
}
