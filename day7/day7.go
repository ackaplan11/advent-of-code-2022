package main

import (
	"fmt"
	"io/ioutil"
)

type Node struct {
	Name     string
	Size     int
	Parent   *Node
	Children []*Node
}

var root Node = Node{
	Name:     "/",
	Size:     0,
	Parent:   nil,
	Children: []*Node{},
}

var pointer *Node = &root

func main() {
	fmt.Println("Advent of Code")
	terminalOutputs, _ := ioutil.ReadFile("day6/input.txt")
	for _, output := range terminalOutputs {
		parseOutput(output)
	}
}

func parseOutput(output []byte) {
	str := string(output)
	if str == "$ cd /" {
		pointer = &root
	} else if str[3:] == "dir" {
		addChild(*pointer, str[4:])
	}

}

func addSize(node Node, size int) {
	node.Size += size
}

func addChild(node Node, childName string) {
	child := Node{
		Name:     childName,
		Size:     0,
		Parent:   &node,
		Children: []*Node{},
	}
	_ = append(node.Children, &child)
}

func calculateSize(node Node) int {
	if len(node.Children) == 0 {
		return node.Size
	}
	for _, child := range node.Children {
		node.Size += calculateSize(*child)
	}
	return node.Size
}
