//Joshua Vann
//COSC 3750
//2-1-26

package ds

type Stack struct {
	list LinkedList
}

// add new head node
func (s *Stack) Push(value string) {
	s.list.InsertAt(0, value)
}

// remove the head
func (s *Stack) Pop() (string, bool) {
	if s.list.Size == 0 {
		return "", false
	}
	r := s.list.Head.data
	s.list.RemoveAt(0)
	return r, true
}
