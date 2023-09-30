// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

type node struct {
	data int
	prev *node
	next *node
}
type dlinkedlist struct {
	head *node
	tail *node
}

func (d *dlinkedlist) AddToHead(v int) *node {
	n := &node{data: v}
	if d.head == nil {
		d.head = n
		d.tail = n
	} else {
		n.next = d.head
		d.head.prev = n
		d.head = n
	}
	return n
}

func (d *dlinkedlist) AddToTail(v int) {
	n := &node{data: v}
	if d.tail == nil {
		d.head = n
		d.tail = n
	} else {
		n.prev = d.tail
		d.tail.next = n
		d.tail = n
	}
}

func (d *dlinkedlist) Delete(n *node) {

	if d.head == n {
		d.head = n.next
	}
	if d.tail == n {
		d.tail = n.prev
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
	if n.prev != nil {
		n.prev.next = n.next
	}

	n = nil
}
func (d *dlinkedlist) MoveToFront(n *node, v int) {
	n.data = v
	if n.prev == nil {
		d.head = n
		d.tail = n
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = d.head
	n.prev = nil
	d.head = n
}
func (d *dlinkedlist) PrintList() {
	n := d.head
	if n == nil {
		fmt.Println("list is empty")
	}
	for n != nil {
		fmt.Printf("{ %d } => ", n.data)
		n = n.next
	}
	fmt.Println(" ")
}
func (d *dlinkedlist) RemoveTail() {
	if d.tail == nil {
		return
	}
	n := d.tail
	if n.prev == nil {

		d.tail = nil
		d.head = nil
	}

	d.tail = n.prev
	d.tail.next = nil
	n = nil
}

func main() {
	l := new(dlinkedlist)
	l.AddToHead(1)
	l.AddToHead(2)
	l.AddToHead(2)
	l.AddToHead(2)
	l.AddToHead(3)
	l.PrintList()
	l.AddToHead(4)
	l.PrintList()
	l.AddToHead(5)
	l.PrintList()
	l.AddToHead(5)
	l.PrintList()
	l.AddToTail(6)
	l.PrintList()
}
