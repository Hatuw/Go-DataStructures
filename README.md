# Golang Data Structures

The implementation of data structures (using Golang)

- Linked List
- Queue
- Stack

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
