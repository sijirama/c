package main

import "fmt"

type Node struct {
	Value any
	Next  *Node
}

var head *Node

func InsertLinkedList(value any) {

	node := new(Node)
	node.Value = value
	node.Next = nil

	if head == nil {
		head = node
		return
	}
	temp := head
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = node
}

func DisplayLinkedList() {
	node := head

	for node != nil {
		fmt.Println(node.Value)
		node = node.Next
	}
}
