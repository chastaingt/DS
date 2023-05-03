package main

import "fmt"

type node struct {
	key   int
	left  *node
	right *node
}

func (n *node) insert(k int) {
	if n.key < k {
		if n.right == nil {
			n.right = &node{key: k}
		} else {
			n.right.insert(k)
		}
	} else if n.key > k {
		if n.left == nil {
			n.left = &node{key: k}
		} else {
			n.left.insert(k)
		}
	}
}

func (n *node) search(k int) bool {

	if n == nil {
		return false
	}
	if n.key < k {
		n.right.search(k)
	} else if n.key > k {
		n.left.search(k)
	}
	return true
}

func main() {
	tree := node{key: 100}
	tree.insert(89)
	tree.insert(198)
	tree.insert(100)
	tree.insert(2)
	tree.insert(7)

	fmt.Println(tree)
	fmt.Println(tree.search(33))
}
