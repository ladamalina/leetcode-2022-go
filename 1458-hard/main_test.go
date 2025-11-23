package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDotProduct(t *testing.T) {
	tests := []struct {
		name         string
		nums1, nums2 []int
		want         int
	}{
		{
			name:  "case 1",
			nums1: []int{2, 1, -2, 5},
			nums2: []int{3, 0, -6},
			want:  18,
		},
		{
			name:  "case 2",
			nums1: []int{3, -2},
			nums2: []int{2, -6, 7},
			want:  21,
		},
		{
			name:  "case 3",
			nums1: []int{-1, -1},
			nums2: []int{1, 1},
			want:  -1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := maxDotProduct(test.nums1, test.nums2)
			assert.Equal(t, test.want, res)
		})
	}
}
