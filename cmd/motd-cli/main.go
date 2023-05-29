package main

import (
	"log"

	"github.com/sebasmannem/motd/pkg/memes"
)

func main() {
	if memeDir, err := memes.NewMemeDir("./testdata/"); err != nil {
		log.Fatalf("Houston, we have a problem: %e", err)
	} else {
		memeDir.Print()
	}
}
