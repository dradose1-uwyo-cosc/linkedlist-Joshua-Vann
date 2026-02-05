//Joshua Vann
//COSC 3750
//2-1-26

package ds

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

// insert at the tail DONE
func (l *LinkedList) Insert(value string) {
	if l.Head == nil {
		l.Head = &Node{data: value, Next: nil}
		l.Tail = l.Head
		l.Size += 1
	} else if l.Size == 1 {
		l.Tail = &Node{data: value, Next: nil}
		l.Head.Next = l.Tail
		l.Size += 1
	} else {
		l.Tail.Next = &Node{data: value, Next: nil}
		l.Tail = l.Tail.Next
		l.Size += 1
	}
}

// inserts at a position, returns true if success or false if not, like if pos doesn't exist DONE
func (l *LinkedList) InsertAt(position int, value string) error {
	if position > l.Size+1 || position < 0 || (l == nil && position != 0) {
		return errors.New("Invalid Position")
	} else if position == 0 {
		l.Size += 1
		if l.Size == 1 {
			l.Head = &Node{data: value, Next: nil}
			l.Tail = l.Head
		} else {
			l.Head = &Node{data: value, Next: l.Head}
		}
	} else if position == l.Size {
		l.Size += 1
		if l.Size == 1 {
			l.Head = &Node{data: value, Next: nil}
			l.Tail = l.Head
		} else {
			l.Tail.Next = &Node{data: value, Next: nil}
			l.Tail = l.Tail.Next
		}
	} else {
		l.Size += 1
		check := l.Head
		for i := 1; i < position; i++ {
			check = check.Next
		}
		check.Next = &Node{data: value, Next: check.Next}
	}
	return nil
}

// remove first occurrence of the value DONE
func (l *LinkedList) Remove(value string) error {
	if l.Size == 0 {
		return errors.New("Unsupported operation on empty list")
	}
	var last *Node
	for i := l.Head; i.data != value; i = i.Next {
		last = i
	}
	t := last.Next.Next
	last.Next.Next = nil
	last.Next = t
	l.Size -= 1
	return nil
}

// remove all occurrences of a value DONE
func (l *LinkedList) RemoveAll(value string) error {
	if l.Size == 0 {
		return errors.New("Unsupported operation on empty list")
	}
	j := 0
	for i := l.Head; i != nil; i = i.Next {
		if i.data == value {
			l.RemoveAt(j)
		} else {
			j++
		}
	}
	return nil
}

// remove at a position, if index exists DONE
func (l *LinkedList) RemoveAt(pos int) error {
	if l.Size == 0 {
		return errors.New("Unsupported operation on empty list")
	}
	if pos > l.Size-1 || pos < 0 {
		return errors.New("Invalid Position")
	} else if pos == 0 {
		l.Size -= 1
		l.Head = l.Head.Next
	} else if pos == l.Size-1 {
		l.Size -= 1
		check := l.Head
		for i := 1; i != pos; i++ {
			check = check.Next
		}
		check.Next = nil
		l.Tail = check
	} else {
		l.Size -= 1
		check := l.Head
		for i := 1; i != pos; i++ {
			check = check.Next
		}
		if pos == l.Size-1 {
			l.Tail = check
		}
		t := check.Next.Next
		check.Next = t
	}
	return nil
}

// checks if the linked list is empty DONE
func (l *LinkedList) IsEmpty() bool {
	return l.Size == 0
}

// get size of ll DONE
func (l *LinkedList) GetSize() int {
	return l.Size
}

// reverse the list DONE
func (l *LinkedList) Reverse() {
	if l.Size < 2 {
		return
	}
	var prev *Node = nil
	var next *Node
	for i := l.Head; i != nil; i = next {
		next = i.Next
		i.Next = prev
		prev = i
	}
	l.Head, l.Tail = l.Tail, l.Head
}

// print the list DONE
func (l *LinkedList) PrintList() {
	for i := l.Head; i != nil; i = i.Next {
		fmt.Println(i.data)
	}
}
