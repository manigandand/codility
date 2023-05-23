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

func (l *DList) Delete(val interface{}) {
	if l.head == nil {
		return
	}

	current := l.head

	for current.next != nil {
		if current.val == val && current.prev == nil {
			l.head = current.next
			current.next.prev = nil
			return
		}
		// current is 2
		// current.next is 3
		if current.next.val == val {
			// if delete last node
			if current.next.next == nil {
				current.next = nil
				l.tail = current
				return
			}

			// 2's next is 4 (3's next)
			current.next.next.prev = current.next.prev
			current.next = current.next.next
			return
		}

		current = current.next
	}
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
	fmt.Println("------------------")

	l.Delete(1)
	l.Show()
	l.RShow()
	fmt.Println("------------------")

	l.Delete(4)
	l.Show()
	l.RShow()
	fmt.Println("------------------")

	l.Delete(5)
	l.Show()
	l.RShow()
	fmt.Println("------------------")
}
