package bugs

import (
	"os"
	"strings"
)

func Get() ([]string, error) {
	file, err := os.ReadFile("bugs.txt")
	if err != nil {
		return nil, err
	}

	var bugString = string(file)
	bugs := strings.Split(bugString, "\n")

	return bugs, nil
}

func Create(title string) {
	buglist, err := Get()

	if err != nil {
		println("Failed to get bugs", err)
		return
	}

	buglist = append(buglist, title)

	// our get function returns a slice of strings
	// Need to convert to one long string, then convert to byte slice
	bugString := strings.Join(buglist, "\n")
	var bugBytes = []byte(bugString)

	// What's the standard when using multiple error handling functions?
	// Should I just write over the existing err, or create a new var like err2?
	// err2 := os.WriteFile("bugs.txt", bugBytes, 0777)
	// if err2 != nil {
	// 	println("Failed to write bugs to file", err)
	// }
	// or
	err = os.WriteFile("bugs.txt", bugBytes, 0777)
	if err != nil {
		println("Failed to write bugs to file", err)
	}
}
