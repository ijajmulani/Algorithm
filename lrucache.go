// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

type node struct {
	data int
	key  int
	prev *node
	next *node
}
type dlinkedlist struct {
	len    int
	head   *node
	tail   *node
	maxCap int
}
type lrucache struct {
	hash map[int]*node
	list *dlinkedlist
	cap  int
}

func InitializeLRUCache(cap int) *lrucache {
	return &lrucache{
		hash: make(map[int]*node, cap),
		list: &dlinkedlist{maxCap: cap},
		cap:  cap,
	}
}

func (l *lrucache) Print() {
	l.list.PrintList()
	fmt.Printf("%+v", l.hash)
	fmt.Println(" \n \n ")
}

func (l *lrucache) Add(k, v int) bool {
	if nodeaddress, ok := l.hash[k]; ok {
		l.list.MoveToFront(nodeaddress, v)
		return true
	}
	if l.list.len >= l.cap {
		removedKey := l.list.RemoveTail()
		if removedKey != -1 {
			delete(l.hash, removedKey)
		}
	}
	add := l.list.Add(k, v)
	l.hash[k] = add
	return true
}

func (l *lrucache) Delete(k int) {
	if nodeaddress, ok := l.hash[k]; ok {
		delete(l.hash, k)
		l.list.Delete(nodeaddress)
	}
}

func (d *dlinkedlist) Add(k int, v int) *node {
	n := &node{key: k, data: v}
	if d.head == nil {
		d.head = n
		d.tail = n
	} else {
		n.next = d.head
		d.head.prev = n
		d.head = n
	}
	d.len++
	return n
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
		fmt.Printf("{ %d , %d } => ", n.key, n.data)
		n = n.next
	}
	fmt.Println(" ")
}
func (d *dlinkedlist) RemoveTail() int {
	if d.tail == nil {
		return -1
	}
	n := d.tail
	k := n.key
	if n.prev == nil {

		d.tail = nil
		d.head = nil
		d.len = 0
		return k
	}

	d.tail = n.prev
	d.tail.next = nil
	n = nil
	d.len--
	return k
}

func main() {
	l := InitializeLRUCache(3)
	l.Add(1, 1)
	l.Add(2, 2)
	l.Add(3, 2)
	l.Add(4, 2)
	l.Add(4, 3)
	l.Print()
	l.Add(4, 4)
	l.Print()
	l.Add(4, 5)
	l.Print()
	l.Add(3, 5)
	l.Print()
}
