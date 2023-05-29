package main

import (
	"fmt"
)

func main() {
	memeLocation := "./testdata/"
	println("Out of all these meme files:")
	printMemeFiles(memeLocation)
	fmt.Printf("This is the one:\n%s", getRandomMemeFile(memeLocation))
}
