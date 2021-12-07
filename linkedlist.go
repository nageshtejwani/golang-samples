package main

import (
	"fmt"
)

type node struct {
	val  int
	next *node
}

type linkedlist struct {
	length int
	head   *node
	tail   *node
}

func insertNode(n *node, list *linkedlist) {
	fmt.Println(n)
	if list.head == nil {
		list.head = n
		list.tail = n
		fmt.Println("Inside if")
	} else {
		fmt.Println("Inside else")
		temp := list.tail
		list.tail = nil
		temp.next = n
		list.length = list.length + 1
	}
	fmt.Println(list)
}

func printList(list *linkedlist) {
	node := list.head
	fmt.Println(node.val)
	for node.next != nil {
		node = node.next
		fmt.Println(node.val)
	}
}

func main() {
	ll := linkedlist{0, nil, nil}
	n := node{1, nil}
	insertNode(&n, &ll)
	k := node{2, nil}
	insertNode(&k, &ll)
	printList(&ll)
}
