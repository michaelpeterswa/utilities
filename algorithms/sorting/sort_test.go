package sorting_test

import (
	"testing"

	"676f.dev/utilities/algorithms/sorting"
	"github.com/stretchr/testify/assert"
)

func TestSortString(t *testing.T) {
	type test struct {
		inputs []string
		want   []string
	}

	tests := []test{
		{inputs: []string{"d", "k", "a", "z", "r", "t"}, want: []string{"a", "d", "k", "r", "t", "z"}},
		{inputs: []string{"hello", "goodbye", "test", "gopher", "zulu", "whiskey"}, want: []string{"goodbye", "gopher", "hello", "test", "whiskey", "zulu"}},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.want, sorting.Sort(tc.inputs))
	}
}

func TestSortInt64(t *testing.T) {
	type test struct {
		inputs []int64
		want   []int64
	}

	tests := []test{
		{inputs: []int64{9, 3, 4, 6, 2, 6, 7}, want: []int64{2, 3, 4, 6, 6, 7, 9}},
		{inputs: []int64{1234592342, 9847345786, 23984678374, 2349863472, 1239238365}, want: []int64{1234592342, 1239238365, 2349863472, 9847345786, 23984678374}},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.want, sorting.Sort(tc.inputs))
	}
}

func TestSortFloat64(t *testing.T) {
	type test struct {
		inputs []float64
		want   []float64
	}

	tests := []test{
		{inputs: []float64{9.4, 3.2, 4.3, 6.5, 2.6, 6.6, 7.9}, want: []float64{2.6, 3.2, 4.3, 6.5, 6.6, 7.9, 9.4}},
		{inputs: []float64{12.34592342, 98.47345786, 23.984678374, 23.49863472, 12.39238365}, want: []float64{12.34592342, 12.39238365, 23.49863472, 23.984678374, 98.47345786}},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.want, sorting.Sort(tc.inputs))
	}
}
