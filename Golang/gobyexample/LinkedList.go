package main

import "fmt"

type Node struct {
	data int
	next *Node
}

// single linked list
type LinkedList struct {
	Head *Node
}

func (li *LinkedList) append(data int) {
	nextNode := &Node{data: data, next: nil}

	if li.Head == nil {
		li.Head = nextNode
		return
	}

	current := li.Head
	for current.next != nil {
		current = current.next
	}
	current.next = nextNode
}

func (li *LinkedList) display() {
	current := li.Head

	for current != nil {
		fmt.Printf("%d\n"  , *&current.data)
		current = current.next
	}
}
