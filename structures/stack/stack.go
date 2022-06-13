package stack

import "sync"

type Stack struct {
	stack []any

	mu sync.Mutex
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.stack)
}

func (s *Stack) Push(a any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.stack = append(s.stack, a)
}

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
