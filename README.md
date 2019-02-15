# Golang Data Structures

The implementation of data structures (using Golang)

- Linked List
- Queue
- Stack
- Set
- Dictionary
- Hash Table
- Binary Search Tree

## Linked List

The linked list provides these methods:

`Append(t Item)` : adds an Item t to the end of the linked list

`Insert(i int, t Item)` : adds an Item t at position i

`RemoveAt(i int)` : removes a node at position i

`IndexOf()` : returns the position of the Item t

`IsEmpty()` : returns true if the list is empy

`Size()` : returns the linked list size

`String()` : returns a string representation of the list

`Head()` : returns the first node, so we can iterate on it

## Queue

The queue provides these methods:

`New()` : creates a new ItemQueue

`EnQueue(t Item)` : adds an Item to the end of the queue

`Dequeue()` : removes an Item from the start of the queue

`Front()` : returns the item next in the queu, without removing it

`IsEmpty()` : returns true if the queue is empty

`Size()` : returns the number of Items in the queue

## Stack

`New()` : creates a new ItemStack

`Push(t Item)` : adds an Item to the top of the stack

`Pop()` : removes an Item from the top of the stack

## Set

`Add(t Item)` : adds a new element to the Set

`Clear()` : removes all elements from the Set

`Delete(item Item)` : removes the Item from the Set and returns Has(Item)

`Has(item Item)` : returns true if the Set contains the Item

`Items()` : returns the Item(s) stored

`Size()` : returns the size of the set

## Dictionary

`Set(k Key, v Value)` : adds a new item to the dictionary

`Delete(k Key)` : removes a value from the dictionary, given its key

`Has(k Key)` : returns true if the key exists in the dictionary

`Get(k Key)` : returns the value associated with the key

`Clear()` : removes all the items from the dictionary

`Keys()` : returns a slice of all the keys present

`Values()` : returns a slice of all the values present

## Hash Table

`Put(k Key, v Value)` : put item with value v and key k into the hashtable

`Remove(k Key)` : remove item with key k from hashtable

`Get(k Key)` : get item with key k from the hashtable

`Size()` : returns the number of the hashtable elements

## Binary Search Tree

`Insert(key int, value Item)` : inserts the Item t in the tree
`Search(key int)` : returns true if the Item t exists in the tree
`InOrderTraverse()` : visits all nodes with in-order traversing
`PreOrderTraverse()` : visits all nodes with pre-order traversing
`PostOrderTraverse()` : visits all nodes with post-order traversing
`Min()` : returns the Item with min value stored in the tree
`Max()` : returns the Item with max value stored in the tree
`Remove(key int)` : removes the Item t from the tree
`String()` : prints a CLI readable rendering of the tree
