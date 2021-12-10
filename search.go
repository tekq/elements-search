package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

func main() {
	var files []string
	var searched string
	var success bool

	searched = os.Args[1]
	root := "/etc/elements/repos/"
	err := filepath.Walk(root, visit(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if strings.Contains(file, searched) {
			fmt.Println("It worked")
			fmt.Print("File searched: ")
			fmt.Print(file)
			fmt.Println()
			success = true
		}

	}
	if success {
		fmt.Println("File found. Good job on the code.")
	} else {
		fmt.Println("File not found")
		log.Fatal(err)
	}
}
