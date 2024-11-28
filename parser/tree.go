package parser

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Node struct {
	Name     string
	IsDir    bool
	Children []*Node
}

func buildTree(paths []string) *Node {
	root := &Node{Name: ".", IsDir: true}
	for _, path := range paths {
		parts := strings.Split(path, string(os.PathSeparator))
		currentNode := root
		for i, part := range parts {
			// Check if part already exists as a child
			var childNode *Node
			for _, child := range currentNode.Children {
				if child.Name == part {
					childNode = child
					break
				}
			}
			if childNode == nil {
				// Create new node
				isDir := i < len(parts)-1
				childNode = &Node{Name: part, IsDir: isDir}
				currentNode.Children = append(currentNode.Children, childNode)
			}
			currentNode = childNode
		}
	}
	return root
}

func printTree(node *Node, prefix string, isLast bool) {
	if node.Name != "." {
		fmt.Print(prefix)
		if isLast {
			fmt.Print("└── ")
			prefix += "    "
		} else {
			fmt.Print("├── ")
			prefix += "│   "
		}
		fmt.Println(node.Name)
	}
	// Sort children by name
	sort.Slice(node.Children, func(i, j int) bool {
		return node.Children[i].Name < node.Children[j].Name
	})
	for i, child := range node.Children {
		printTree(child, prefix, i == len(node.Children)-1)
	}
}

func countNodes(node *Node) (dirs int, files int) {
	if node.IsDir {
		dirs++
	} else {
		files++
	}
	for _, child := range node.Children {
		d, f := countNodes(child)
		dirs += d
		files += f
	}
	return
}
