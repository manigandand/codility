package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

// insert
func (n *Node) Insert(k int) {
	if n.key < k {
		// move to right
		if n.right == nil {
			n.right = &Node{key: k}
		} else {
			n.right.Insert(k)
		}
	}

	if n.key > k {
		// move to left
		if n.left == nil {
			n.left = &Node{key: k}
		} else {
			n.left.Insert(k)
		}
	}
}

// search
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.key < k {
		// move to right
		return n.right.Search(k)
	} else if n.key > k {
		// move to left
		return n.left.Search(k)
	}

	return true
}

func main() {
	tree := &Node{key: 100}
	tree.Insert(250)
	tree.Insert(150)
	tree.Insert(50)
	tree.Insert(10)
	tree.Insert(20)

	fmt.Printf("%+v\n", tree)

	fmt.Println(tree.Search(30))
}
