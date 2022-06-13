package queue

import "sync"

type Queue struct {
	queue []any

	mu sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.queue)
}

func (q *Queue) Push(a any) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.queue = append(q.queue, a)
}

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
