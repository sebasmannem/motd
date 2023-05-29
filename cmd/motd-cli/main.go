package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var entries []os.DirEntry
	var err error
	if entries, err = os.ReadDir("./testdata/"); err != nil {
		log.Fatal(err)
	} else {
		for _, e := range entries {
			fmt.Println(filepath.Join("./testdata", e.Name()))
		}
	}
}
