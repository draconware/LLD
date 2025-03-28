package main

import "fmt"

type File struct {
	filename string
}

func (fi *File) print(indentation string) {
	fmt.Println(indentation + fi.filename)
}

func (fi *File) clone() INode {
	return &File{filename: fi.filename}
}
