package main

import "fmt"

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func main() {
	root := &Tree{Value: 1}
	root.Left = &Tree{Value: 2}
	root.Left.Left = &Tree{Value: 3}
	root.Left.Right = &Tree{Value: 4}
	root.Right = &Tree{Value: 5}

	printTree(root)
	fmt.Printf("\n ------- \n")
	result := invert(root)
	printTree(result)
	fmt.Printf("\n ------- \n")
}

func printTree(tree *Tree) {
	if tree != nil {
		fmt.Printf("%d ", tree.Value)

		printTree(tree.Left)
		printTree(tree.Right)
	}
}

func invert(tree *Tree) *Tree {
	if tree == nil {
		return nil
	}

	tree.Left, tree.Right = tree.Right, tree.Left

	tree.Left = invert(tree.Left)
	tree.Right = invert(tree.Right)

	return tree
}
