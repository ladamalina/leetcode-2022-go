package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "case 1",
			nums: []int{1, 1, 1, 1, 2, 3, 5, 1},
			want: 5,
		},
		{
			name: "case 2",
			nums: []int{5, 2, 7, 9, 16},
			want: 5,
		},
		{
			name: "case 3",
			nums: []int{1000000000, 1000000000, 1000000000},
			want: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, longestSubarray(test.nums))
		})
	}
}
