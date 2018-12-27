package collections

import "fmt"

type Dequer interface {
	IsEmpty() bool
	AddFirst(interface{})
	AddLast(interface{})
	RemoveFirst() interface{}
	RemoveLast() interface{}
	String() string
}

type Node struct { //will make it a tailq
	item interface{}
	next *Node
	prev *Node
}

type Deque struct {
	head *Node
	tail *Node
}

func (d *Deque) IsEmpty() bool {
	return d.head == nil || d.tail == nil
}

func (d *Deque) AddFirst(item interface{}) {
	oldfirst := d.head
	d.head = new(Node)
	d.head.item = item
	d.head.next = oldfirst
	d.head.prev = nil
	if oldfirst != nil {
		oldfirst.prev = d.head
	} else { // means first element was added fix tail too
		d.tail = d.head
	}
}

func (d *Deque) AddLast(item interface{}) {
	oldlast := d.tail
	d.tail = new(Node)
	d.tail.item = item
	d.tail.prev = oldlast
	d.tail.next = nil
	if oldlast != nil {
		oldlast.next = d.tail
	} else { // means first elem was added through here.
		d.head = d.tail
	}
}

func (d *Deque) RemoveFirst() interface{} {
	if d.IsEmpty() {
		return ""
	}

	oldhead := d.head
	d.head = oldhead.next
	if d.head == nil { //nil the list
		d.tail = nil
	} else {
		d.head.prev = nil
	}
	return oldhead.item
}

func (d *Deque) RemoveLast() interface{} {
	if d.IsEmpty() {
		return ""
	}

	oldlast := d.tail
	d.tail = oldlast.prev
	if d.tail == nil { //nil the list
		d.head = nil
	} else {
		d.tail.next = nil
	}
	return oldlast.item
}

func (d *Deque) String() string {
	ret := ""
	if d.IsEmpty() {
		return "EMPTY"
	} else {
		for elem := d.head; elem != nil; elem = elem.next { //forward walk
			if elem.prev == nil && elem.next == nil {
				ret += fmt.Sprintf("HT(%s)", elem.item)
			} else if elem.prev == nil {
				ret += fmt.Sprintf("H(%s) - ", elem.item)
			} else if elem.next == nil {
				ret += fmt.Sprintf("T(%s)", elem.item)
			} else {
				ret += fmt.Sprintf("%s - ", elem.item)
			}
		}
	}
	return ret
}
