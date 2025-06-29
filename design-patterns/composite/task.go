/*
🧪 Your Challenge (Interview Style)
You're building a file explorer.

Your Task:
Design a structure that lets you call Display(indentLevel int) on:

Files

Folders (which can contain files and other folders)

type FileSystemItem interface {
	Display(indentLevel int)
}
Each file prints like: -- document.txt
Each folder prints its name and recursively prints its children with indentation.

Example Output:

📁 root
  -- file1.txt
  📁 subfolder
    -- file2.txt
*/

package main

import (
	"fmt"
	"strings"
)

type FileSystemItem interface {
	Display(indentLevel int)
}

type File struct {
	name string
}

func (f File) Display(indentLevel int) {
	fmt.Println(strings.Repeat(" ", indentLevel) + f.name)
}

type Folder struct {
	name string
	item []FileSystemItem
}

func (f Folder) Display(indentLevel int) {
	fmt.Println(strings.Repeat(" ", indentLevel) + f.name)
	for _, item := range f.item {
		item.Display(indentLevel + 1)
	}
}

func main() {
	file1 := File{name: "file1.txt"}
	file2 := File{name: "file2.txt"}
	folder1 := Folder{name: "folder1", item: []FileSystemItem{file1, file2}}
	root := Folder{name: "root", item: []FileSystemItem{folder1}}
	root.item = append(root.item, File{name: "root.txt"})
	root.item = append(root.item, []FileSystemItem{folder1}...)
	root.Display(0)
}
