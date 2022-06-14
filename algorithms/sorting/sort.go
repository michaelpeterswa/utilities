package sorting

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// Sort() wraps the sort.Slice() function and uses the generic type constraints.Ordered.
// If true, the bool isReverse parameter can reverse the sort order (descending).
// If false, the sort order is ascending.
//
// See constraints.Ordered for more information about usable types.
func Sort[C constraints.Ordered](c []C, isReverse bool) []C {
	if !isReverse {
		sort.Slice(c, func(i, j int) bool {
			return c[i] < c[j]
		})
	} else {
		sort.Slice(c, func(i, j int) bool {
			return c[i] > c[j]
		})
	}
	return c
}
