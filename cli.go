package bugs

import (
	"fmt"
	"os"
	"strings"
)

const (
	VerbHelp   = "help"
	VerbList   = "list"
	VerbCreate = "create"
	VerbUpdate = "update"
)

// Action is returned by ParseArgs
// It represents what the cli should do, and
// the variables it needs to do it. Verbs are commands
// in the cli like "create", "list", and "update". Some verbs
// need fewer variables than others, but each verb can depend
// on the variables it needs being populated by ParseArgs
type Action struct {
	Verb      string
	BugTitle  string
	BugID     string
	BugStatus string
}

// ParseArgs takes in all of the cli arguments, determines what
// command or "verb" the user wants to perform, and populates the
// required variables. It returns that as an Action type
// If any of those steps goes wrong, it'll return the help message
func ParseArgs(args []string) Action {
	if len(args) < 1 {
		return Action{
			Verb: VerbHelp,
		}
	}

	switch args[0] {
	case "list", "List":
		return Action{
			Verb: VerbList,
		}
	case "create", "Create":
		if len(args) < 2 {
			return Action{
				Verb: VerbHelp,
			}
		}
		return Action{
			Verb:     VerbCreate,
			BugTitle: strings.Join(args[1:], " "),
		}
	case "update", "Update":
		if len(args) != 3 {
			return Action{
				Verb: VerbHelp,
			}
		}
		return Action{
			Verb:      VerbUpdate,
			BugID:     args[1],
			BugStatus: args[2],
		}
	default:
		return Action{
			Verb: VerbHelp,
		}
	}
}

// Main runs the cli for bugs. It calls the ParseArgs function,
// runs the actions returned by ParseArgs, catches any errors,
// displays messages and errors and exits appropriately
func Main() {
	saveFile := ".buglist.json"
	buglist, err := OpenBugList(saveFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open buglist, %v", err)
		os.Exit(1)
	}

	// First arg is just the binary path. ParseArgs shouldn't care about that
	action := ParseArgs(os.Args[1:])

	switch action.Verb {
	case VerbCreate:
		_, err := buglist.CreateBug(action.BugTitle)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		err = buglist.Write(saveFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case VerbUpdate:
		err := buglist.UpdateBugStatus(action.BugID, action.BugStatus)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		err = buglist.Write(saveFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case VerbHelp:
		help()
	case VerbList:
		list(buglist)
	}

}

// list displays all of the bugs in a given Buglist for the user
// It formats it in an easy to read way
func list(buglist *Buglist) {

	fmt.Println("ID Status Title")
	fmt.Println("=================")

	for _, value := range buglist.All() {
		fmt.Println(value.ID, value.Status, " ", value.Title)
	}

}

// help displays the help message
func help() {
	multilineHelpString := `
	Commands:
	list
	create
	update
	help

	Usage Info:

	List will list the ID, Title, and Status of all your bugs:
	Example:
	$ bugs list

	Create will create an open bug with whatever title you provide
	Examples:
	$ bugs create <title of bug to create>
	$ bugs create new bug!

	Update will update the status of the bug with the ID you provided
	Examples:
	$ bugs update <id of bug to update> <new status>
	$ bugs update 2 Closed

	Help will print this help message
	Examples:
	$ bugs help
	`
	fmt.Println(multilineHelpString)
}
