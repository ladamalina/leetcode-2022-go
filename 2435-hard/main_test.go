package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfPaths(t *testing.T) {
	tests := []struct {
		name string
		g    [][]int
		k    int
		want int
	}{
		{
			name: "case 1",
			g:    [][]int{{5, 2, 4}, {3, 0, 5}, {0, 7, 2}},
			k:    3,
			want: 2,
		},
		{
			name: "case 2",
			g:    [][]int{{0, 0}},
			k:    5,
			want: 1,
		},
		{
			name: "case 3",
			g:    [][]int{{7, 3, 4, 9}, {2, 3, 6, 2}, {2, 3, 7, 0}},
			k:    1,
			want: 10,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.want, numberOfPaths(test.g, test.k))
		})
	}
}
