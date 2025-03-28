package main

import "fmt"

func main() {
	file1 := &File{filename: "File1"}
	file2 := &File{filename: "File2"}
	file3 := &File{filename: "File3"}
	file4 := &File{filename: "File4"}

	folder1 := &Folder{
		foldername: "Folder1",
		children:   []INode{file1},
	}

	folder2 := &Folder{
		foldername: "Folder2",
		children:   []INode{folder1, file2, file3, file4},
	}

	fmt.Println("Printing hirearchy for folder 2")
	folder2.print(" ")

	clonedFolder := folder2.clone()
	fmt.Println("Printing hirearchy for cloned folder 2")
	clonedFolder.print(" ")
}
