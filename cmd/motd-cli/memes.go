package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

type MemeFile string
type MemeFiles []MemeFile
type MemeDir struct {
	path  string
	files MemeFiles
}

func NewMemeDir(path string) MemeDir {
	return MemeDir{
		path: path,
	}
}

func (md *MemeDir) getFiles() MemeFiles {
	if md.files != nil {
		return md.files
	}
	var entries []os.DirEntry
	var err error
	if entries, err = os.ReadDir(md.path); err != nil {
		log.Fatal(err)
	} else {
		for _, entry := range entries {
			md.files = append(md.files, MemeFile(filepath.Join(md.path, entry.Name())))
		}
		return md.files
	}
	return nil
}

func (md *MemeDir) Print() {
	if entries := md.getFiles(); entries == nil {
		fmt.Printf("Location %s does no contain memes\n", md.path)
	} else {
		fmt.Printf("Out of all the meme files in %s:\n", md.path)
		for _, entry := range entries {
			fmt.Printf("- %s\n", entry)
		}
		fmt.Printf("This is the one:\n%s", md.GetRandomFile())
	}
}

func (md MemeDir) GetRandomFile() MemeFile {
	return md.files.getRandom()
}

func (mfs MemeFiles) getRandom() MemeFile {
	return mfs[rand.Intn(len(mfs))]
}
