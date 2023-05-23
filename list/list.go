package main

import "fmt"

type Node struct {
	val  interface{}
	next *Node
}

type List struct {
	head *Node

	last *Node
}

func (l *List) Insert(val interface{}) {
	node := &Node{val: val}

	if l.head == nil {
		l.head = node
		l.last = node
		return
	}

	current := l.head

	for current.next != nil {
		current = current.next
	}

	current.next = node
}

func (l *List) InsertAt(val interface{}, pos int) {
	node := &Node{val: val}

	if l.head == nil {
		l.head = node
		l.last = node
		return
	}

	current := l.head

	for i := 0; i < pos; i++ {
		current = current.next
	}

	node.next = current.next
	current.next = node
}

func (l *List) Delete(val interface{}) {
	if l.head == nil {
		return
	}

	current := l.head

	for current.next != nil {
		if current.next.val == val {
			current.next = current.next.next
			return
		}

		if current.val == val {
			l.head = l.head.next
			return
		}

		current = current.next
	}
}

func (l *List) Show() {
	current := l.head

	for current != nil {
		fmt.Printf("%d -> ", current.val)
		current = current.next
	}
	println()
}

func main() {
	l := &List{}

	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)
	l.Insert(5)
	l.Insert(6)

	l.Show()

	l.Delete(1)
	l.Show()

	l.Delete(4)
	l.Show()

	l.Delete(6)
	l.Show()
}
