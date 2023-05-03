package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func (l *linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d", toPrint.data)
		toPrint = toPrint.next
		l.length--
		fmt.Printf("\n")
	}
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: 98}
	node2 := &node{data: 65}
	node3 := &node{data: 9}
	node4 := &node{data: 92}
	node5 := &node{data: 5}
	node6 := &node{data: 29}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.deleteWithValue(29)
	mylist.printListData()
}
