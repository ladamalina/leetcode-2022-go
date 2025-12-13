package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int64
	}{
		{
			name: "case 1",
			nums: []int{1, 2},
			k:    1,
			want: 3,
		},
		{
			name: "case 2",
			nums: []int{-1, -2, -3, -4, -5},
			k:    4,
			want: -10,
		},
		{
			name: "case 3",
			nums: []int{-5, 1, 2, -3, 4},
			k:    2,
			want: 4,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, maxSubarraySum(test.nums, test.k), test.want)
		})
	}
}
