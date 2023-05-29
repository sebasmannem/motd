package main

import (
	"log"

	imgview "github.com/sebasmannem/motd/pkg/img_view"
	"github.com/sebasmannem/motd/pkg/memes"
)

func main() {
	if memeDir, err := memes.NewMemeDir("./testdata/"); err != nil {
		log.Fatalf("Houston, we have a problem: %e", err)
	} else {
		imgview.DrawOnConsole(memeDir.GetRandomFile().String())
	}
}
