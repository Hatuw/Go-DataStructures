package linkedlist

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item: the type of the linked list
type Item generic.Type

// Node: a single node that composes the list
type Node struct {
	content Item
	next    *Node
}

// ItemLinkedList: the linked list of Items
type ItemLinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

// Append: adds an Item to the end of the linked list
func (ll *ItemLinkedList) Append(t Item) {
	ll.lock.Lock()
	node := Node{t, nil}
	if ll.head == nil {
		ll.head = &node
	} else {
		last := ll.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &node
	}
	ll.size++
	ll.lock.Unlock()
}

// Insert: adds an Item t at position i
func (ll *ItemLinkedList) Insert(i int, t Item) error {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	if i < 0 || i > ll.size {
		return fmt.Errorf("Index out of bounds")
	}
	addNode := Node{t, nil}
	if i == 0 {
		addNode.next = ll.head
		ll.head = &addNode
	} else {
		node := ll.head
		for j := 0; j < i-2; j++ {
			node = node.next
		}
		addNode.next = node.next
		node.next = &addNode
		ll.size++
	}
	return nil
}

// RemoveAt: removes a node at position i
func (ll *ItemLinkedList) RemoveAt(i int) (*Item, error) {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	if i < 0 || i > ll.size {
		return nil, fmt.Errorf("Index out of bounds")
	}
	node := ll.head
	for j := 0; j < i-1; j++ {
		node = node.next
	}
	remove := node.next
	node.next = remove.next
	ll.size--
	return &remove.content, nil
}

// IndexOf: returns the position of the item t
func (ll *ItemLinkedList) IndexOf(t Item) int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	node := ll.head
	for j := 0; ; j++ {
		if node.content == t {
			return j
		}
		if node.next == nil {
			return -1
		}
		node = node.next
	}
}

// IsEmpty: returns true if the List is empty
func (ll *ItemLinkedList) IsEmpty() bool {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	if ll.head == nil {
		return true
	}
	return false
}

// Size: returns the linked list size
func (ll *ItemLinkedList) Size() int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	size := 1
	node := ll.head
	for ; ; size++ {
		if node == nil || node.next == nil {
			break
		}
		node = node.next
	}
	return size
}

// String: returns a string representation of the list
func (ll *ItemLinkedList) String() {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	node := ll.head
	for j := 0; ; j++ {
		if node == nil {
			break
		}
		fmt.Print(node.content)
		fmt.Print(" ")
		node = node.next
	}
	fmt.Println()
}

// Head: returns the first node, so we can iterate on it
func (ll *ItemLinkedList) Head() *Node {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	return ll.head
}
