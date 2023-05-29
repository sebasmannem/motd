package memes

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

type MemeFile struct {
	name string
	path string
}

func (mf MemeFile) String() string {
	return filepath.Join(mf.path, mf.name)
}

type MemeFiles []MemeFile

func (mfs MemeFiles) getRandom() MemeFile {
	return mfs[rand.Intn(len(mfs))]
}

type MemeDir struct {
	path  string
	files MemeFiles
}

func NewMemeDir(path string) (*MemeDir, error) {
	md := MemeDir{
		path: path,
	}
	if md.getFiles() != nil {
		return &md, nil
	}
	return nil, fmt.Errorf("folder does not contain memes!!!")
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
			md.files = append(md.files, MemeFile{
				path: md.path,
				name: entry.Name(),
			})
		}
		return md.files
	}
	return nil
}

func (md MemeDir) Print() {
	fmt.Printf("Out of all the meme files in %s:\n", md.path)
	for _, entry := range md.files {
		fmt.Printf("- %s\n", entry)
	}
	// md.GetRandomFile() returns a MemeFile, but that has a String() method,
	// which is automatically used by fmt.Printf("%s")
	fmt.Printf("This is the one:\n%s", md.GetRandomFile())
}

func (md MemeDir) GetRandomFile() MemeFile {
	return md.files.getRandom()
}
