package main

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}
	file4 := &File{name: "File4"}

	folder1 := &Folder{
		name:       "Folder1",
		components: []Component{file1, file2},
	}

	folder2 := &Folder{
		name:       "Folder2",
		components: []Component{folder1},
	}
	folder2.add(file3)
	folder2.add(file4)

	folder2.search("rose")
}
