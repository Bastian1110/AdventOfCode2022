package main

import (
	"bufio"
	"fmt"
	"os"
)

type Diectory struct {
	name   string
	root   *Diectory
	files  []File
	childs []Diectory
	size   int
}

func (dir *Diectory) init(root *Diectory, name string) *Diectory {
	dir.name = name
	dir.root = root
	dir.childs = []Diectory{}
	dir.files = []File{}
	dir.size = 0
	return dir
}

func NewDirectory(root *Diectory, name string) *Diectory {
	return new(Diectory).init(root, name)
}

func (dir *Diectory) AddFile(temp *File) {
	dir.files = append(dir.files, *temp)
	dir.size += temp.size
}

func (dir *Diectory) AddChild(temp *Diectory) {
	dir.childs = append(dir.childs, *temp)
	dir.size += temp.size
}

func (dir *Diectory) Path() {
	if dir.root != nil {
		fmt.Println(dir.name)
		dir.root.Path()
	} else {
		fmt.Println("/")
	}
}

func (dir *Diectory) Print(tab int) {
	for i := 0; i < tab*2; i++ {
		fmt.Print(" ")
	}
	fmt.Println(">", dir.name)
	for i := 0; i < tab*2; i++ {
		fmt.Print(" ")
	}
	fmt.Println("Size : ", dir.size)
	for _, c := range dir.files {
		for i := 0; i < tab*2; i++ {
			fmt.Print(" ")
		}
		fmt.Println("-", c.name, "Size : ", c.size)
	}
	for _, c := range dir.childs {
		c.Print(tab + 1)
	}
}

type File struct {
	name string
	size int
}

func (fil *File) init(name string, size int) *File {
	fil.name = name
	fil.size = size
	return fil
}

func NewFile(name string, size int) *File {
	return new(File).init(name, size)
}

func ReadFile(path string) []string {
	var res []string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Cant open the file 😬")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

func summ(arr []int) int {
	total := 0
	for _, a := range arr {
		total += a
	}
	return total
}

func solutionOne(data []string) {

}

func solutionTwo(data []string) {

}

func main() {
	lines := ReadFile("test.txt")
	solutionOne(lines)

	file1 := NewFile("one", 1)
	root := NewDirectory(nil, "root")
	root.AddFile(file1)

	file2 := NewFile("two", 2)
	child := NewDirectory(root, "child")
	child.AddFile(file2)

	file3 := NewFile("tre", 3)
	miau := NewDirectory(child, "miau")
	miau.AddFile(file3)

	child.AddChild(miau)
	root.Print(0)
	root.AddChild(child)
	root.Print(0)

}
