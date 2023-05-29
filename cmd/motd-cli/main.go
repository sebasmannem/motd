package main

import "log"

func main() {
	if memeDir, err := NewMemeDir("./testdata/"); err != nil {
		log.Fatalf("Houston, we have a problem: %e", err)
	} else {
		memeDir.Print()
	}
}
