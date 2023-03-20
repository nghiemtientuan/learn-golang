package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// Your solution goes here. Good luck!
	fmt.Print(strings.Join(listFiles("testdata"), " "))
}

func listFiles(dirname string) []string {
	var dirs []string

	files, _ := ioutil.ReadDir(dirname)

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
