package main

import (
	"strconv"
)

type File struct {
	Directory *Directory
	Name      string
	Size      int
}

func NewFile(name string, fileSize string, directory *Directory) File {
	size, err := strconv.Atoi(fileSize)
	if err != nil {
		panic(err)
	}

	return File{Name: name, Size: size, Directory: directory}

}

func (f File) GetParentDirectory() *Directory {
	return f.Directory
}
