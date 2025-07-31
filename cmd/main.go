package main

import (
	"os"

	"github.com/unkownblunders/bugs"
)

func list(buglist *bugs.Buglist) {

	println("# Status Title")
	println("=================")

	for index, value := range buglist.All() {
		println(index, value.Status, " ", value.Title)
	}

}

func help() {
	println("Commands:")
	println("create")
	println("list")
	println("update")
	println("help")
}

func main() {
	saveFile := "bugs.txt"
	buglist, err := bugs.Get(saveFile)

	if err != nil {
		println("Failed to get bugs", err)
		return
	}

	if len(os.Args) >= 2 {

		switch os.Args[1] {
		case "create":
			for _, title := range os.Args[2:] {
				buglist.Create(title, "Open")
				buglist.Write(saveFile)
			}
		case "update":
			err := buglist.UpdateStatus(os.Args[2], os.Args[3])
			if err != nil {
				println("no such bug", err)
			} else {
				buglist.Write(saveFile)
			}
		case "help":
			help()
		case "list":
			list(buglist)
		default:
			println("Unknown command: ", os.Args[1])
			println("")
			help()
		}

	}

}
