package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

func initFile(name string, size int) *File {
	f := new(File)
	f.name = name
	f.size = size
	return f
}

type Dir struct {
	name       string
	childDirs  []*Dir
	childFiles []*File
	parent     *Dir
	size       int
}

func initDir(name string) *Dir {
	d := new(Dir)
	d.name = name
	return d
}

func (d *Dir) addChildDir(child *Dir) {
	child.parent = d
	d.childDirs = append(d.childDirs, child)
}

func (d *Dir) addChildFile(child *File) {
	d.childFiles = append(d.childFiles, child)
}

func (d *Dir) findChildDir(name string) *Dir {
	for _, child := range d.childDirs {
		if child.name == name {
			return child
		}
	}
	return nil
}

func (d *Dir) getSize() int {
	size := 0
	for _, child := range d.childFiles {
		size += child.size
	}
	for _, child := range d.childDirs {
		size += child.getSize()
	}
	d.size = size
	return d.size
}

func (d *Dir) dfs() (int, int, []int) {
	size, res, l, tmp, tmp2, tmp3 := 0, 0, []int{}, 0, 0, []int{}
	for _, child := range d.childFiles {
		size += child.size
	}
	for _, child := range d.childDirs {
		tmp, tmp2, tmp3 = child.dfs()
		size += tmp
		res += tmp2
		l = append(l, tmp3...)
	}
	if size <= 100000 {
		res += size
	}
	l = append(l, size)
	return size, res, l
}

func main() {
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("example.txt")
	data := strings.Split(strings.Trim(string(input), "\n"), "\n")

	// init root
	curr := initDir("/")
	root := curr

	for _, line := range data[1:] {
		// fmt.Println(line)
		args := strings.Split(line, " ")
		if args[0] == "$" {
			if args[1] == "cd" {
				if args[2] == ".." {
					//  cd into parent
					curr = curr.parent
					// fmt.Printf("%s has size %d\n", curr.name, curr.getSize())
				} else {
					// cd into child
					curr = curr.findChildDir(args[2])
				}
			}
		} else {
			// ls output
			if args[0] == "dir" {
				child := initDir(args[1])
				curr.addChildDir(child)
			} else {
				size, _ := strconv.Atoi(args[0])
				child := initFile(args[1], size)
				curr.addChildFile(child)
			}
		}
	}

	tot, res, l := root.dfs()
	sort.Ints(l)
	diff := 30000000 - (70000000 - tot)

	// fmt.Printf("total size: %d, result: %d\n", tot, res)
	fmt.Println(res)

	for _, v := range l {
		if v >= diff {
			fmt.Println(v)
			break
		}
	}

}
