package main

import "fmt"

type Node struct {
	val  interface{}
	prev *Node
	next *Node
}

type DList struct {
	head *Node
	tail *Node
}

func (l *DList) Insert(val interface{}) {
	node := &Node{val: val}

	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	current := l.head

	for current.next != nil {
		current = current.next
	}

	node.prev = current

	current.next = node
	l.tail = node
}

func (l *DList) Show() {
	current := l.head

	for current != nil {
		fmt.Printf("%d -> ", current.val)
		current = current.next
	}
	println()
}

func (l *DList) RShow() {
	current := l.tail

	for current != nil {
		fmt.Printf("%d -> ", current.val)
		current = current.prev
	}
	println()
}

func main() {
	l := &DList{}

	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)
	l.Insert(5)

	l.Show()
	l.RShow()
}
