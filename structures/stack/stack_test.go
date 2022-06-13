package stack_test

import (
	"testing"

	"676f.dev/utilities/structures/stack"
	"github.com/stretchr/testify/assert"
)

func TestStackPushAndPop(t *testing.T) {
	type test struct {
		inputs []string
		want   any
	}

	tests := []test{
		{inputs: []string{}, want: nil},
		{inputs: []string{"a"}, want: "a"},
		{inputs: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}, want: "j"},
	}

	for _, tc := range tests {
		s := stack.NewStack()
		for _, v := range tc.inputs {
			s.Push(v)
		}
		assert.Equal(t, tc.want, s.Pop())
	}
}

func TestStackLen(t *testing.T) {
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
		s := stack.NewStack()
		for _, v := range tc.inputs {
			s.Push(v)
		}
		assert.Equal(t, tc.want, s.Len())
	}
}
