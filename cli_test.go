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
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestParseArgsReturnsHelpActionWhenNoArgs(t *testing.T) {
	action := bugs.ParseArgs([]string{"dummy command"})

	if action.Verb != bugs.VerbHelp {
		t.Errorf("Wrong Action, Expected help got %q", action.Verb)
	}
}

func TestParseArgsReturnsHelpActionWhenUnknownCommand(t *testing.T) {
	action := bugs.ParseArgs([]string{"dummy command", "junk command"})

	if action.Verb != bugs.VerbHelp {
		t.Errorf("Wrong Action, Expected help got %q", action.Verb)
	}
}

func TestParseArgsReturnsListActionWhenFirstArgIsList(t *testing.T) {
	action := bugs.ParseArgs([]string{"dummy command", "list"})

	if action.Verb != bugs.VerbList {
		t.Errorf("Wrong Action, Expected list got %q", action.Verb)
	}
}

func TestParseArgsReturnsHelpWhenFirstArgIsCreateWithNoBugTitle(t *testing.T) {
	action := bugs.ParseArgs([]string{"dummy command", "create"})

	if action.Verb != bugs.VerbHelp {
		t.Errorf("Wrong Action, Expected help got %q", action.Verb)
	}
}

func TestParseArgsReturnsCreateActionAndBugTitleWhenFirstArgIsCreateWithTitlesFollowing(t *testing.T) {
	action := bugs.ParseArgs([]string{"dummy command", "create", "new", "bug"})

	if action.Verb != bugs.VerbCreate {
		t.Errorf("Wrong Action, Expected create got %q", action.Verb)
	}

	if action.BugTitle != "new bug" {
		t.Errorf("Wrong bug title, expected new bug got %q", action.BugTitle)
	}
}

func TestParseArgsReturnsHelpWhenFirstArgIsUpdateWithNoBugIdOrStatus(t *testing.T) {
	action := bugs.ParseArgs([]string{"dummy command", "update"})

	if action.Verb != bugs.VerbHelp {
		t.Errorf("Wrong Action, Expected help got %q", action.Verb)
	}
}

func TestParseArgsReturnsHelpWhenFirstArgIsUpdateAndThereAreAdditionalArgsAfterOurIdAndStatus(t *testing.T) {
	action := bugs.ParseArgs([]string{"dummy command", "update", "1", "Closed", "extra-arg"})

	if action.Verb != bugs.VerbHelp {
		t.Errorf("Wrong Action, Expected help got %q", action.Verb)
	}
}

func TestParseArgsReturnsUpdateWhenFirstArgIsUpdateAndThereIsAnIdAndAStatus(t *testing.T) {
	action := bugs.ParseArgs([]string{"dummy command", "update", "1", "Closed"})

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
