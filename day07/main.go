package main

import (
	"bufio"
	"io"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

type Inode struct {
	Name string
	Size uint64
	IsDir bool
	Children map[string]*Inode
	Parent *Inode
}

func ParseFsTree(input io.Reader) *Inode {
	scanner := bufio.NewScanner(input)

	root := Inode{
		Name: "/",
		Children: make(map[string]*Inode),
		IsDir: true,
	}

	currentNode := &root
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$ cd") {
			// CD
			dirName := strings.TrimPrefix(line, "$ cd ")
			if dirName == "/" {
				currentNode = &root
			} else if dirName == ".." {
				currentNode = currentNode.Parent
			} else {
				currentNode = currentNode.Children[dirName]
			}
		} else if strings.HasPrefix(line, "dir ") {
			// dir
			dirName := strings.TrimPrefix(line, "dir ")
			newNode := Inode{
				Name: dirName,
				IsDir: true,
				Children: make(map[string]*Inode),
				Parent: currentNode,
			}
			currentNode.Children[dirName] = &newNode
		} else if !strings.HasPrefix(line, "$ ls") {
			// file
			spaceIdx := strings.Index(line, " ")
			size := utils.Atoi64OrPanic(line[0:spaceIdx])
			fileName := line[spaceIdx + 1:]
			newNode := Inode{
				Name: fileName,
				Size: uint64(size),
				IsDir: false,
				Parent: currentNode,
			}
			currentNode.Children[fileName] = &newNode
		}
	}
	return &root
}

func findDirs(dir *Inode, sizeThreshold uint64) (candidates []uint64, dirSize uint64) {
	for _, node := range dir.Children {
		if node.IsDir {
			subCandidates, subSize := findDirs(node, sizeThreshold)
			dirSize += subSize
			candidates = append(candidates, subCandidates...)
			if subSize < sizeThreshold {
				candidates = append(candidates, subSize)
			}
		} else {
			dirSize += node.Size
		}
	}
	return candidates, dirSize
}

func Part1(input io.Reader) int {

	root := ParseFsTree(input)

	dirs, _ := findDirs(root, 100000)
	sum := uint64(0)
	for _, dirSize := range dirs {
		sum += dirSize 
	}
	return int(sum)
}

func Part2(input io.Reader) int {
	totalSpace := uint64(70000000)
	requiredSpace := uint64(30000000)

	root := ParseFsTree(input)

	smallDirs, rootSize := findDirs(root, requiredSpace)
	neededFree := requiredSpace - (totalSpace - rootSize)

	candidateSize := rootSize
	for _, candidate := range smallDirs {
		if candidate >= neededFree && candidate < candidateSize {
			candidateSize = candidate
		}
	}
	
	return int(candidateSize)
}

func main() {
	utils.Run("day07.txt", Part1, Part2)
}
