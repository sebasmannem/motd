package main

import (
	"log"
	"os"

	imgview "github.com/sebasmannem/motd/pkg/img_view"
	kitty "github.com/sebasmannem/motd/pkg/kitty-display"
	"github.com/sebasmannem/motd/pkg/memes"
)

func main() {
	if memeDir, err := memes.NewMemeDir("./testdata/"); err != nil {
		log.Fatalf("Houston, we have a problem: %e", err)
	} else {
		path := memeDir.GetRandomFile().String()
		if _, is_defined := os.LookupEnv("KITTY_WINDOW_ID"); is_defined {
			kitty.FromFile(path)
		} else {
			imgview.DrawOnConsole(path)
		}
	}
}
