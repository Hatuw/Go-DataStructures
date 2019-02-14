// Package binarysearchtree creates a ItemBinarySearchTree data structure for the Item type
package binarysearchtree

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item: the type fo the binary search tree
type Item generic.Type

// Node: a single node that composes the tree
type Node struct {
	key int
	value Item
	left *Node
	right *Node
}

// ItemBinarySearchTree: the binary search tree of Items
type ItemBinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

// Insert: inserts the Item t in the tree
func (bst *ItemBinarySearchTree) Insert(key int, value Item) {

}

// internal function to find the correct place for a node in a tree
func insertNode(node, newNode *Node) {

}

// InOrderTraverse: visits all nodes with in-order traversing
func (bst *ItemBinarySearchTree) InOrderTraverse(f func(Item)) {

}

// internal recursive function to traverse in order
func inOrderTraverse(n *Node, f func(Item)) {

}

// PreOrderTraverse: visits all nodes with pre-order traversing
func (bst *ItemBinarySearchTree) PreOrderTraverse(f func(Item)) {

}

// internal recursive function to traverse pre order
func preOrderTraverse(n *Node, f func(Item)) {

}

// PostOrderTraverse: visits all nodes with post_order traversing
func (bst *ItemBinarySearchTree) PostOrderTraverse(f func(Item)) {

}

// internal recursive function to traverse post order
func postOrderTraverse(n *Node, f func(Item)) {

}

// Min: returns the Item with min value stored in the tree
func (bst *ItemBinarySearchTree) Min() *Item {
	return &bst.root.value
}

// Max: returns the Item with max value stored in the tree
func (bst *ItemBinarySearchTree) Max() *Item {
	return &bst.root.value
}

// Search: returns true if the Item t exists in the tree
func (bst *ItemBinarySearchTree) Search(key int) bool {
	return false
}

// internal recursive function to search an item in the tree
func search(n *Node, key int) bool {
	return false
}

// Remove: removes the Item with key `key` from the tree
func (bst *ItemBinarySearchTree) Remove(key int) {

}

// internal recursive function to remove an item
func remove(node *Node, key int) *Node {
	return node
}

// String: prints a visual representation of the tree
func (bst *ItemBinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")

}

// internal recursive function to print a tree
func stringify(n *Node, level int) {

}
