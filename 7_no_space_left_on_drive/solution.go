package main

import (
	"bufio"
	"fmt"

	"github.com/ecuyle/advent-of-code-2022/utils"
)

type node struct {
	name     string
	kind     string
	size     int
	children []*node
	parent   *node
}

func makeNode(name string, kind string, size int) *node {
	return &node{name, kind, size, []*node{}, nil}
}

func updateParentSize(parent *node, size int) {
	if parent == nil {
		return
	}

	parent.size += size
	updateParentSize(parent.parent, size)
}

func addChild(parent *node, child *node) {
	fmt.Println("Adding parent to child", parent.name, child.name)
	parent.children = append(parent.children, child)
	child.parent = parent
	updateParentSize(parent, child.size)
}

func getRoot(node *node) *node {
	if node.parent == nil {
		return node
	}

	return getRoot(node.parent)
}

func cd(targetDir string, cwd *node) (*node, error) {
	fmt.Println("Attempting to cd:", cwd.name, targetDir)

	if targetDir == "/"{
		return getRoot(cwd), nil
	}

	if targetDir == "." {
		return cwd, nil
	}

	if targetDir == ".." {
		fmt.Println("Going up to", cwd.parent.name)
		return cwd.parent, nil
	}

	for _, child := range cwd.children {
		if child.name == targetDir {
			return child, nil
		}
	}

	return nil, fmt.Errorf("could not change directory to %s", targetDir)
}

func createFsTree(inputPath string) *node {
	file := utils.Readfile(inputPath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	root := makeNode("/", "d", 0)
	cwd := root

	// important learning lesson today:
	// fmt.Println(a)  // < pointer to `a`
	// fmt.Println(*a) // < value at pointer `a`
	// fmt.Println(&a) // < address at pointer `a`

	for scanner.Scan() {
		line := scanner.Text()

		if line[0:4] == "$ cd" {
			var targetDir string
			_, err := fmt.Sscanf(line, "$ cd %s", &targetDir)
			utils.CheckError(err)
			dir, err := cd(targetDir, cwd)
			utils.CheckError(err)
			cwd = dir

			fmt.Println("Changed current directory to", cwd)
		} else if line[0:4] != "$ ls" {
			var size int
			var filename string
			fmt.Sscanf(line, "%d %s", &size, &filename)
			fmt.Sscanf(line, "dir %s", &filename)
			var newNode *node

			if size == 0 {
				newNode = makeNode(filename, "d", size)
			} else {
				newNode = makeNode(filename, "f", size)
			}

			addChild(cwd, newNode)
		}
	}

	return root
}

func findSumOfDirsUnderSize(root *node, maxSize int) int {
	sum := 0

	if root.size <= maxSize && root.kind == "d" {
		sum += root.size
	}

	for _, children := range root.children {
		sum += findSumOfDirsUnderSize(children, maxSize)
	}

	return sum
}

func partOne(inputPath string) int {
	root := createFsTree(inputPath)
	return findSumOfDirsUnderSize(root, 100000)
}

func partTwo(inputPath string) {
	file := utils.Readfile(inputPath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func main() {
	fmt.Println(partOne("./input_test.txt"))
	fmt.Println(partOne("./input.txt"))
	// partTwo("./input_test.txt")
}
