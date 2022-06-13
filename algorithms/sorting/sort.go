package sorting

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Sort[C constraints.Ordered](c []C) []C {
	sort.Slice(c, func(i, j int) bool {
		return c[i] < c[j]
	})

	return c
}
