package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxPartitionFactor(t *testing.T) {
	tests := []struct {
		name string
		nums [][]int
		want int
	}{
		{
			name: "case 1",
			nums: [][]int{{0, 0}, {0, 2}, {2, 0}, {2, 2}},
			want: 4,
		},
		{
			name: "case 2",
			nums: [][]int{{0, 0}, {0, 1}, {10, 0}},
			want: 11,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, maxPartitionFactor(test.nums))
		})
	}
}
