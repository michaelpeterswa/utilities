package stack

import "sync"

// Stack is a thread-safe LIFO data structure.
type Stack struct {
	stack []any

	mu sync.Mutex
}

// NewStack returns a new Stack.
func NewStack() *Stack {
	return &Stack{}
}

// Len returns the number of elements in the stack.
func (s *Stack) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.stack)
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(a any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.stack = append(s.stack, a)
}

// Pop removes and returns the element at the top of the stack.
func (s *Stack) Pop() any {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.stack) == 0 {
		return nil
	}

	tmp := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return tmp
}
