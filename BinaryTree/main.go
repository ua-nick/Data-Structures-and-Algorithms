package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (tree *BinaryTree) insertItem(i int) {
	if tree.root == nil {
		tree.root = &Node{data: i}
		return
	}
	currentNode := tree.root
	for {
		if i > currentNode.data {
			if currentNode.right == nil {
				currentNode.right = &Node{data: i}
				return
			}
			currentNode = currentNode.right
		} else {
			if currentNode.left == nil {
				currentNode.left = &Node{data: i}
				return
			}
			currentNode = currentNode.left
		}
	}
}

func (tree *BinaryTree) searchItem(i int) (*Node, bool) {
	if tree.root == nil {
		return nil, false
	}
	currentNode := tree.root
	for currentNode != nil {
		if i == currentNode.data {
			return currentNode, true
		} else if i > currentNode.data {
			currentNode = currentNode.right
		} else if i < currentNode.data {
			currentNode = currentNode.left
		}
	}
	return nil, false
}

func (tree *BinaryTree) printItems(subtree *Node) {
	if subtree.left != nil {
		tree.printItems(subtree.left)
	}
	fmt.Println(subtree.data)
	if subtree.right != nil {
		tree.printItems(subtree.right)
	}
}
