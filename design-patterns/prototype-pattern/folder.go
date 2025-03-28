package main

import "fmt"

type Folder struct {
	children   []INode
	foldername string
}

func (fo *Folder) print(indentation string) {
	fmt.Println(indentation + fo.foldername)
	for _, child := range fo.children {
		child.print(indentation + indentation)
	}
}

func (fo *Folder) clone() INode {
	folder := &Folder{
		children:   make([]INode, 0),
		foldername: fo.foldername,
	}
	for _, child := range fo.children {
		copyChild := child.clone()
		folder.children = append(folder.children, copyChild)
	}
	return folder
}
