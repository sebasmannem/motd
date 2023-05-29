package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func printMemeFiles(path string) {
	var entries []os.DirEntry
	var err error
	if entries, err = os.ReadDir(path); err != nil {
		log.Fatal(err)
	} else {
		for _, e := range entries {
			fmt.Println(filepath.Join(path, e.Name()))
		}
	}
}

func main() {
	printMemeFiles("./testdata/")
}
