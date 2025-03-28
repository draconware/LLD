package main

import "fmt"

type Component interface {
	search(string)
}

type Folder struct {
	components []Component
	name       string
}

func (fo *Folder) search(key string) {
	fmt.Printf("Searching recursively for keyword: %s in folder: %s.\n", key, fo.name)
	for _, child := range fo.components {
		child.search(key)
	}
}

func (fo *Folder) add(component Component) {
	fo.components = append(fo.components, component)
}

type File struct {
	name string
}

func (fi *File) search(key string) {
	fmt.Printf("Searching for keyword: %s in file: %s.\n", key, fi.name)
}
