package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (b *BinaryTree) insertNode(value int) {
	if b.root == nil {
		b.root = &Node{value: value}
	} else {
		b.root.insertNode(value)
	}
}

func (n *Node) insertNode(value int) {
	if value < n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insertNode(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insertNode(value)
		}
	}
}

func (b *BinaryTree) removeNode(value int) {
	if b.root != nil {
		b.root = b.root.removeNode(value)
	}
}

func (n *Node) removeNode(value int) *Node {
	if n == nil {
		return nil
	}
	if value == n.value {
		if n.left == nil && n.right == nil {
			return nil
		}
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		tempNode := n.right
		for tempNode.left != nil {
			tempNode = tempNode.left
		}
		n.value = tempNode.value
		n.right = n.right.removeNode(tempNode.value)
	} else if value < n.value {
		n.left = n.left.removeNode(value)
	} else {
		n.right = n.right.removeNode(value)
	}
	return n
}

func (b *BinaryTree) printTree() {
	if b.root != nil {
		b.root.printTree(0)
	}
}

func (n *Node) printTree(level int) {
	if n != nil {
		n.right.printTree(level + 1)
		for i := 0; i < level; i++ {
			fmt.Print("   ")
		}
		fmt.Printf("%d\n", n.value)
		n.left.printTree(level + 1)
	}
}

func main() {
	tree := BinaryTree{}
	tree.insertNode(1)
	tree.insertNode(2)
	tree.insertNode(3)
	tree.insertNode(5)
	tree.insertNode(4)
	tree.printTree()
	tree.removeNode(3)
	tree.printTree()
}
