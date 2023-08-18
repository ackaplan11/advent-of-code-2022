package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	Size     int
	Parent   *Node
	Children map[string]*Node
}

func main() {
	root := Node{ // create root node
		Size:     0,
		Parent:   nil,
		Children: map[string]*Node{},
	}
	curr := &root // create pointer to current node

	content, _ := ioutil.ReadFile("day7/input.txt")
	terminalOutputs := strings.Split(string(content), "\n")
	goToRootRegex := regexp.MustCompile(`^\$ cd /$`)
	goToParentRegex := regexp.MustCompile(`^\$ cd \.\.$`)
	goToChildRegex := regexp.MustCompile(`^\$ cd ([a-z]+)$`)
	dirRegex := regexp.MustCompile(`^dir\s(\S+)`)
	fileRegex := regexp.MustCompile(`^(\d+) (.+)$`)

	for _, output := range terminalOutputs {

		if goToRootRegex.MatchString(output) {
			curr = &root // set pounter to root
		} else if goToParentRegex.MatchString(output) {
			curr = curr.Parent // set pointer to parent
		} else {
			goToChildMatch := goToChildRegex.FindStringSubmatch(output)
			dirMatch := dirRegex.FindStringSubmatch(output)
			fileMatch := fileRegex.FindStringSubmatch(output)

			if len(goToChildMatch) > 0 {
				key := goToChildMatch[1]
				curr = curr.Children[key] // set pointer to child
			} else if len(dirMatch) > 0 {
				key := dirMatch[1]
				childNode := Node{
					Size:     0,
					Parent:   curr,
					Children: map[string]*Node{},
				}
				curr.Children[key] = &childNode // add new child node with key name
			} else if len(fileMatch) > 0 {
				sizeStr := fileMatch[1]
				size, err := strconv.Atoi(sizeStr)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				curr.Size += size // add size to current node
			}
		}
	}

	curr = &root
	calculateSize(curr)
	var totalSize int
	sumSizes(&totalSize, curr)
	fmt.Println(findSizeToDelete(&root))
}

func calculateSize(node *Node) int {
	if len(node.Children) == 0 {
		return node.Size
	} else {
		for _, child := range node.Children {
			node.Size += calculateSize(child)
		}
	}
	return node.Size
}

func sumSizes(totalSize *int, node *Node) {
	if node.Size <= 100000 {
		*totalSize += node.Size
	}
	if len(node.Children) == 0 {
		return
	}
	for _, child := range node.Children {
		sumSizes(totalSize, child)
	}
}

func findSizeToDelete(root *Node) int {
	minSize := math.Abs(float64(40000000 - root.Size))
	deletedSize := root.Size
	curr := root
	queue := convertMapToSlice(root.Children)
	for len(queue) > 0 {
		curr = queue[0]
		queue = queue[1:]
		if float64(curr.Size) > minSize && curr.Size < deletedSize {
			deletedSize = curr.Size
		}
		addToQueue := convertMapToSlice(curr.Children)
		for _, child := range addToQueue {
			queue = append(queue, child)
		}

	}
	return deletedSize
}

func convertMapToSlice(children map[string]*Node) []*Node {
	queue := []*Node{}
	for _, value := range children {
		queue = append(queue, value)
	}
	return queue
}
