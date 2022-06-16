package shortid_test

import (
	"fmt"
	"testing"

	"676f.dev/utilities/tools/shortid"
	"github.com/stretchr/testify/assert"
)

func TestShortIDLength(t *testing.T) {
	type test struct {
		want int
	}

	tests := []test{
		{want: 6},
		{want: 12},
		{want: 18},
		{want: 24},
	}

	shortid := shortid.NewShortID(shortid.Base58CharacterSet)

	for _, tc := range tests {
		str, err := shortid.Generate(tc.want)
		if err != nil {
			t.Fail()
		}
		fmt.Println(str)
		assert.Equal(t, tc.want, len(str))
	}
}
