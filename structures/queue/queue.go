package queue

import "sync"

// Queue is a thread-safe FIFO data structure.
type Queue struct {
	queue []any

	mu sync.Mutex
}

// NewQueue returns a new Queue.
func NewQueue() *Queue {
	return &Queue{}
}

// Len returns the number of elements in the queue.
func (q *Queue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.queue)
}

// Push adds an element to the end of the queue.
func (q *Queue) Push(a any) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.queue = append(q.queue, a)
}

// Pop removes and returns the element at the front of the queue.
func (q *Queue) Pop() any {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.queue) == 0 {
		return nil
	}

	tmp := q.queue[0]
	q.queue = q.queue[1:]
	return tmp
}
