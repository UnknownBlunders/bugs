package main

import "github.com/unkownblunders/bugs"

func main() {

	buglist := bugs.Get()

	println("Header")

	for bug, value := range buglist {
		println(bug, value)
	}

}
