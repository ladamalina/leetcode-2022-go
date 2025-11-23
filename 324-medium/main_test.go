package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWiggleSort(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "case 1",
			nums: []int{1, 5, 1, 1, 6, 4},
		},
		{
			name: "case 2",
			nums: []int{1, 3, 2, 2, 3, 1},
		},
		{
			name: "case 3",
			nums: []int{4, 5, 5, 6},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wiggleSort(test.nums)
			for i := 1; i < len(test.nums); i++ {
				if i%2 == 0 {
					assert.Greater(t, test.nums[i-1], test.nums[i])
				} else {
					assert.Less(t, test.nums[i-1], test.nums[i])
				}
			}
		})
	}
}
