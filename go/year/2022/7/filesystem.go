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
	root := Directory{Name: "/"}
	f := FileSystem{}
	f.Directories[root.Name] = &root
	f.CurrentDir = f.Directories[root.Name]

	for _, cmd := range commands[1:] {
		// only ls has output, this could maybe be handled better
		// if cmd.output != nil {
		// 	for _, v := range cmd.output {
		// 		v := strings.Split(v, " ")
		// 		if v[0] == "dir" {
		// 			newDir := NewDirectory(v[1])
		// 			f.Directories[newDir.Name] = &newDir
		// 		}
		// 	}
		// } else {
		// 	runCmd()
		// }
		runCmd(cmd, &f)
	}

	return f
}

func (f *FileSystem) prev() {
	f.CurrentDir = f.CurrentDir.GetParent()
}

func (f *FileSystem) next(s string) {
	f.CurrentDir = f.Directories[s]
}

// Turn this into a type
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
				newDir := NewDirectory(v[1])
				f.CurrentDir.Children[newDir.Name] = &newDir
			} else {
				name := v[0]
				size := v[1]
				newFile := NewFile(name, size, f.CurrentDir)
				f.CurrentDir.Files = append(f.CurrentDir.Files, newFile)
			}
		}
	}
}
