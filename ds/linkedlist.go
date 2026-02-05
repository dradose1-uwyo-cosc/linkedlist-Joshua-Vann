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
		} else {
			l.Head = &Node{data: value, Next: l.Head.Next}
		}
	} else {
		l.Size += 1
		check := l.Head
		for i := 0; i < position; i++ {
			check = check.Next
		}
		check.Next = &Node{data: value, Next: check.Next}
	}
	return nil
}

// remove first occurrence of the value DONE
func (l *LinkedList) Remove(value string) error {
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
	j := 0
	for i := l.Head; i != nil; i = i.Next {
		if i.data == value {
			l.RemoveAt(0)
		}
		j++
	}
	return nil
}

// remove at a position, if index exists DONE
func (l *LinkedList) RemoveAt(pos int) error {
	if pos > l.Size-1 || pos < 0 {
		return errors.New("Invalid Position")
	} else if pos == 0 {
		l.Size -= 1
		l.Head = l.Head.Next
	} else {
		l.Size -= 1
		check := l.Head
		for i := 1; i != pos; i++ {
			check = check.Next
		}
		if pos == l.Size-1 {
			l.Tail = check
		}
		last := check.Next.Next
		t := check.Next.Next
		check.Next = t
		last.Next = nil
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
	var prev *Node = nil
	for i := l.Head; i != nil; i = i.Next {
		i.Next = prev
		prev = i
	}
}

// print the list DONE
func (l *LinkedList) PrintList() {
	for i := l.Head; i != nil; i = i.Next {
		fmt.Println(i.data)
	}
}
