package queue

import "sync"

type Queue struct {
	queue []any

	mu sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(a any) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.queue = append(q.queue, a)
}
