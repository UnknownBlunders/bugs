package main

import (
	"os"

	"github.com/unkownblunders/bugs"
)

func list() {

	buglist, err := bugs.Get("bugs.txt")

	if err != nil {
		println("Failed to get bugs", err)
		return
	}

	println("# Status Title")
	println("=================")

	for index, value := range buglist {
		println(index, value.Status, "	", value.Title)
	}

}

func main() {

	if len(os.Args) >= 2 {
		for _, title := range os.Args[1:] {
			bugs.Create("bugs.txt", title, "Open")
		}
	}
	list()

}
