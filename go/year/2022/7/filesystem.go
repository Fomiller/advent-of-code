package main

import (
	"log"
	"strings"
)

type FileSystem struct {
	CurrentDir  *Directory
	Directories map[string]*Directory
}

type command struct {
	input  []string
	output []string
}

func NewFileSystem(commands []command) FileSystem {
	root := NewDirectory("/", nil)
	f := FileSystem{}
	f.Directories = make(map[string]*Directory)
	f.Directories[root.Name] = &root
	f.CurrentDir = f.Directories[root.Name]

	for _, cmd := range commands[1:] {
		runCmd(cmd, &f)
	}

	return f
}

func (f *FileSystem) prev() {
	if f.CurrentDir != nil {
		f.CurrentDir = f.CurrentDir.GetParent()
	} else {
		log.Println("No parent Dir.")
	}
}

func (f *FileSystem) next(s string) {
	f.CurrentDir = f.CurrentDir.Children[s]
}

// Turn this into a type
// create command type with cmd.run()
func runCmd(cmd command, f *FileSystem) {
	if strings.Contains(cmd.input[0], "cd") {
		f.cdCmd(cmd)
	} else if strings.Contains(cmd.input[0], "ls") {
		f.lsCmd(cmd)
	} else {
		log.Printf("invalid command: %v", cmd)
	}
}

func (f *FileSystem) cdCmd(cmd command) {
	if cmd.input[1] == ".." {
		f.prev()
	} else {
		f.next(cmd.input[1])
	}

}

func (f *FileSystem) lsCmd(cmd command) {
	if cmd.output != nil {
		for _, v := range cmd.output {
			v := strings.Split(v, " ")

			if v[0] == "dir" {
				newDir := NewDirectory(v[1], f.CurrentDir)
				f.CurrentDir.Children[newDir.Name] = &newDir

			} else {
				size := v[0]
				name := v[1]
				newFile := NewFile(name, size, f.CurrentDir)
				f.CurrentDir.Files = append(f.CurrentDir.Files, newFile)
			}
		}
	}
}

func (f *FileSystem) GetLargeDirs() []Directory {
	var largeDirs []Directory
	for _, v := range f.CurrentDir.Children {
		if v.Size <= 100000 {
			largeDirs = append(largeDirs, *v)
		}
	}
	return largeDirs
}
