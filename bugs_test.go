package bugs_test

import (
	"math"
	"slices"
	"testing"

	"github.com/unknownblunders/bugs"
)

func TestOpenBugList_ReadsBuglistFromExistingFile(t *testing.T) {
	t.Parallel()
	bugList, err := bugs.OpenBugList("testdata/test-bugs.txt")

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	assertTestBugList(t, bugList)

}

func TestOpenBugList_CreatesNewEmptyBuglistWhenBuglistFileDoesntExistYet(t *testing.T) {
	t.Parallel()
	saveFilePath := t.TempDir() + "/buglist.json"

	buglist, err := bugs.OpenBugList(saveFilePath)
	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	// Check that newly created buglist is empty
	if len(buglist.All()) > 0 {
		// ends the test
		t.Fatal(err)
	}

}

func TestWrite_WritesGivenBuglistToAFile(t *testing.T) {
	// Get the testdata buglist, write it to a new location
	// Get the buglist from the new location
	// assert that it should be an identical copy

	t.Parallel()
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

func TestGetBug_FindsBugByTitle(t *testing.T) {
	t.Parallel()
	buglist, err := bugs.OpenBugList("testdata/test-bugs.txt")
	if err != nil {
		// ends the test
		t.Fatal(err)
	}

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

func TestGetBug_ErrorsIfBugNotFound(t *testing.T) {
	t.Parallel()
	buglist, err := bugs.OpenBugList("testdata/test-bugs.txt")
	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	_, err = buglist.GetBug("nonexistent id")

	if err == nil {
		t.Error("Want error for nonexistent bug, got nil")
	}
}

func TestCreateBug_AddsGivenBugToGivenBuglist(t *testing.T) {
	t.Parallel()
	var newList bugs.Buglist
	id, err := newList.CreateBug("BugA")
	if err != nil {
		// ends the test
		t.Fatal(err)
	}
	bugA, err := newList.GetBug(id)

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	if bugA.Title != "BugA" {
		t.Errorf("Found mismatch: wrong title")
	}

	id, err = newList.CreateBug("BugB")
	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	bugB, err := newList.GetBug(id)

	if err != nil {
		// ends the test
		t.Fatal(err)
	}

	if bugB.Title != "BugB" {
		t.Errorf("Found mismatch: wrong title")
	}

}

func TestUpdateBugStatus_UpdatesStatusOfGivenBugByIDToGivenBugStatusString(t *testing.T) {
	t.Parallel()
	var newList bugs.Buglist
	id, err := newList.CreateBug("test bug")
	if err != nil {
		// ends the test
		t.Fatal(err)
	}
	err = newList.UpdateBugStatus(id, bugs.StatusClosed)
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

func TestCreateBug_ErrorsIfNextIDWouldOverflowMaxIntIfIncremented(t *testing.T) {
	t.Parallel()
	buglist := bugs.Buglist{
		NextID: math.MaxUint64,
	}
	t.Logf("NextID: %d", buglist.NextID)
	_, err := buglist.CreateBug("test bug")
	t.Logf("NextID: %d", buglist.NextID)

	if err == nil {
		t.Error("Want error for overflowing nextID, got nil")
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
