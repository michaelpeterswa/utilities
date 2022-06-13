package queue_test

import (
	"testing"

	"676f.dev/utilities/structures/queue"
	"github.com/stretchr/testify/assert"
)

func TestQueuePushAndPop(t *testing.T) {
	type test struct {
		inputs []string
		want   any
	}

	tests := []test{
		{inputs: []string{}, want: nil},
		{inputs: []string{"a"}, want: "a"},
		{inputs: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}, want: "a"},
	}

	for _, tc := range tests {
		q := queue.NewQueue()
		for _, v := range tc.inputs {
			q.Push(v)
		}
		assert.Equal(t, tc.want, q.Pop())
	}
}

func TestQueueLen(t *testing.T) {
	type test struct {
		inputs []string
		want   int
	}

	tests := []test{
		{inputs: []string{}, want: 0},
		{inputs: []string{"a"}, want: 1},
		{inputs: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}, want: 10},
	}

	for _, tc := range tests {
		q := queue.NewQueue()
		for _, v := range tc.inputs {
			q.Push(v)
		}
		assert.Equal(t, tc.want, q.Len())
	}
}
