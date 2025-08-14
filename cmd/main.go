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

func readFile(saveFile string) (buglist *bugs.Buglist) {
	buglist, err := bugs.Get(saveFile)

	if err != nil {
		fmt.Println("Failed to get bugs", err)
		os.Exit(1)
	}

	return buglist
}

func help() {
	fmt.Println("Commands:")
	fmt.Println("init")
	fmt.Println("create")
	fmt.Println("list")
	fmt.Println("update")
	fmt.Println("help")
}

func main() {
	saveFile := ".buglist.json"

	if len(os.Args) >= 2 {

		switch os.Args[1] {
		case "create":
			for _, title := range os.Args[2:] {
				buglist := readFile(saveFile)
				buglist.CreateBug(title)
				buglist.Write(saveFile)
			}
		case "update":
			buglist := readFile(saveFile)
			err := buglist.UpdateBugStatus(os.Args[2], os.Args[3])
			if err != nil {
				fmt.Println("no such bug", err)
			} else {
				buglist.Write(saveFile)
			}
		case "help":
			help()
		case "list":
			buglist := readFile(saveFile)
			list(buglist)
		case "init":
			err := bugs.InitializeSaveFile(saveFile)
			if err != nil {
				fmt.Println("Error creating save file", err)
			}
		default:
			fmt.Println("Unknown command: ", os.Args[1])
			fmt.Println("")
			help()
		}

	}

}
