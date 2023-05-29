package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

func getMemeFiles(path string) []string {
	var memeFiles []string
	var entries []os.DirEntry
	var err error
	if entries, err = os.ReadDir(path); err != nil {
		log.Fatal(err)
	} else {
		for _, e := range entries {
			memeFiles = append(memeFiles, filepath.Join(path, e.Name()))
		}
		return memeFiles
	}
	return nil
}

func printMemeFiles(path string) {
	if entries := getMemeFiles(path); entries != nil {
		for _, entry := range entries {
			fmt.Println(entry)
		}
	}
}

func getRandomMemeFile(path string) string {
	if entries := getMemeFiles(path); entries != nil {
		return entries[rand.Intn(len(entries))]
	}
	return ""
}
