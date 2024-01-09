package main

import (
	"errors"
	"fmt"
)

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
		fmt.Printf("%d\n", *&current.data)
		current = current.next
	}
}

//------------------------------------------------------------------------------

type User struct {
	name  string
	phone int
}

func getUserMap(names []string, phoneNumbers []int) (map[string]User, error) {
	var UserMap = make(map[string]User)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("Failure to create user map")
	}

	for i := 0; i < len(names); i++ {
		name := names[i]
		phonuNumber := phoneNumbers[i]

		UserMap[name] = User{
			name:  name,
			phone: phonuNumber,
		}
	}

	return UserMap, nil
}
