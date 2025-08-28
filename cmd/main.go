package main

import (
	"fmt"
	"os"

	"github.com/unkownblunders/bugs"
)

func list(buglist *bugs.Buglist) {

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

func main() {
	saveFile := ".buglist.json"
	buglist, err := bugs.OpenBugList(saveFile)
	if err != nil {
		fmt.Println("Failed to open buglist", err)
	}

	action := bugs.ParseArgs(os.Args)

	switch action.Verb {
	case bugs.VerbCreate:
		buglist.CreateBug(action.BugTitle)
		buglist.Write(saveFile)
	case "update":
		err := buglist.UpdateBugStatus(action.BugID, action.BugStatus)
		if err != nil {
			fmt.Println("no such bug", err)
		} else {
			buglist.Write(saveFile)
		}
	case "help":
		help()
	case "list":
		list(buglist)
	}

}
