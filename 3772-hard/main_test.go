package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubgraphScore(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		good  []int
		want  []int
	}{
		{
			name:  "case 1",
			n:     3,
			edges: [][]int{{0, 1}, {1, 2}},
			good:  []int{1, 0, 1},
			want:  []int{1, 1, 1},
		},
		{
			name:  "case 2",
			n:     5,
			edges: [][]int{{1, 0}, {1, 2}, {1, 3}, {3, 4}},
			good:  []int{0, 1, 0, 1, 1},
			want:  []int{2, 3, 2, 3, 3},
		},
		{
			name:  "case 3",
			n:     2,
			edges: [][]int{{0, 1}},
			good:  []int{0, 0},
			want:  []int{-1, -1},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.want, maxSubgraphScore(test.n, test.edges, test.good))
		})
	}
}
