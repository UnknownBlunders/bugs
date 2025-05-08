package main

import (
	"os"

	"github.com/unkownblunders/bugs"
)

func list() {

	buglist, err := bugs.Get()

	if err != nil {
		println("Failed to get bugs", err)
		return
	}

	println("# Title")

	for bug, value := range buglist {
		println(bug, value)
	}

}

func main() {

	if len(os.Args) >= 2 {
		for _, title := range os.Args[1:] {
			bugs.Create(title)
		}

	}

	list()

}
