//Joshua Vann
//COSC 3750
//2-1-26

package ds

import "errors"

type Queue struct {
	list LinkedList
}

// add tail node
func (q *Queue) Push(value string) {
	q.list.InsertAt(q.list.GetSize()-1, value)
}

// remove the head
func (q *Queue) Pop() (string, error) {
	if q.list.Size == 0 {
		return "", errors.New("Cannot perform on empty Queue")
	}
	r := q.list.Head.data
	q.list.RemoveAt(0)
	return r, nil
}
