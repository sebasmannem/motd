package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	entries, err := os.ReadDir("./testdata/")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(filepath.Join("./testdata", e.Name()))
	}
}
