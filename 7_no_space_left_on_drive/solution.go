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

	if targetDir == "/" {
		return getRoot(cwd), nil
	}

	if targetDir == "." {
		return cwd, nil
	}

	if targetDir == ".." {
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

	for _, child := range root.children {
		sum += findSumOfDirsUnderSize(child, maxSize)
	}

	return sum
}

func partOne(inputPath string) int {
	root := createFsTree(inputPath)
	return findSumOfDirsUnderSize(root, 100000)
}

func findSmallestDirToDelete(root *node, targetSize int) (*node, error) {
	if root.kind == "f" || root.size < targetSize {
		return nil, fmt.Errorf("could not find directory smaller than %d", targetSize)
	}

	closest := root

	for _, child := range root.children {
		result, err := findSmallestDirToDelete(child, targetSize)

		if err != nil {
			continue
		}

		if result.size < closest.size {
			closest = result
		}
	}

	return closest, nil
}

func partTwo(inputPath string) int {
	const TOTAL_DISK_SPACE = 70000000
	const REQUIRED_FREE_SPACE = 30000000
	root := createFsTree(inputPath)
	unusedFreeSpace := TOTAL_DISK_SPACE - root.size
	dir, err := findSmallestDirToDelete(root, REQUIRED_FREE_SPACE-unusedFreeSpace)
	utils.CheckError(err)

	return dir.size
}

func main() {
	fmt.Println("Part 1 Test:", partOne("./input_test.txt"))
	fmt.Println("Part 1 Actual:", partOne("./input.txt"))
	fmt.Println("Part 2 Test:", partTwo("./input_test.txt"))
	fmt.Println("Part 2 Actual:", partTwo("./input.txt"))
}
