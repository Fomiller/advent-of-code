package main

import "fmt"

type Directory struct {
	Name string
	Size int

	Parent   *Directory
	Children map[string]*Directory

	Files []File
}

func NewDirectory(name string, parent *Directory) Directory {
	var d Directory
	d.Name = name
	d.Children = make(map[string]*Directory)
	d.Parent = parent

	return d
}

func (d Directory) GetFiles() []File {
	return d.Files
}

func (d *Directory) GetDirectorySize() int {
	for _, v := range d.Files {
		d.Size += v.Size
	}

	if d.HasChildren() {
		fmt.Println("children")
		for _, v := range d.Children {
			d.Size += v.Size
		}
	}

	return d.Size
}

func (d Directory) GetParent() *Directory {
	return d.Parent
}

func (d Directory) HasChildren() bool {
	if d.Children == nil {
		return false
	} else {
		return true
	}
}

func (d *Directory) AddDirectory(newDir *Directory) {
	d.Children[d.Name] = newDir
}

func (d *Directory) AddFile(newFile File) {
	d.Files = append(d.Files, newFile)
}

func (d *Directory) walk() {
	if d.HasChildren() {
		for _, v := range d.Children {
			v.GetDirectorySize()
			v.walk()
		}
	}
}
