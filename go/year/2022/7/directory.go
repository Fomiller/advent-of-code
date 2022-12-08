package main

import "strings"

type Directory struct {
	Name string
	Size int

	Parent   *Directory
	Children map[string]*Directory
	Next     *Directory
	Prev     *Directory

	Files []File
}

func NewDirectory(s string) Directory {
	var d Directory
	data := strings.Split(s, " ")
	d.Name = data[1]

	return d
}

func (d Directory) GetFiles() []File {
	return d.Files
}

func (d Directory) GetDirectorySize() int {
	var DirectorySize int
	for _, v := range d.Files {
		DirectorySize += v.Size
	}
	return DirectorySize
}

func (d Directory) GetParent() *Directory {
	return d.Parent
}

func (d *Directory) AddDirectory(newDir *Directory) {
	d.Children[d.Name] = newDir
}

func (d *Directory) AddFile(newFile File) {
	d.Files = append(d.Files, newFile)
}
