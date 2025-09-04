package bugs_test

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
	"github.com/unknownblunders/bugs"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"bugs": bugs.Main,
	})
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestParseArgs_ReturnsHelpActionWhenNoArgs(t *testing.T) {
	t.Parallel()
	action := bugs.ParseArgs([]string{})

	if action.Verb != bugs.VerbHelp {
		t.Errorf("Wrong Action, Expected help got %q", action.Verb)
	}
}

func TestParseArgs_ReturnsHelpActionWhenUnknownCommand(t *testing.T) {
	t.Parallel()
	action := bugs.ParseArgs([]string{"junk command"})

	if action.Verb != bugs.VerbHelp {
		t.Errorf("Wrong Action, Expected help got %q", action.Verb)
	}
}

func TestParseArgs_ReturnsListActionWhenFirstArgIsList(t *testing.T) {
	t.Parallel()
	action := bugs.ParseArgs([]string{"list"})

	if action.Verb != bugs.VerbList {
		t.Errorf("Wrong Action, Expected list got %q", action.Verb)
	}
}

func TestParseArgs_ReturnsHelpWhenFirstArgIsCreateWithNoBugTitle(t *testing.T) {
	t.Parallel()
	action := bugs.ParseArgs([]string{"create"})

	if action.Verb != bugs.VerbHelp {
		t.Errorf("Wrong Action, Expected help got %q", action.Verb)
	}
}

func TestParseArgs_ReturnsCreateActionAndBugTitleWhenFirstArgIsCreateWithTitlesFollowing(t *testing.T) {
	t.Parallel()
	action := bugs.ParseArgs([]string{"create", "new", "bug"})

	if action.Verb != bugs.VerbCreate {
		t.Errorf("Wrong Action, Expected create got %q", action.Verb)
	}

	if action.BugTitle != "new bug" {
		t.Errorf("Wrong bug title, expected new bug got %q", action.BugTitle)
	}
}

func TestParseArgs_ReturnsHelpWhenFirstArgIsUpdateWithNoBugIdOrStatus(t *testing.T) {
	t.Parallel()
	action := bugs.ParseArgs([]string{"update"})

	if action.Verb != bugs.VerbHelp {
		t.Errorf("Wrong Action, Expected help got %q", action.Verb)
	}
}

func TestParseArgs_ReturnsHelpWhenFirstArgIsUpdateAndThereAreAdditionalArgsAfterOurIdAndStatus(t *testing.T) {
	t.Parallel()
	action := bugs.ParseArgs([]string{"update", "1", "Closed", "extra-arg"})

	if action.Verb != bugs.VerbHelp {
		t.Errorf("Wrong Action, Expected help got %q", action.Verb)
	}
}

func TestParseArgs_ReturnsUpdateWhenFirstArgIsUpdateAndThereIsAnIdAndAStatus(t *testing.T) {
	t.Parallel()
	action := bugs.ParseArgs([]string{"update", "1", "Closed"})

	if action.Verb != bugs.VerbUpdate {
		t.Errorf("Wrong Action, Expected update got %q", action.Verb)
	}

	if action.BugID != "1" {
		t.Errorf("Wrong bug ID, expected 1 got %q", action.BugID)
	}

	if action.BugStatus != "Closed" {
		t.Errorf("Wrong bug Status, expected Closed got %q", action.BugStatus)
	}
}
