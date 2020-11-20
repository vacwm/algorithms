package collections

import (
	"errors"
)

// Stack supports retrieval by LIFO order.
type Stack struct {
	arr []int
}

// Push inserts `elem` to the top of stack s.
func (s Stack) Push(elem int) {
	s.arr = append(s.arr, elem)
}

// Pop removes the top item of stack s, if applicable.
func (s Stack) Pop() (int, error) {
	if len(s.arr) == 0 {
		return 0, errors.New("empty stack")
	}
	val := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return val, nil
}
