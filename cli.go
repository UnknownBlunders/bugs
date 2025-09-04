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

type Action struct {
	Verb      string
	BugTitle  string
	BugID     string
	BugStatus string
}

func ParseArgs(args []string) Action {
	if len(args) < 2 {
		return Action{
			Verb: VerbHelp,
		}
	}

	switch args[1] {
	case "list", "List":
		return Action{
			Verb: VerbList,
		}
	case "create", "Create":
		if len(args) < 3 {
			return Action{
				Verb: VerbHelp,
			}
		}
		return Action{
			Verb:     VerbCreate,
			BugTitle: strings.Join(args[2:], " "),
		}
	case "update", "Update":
		if len(args) != 4 {
			return Action{
				Verb: VerbHelp,
			}
		}
		return Action{
			Verb:      VerbUpdate,
			BugID:     args[2],
			BugStatus: args[3],
		}
	default:
		return Action{
			Verb: VerbHelp,
		}
	}
}

func Main() {
	saveFile := ".buglist.json"
	buglist, err := OpenBugList(saveFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open buglist, %v", err)
		os.Exit(1)
	}

	action := ParseArgs(os.Args)

	switch action.Verb {
	case VerbCreate:
		buglist.CreateBug(action.BugTitle)

		err := buglist.Write(saveFile)
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

func list(buglist *Buglist) {

	fmt.Println("ID Status Title")
	fmt.Println("=================")

	for _, value := range buglist.All() {
		fmt.Println(value.ID, value.Status, " ", value.Title)
	}

}

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
