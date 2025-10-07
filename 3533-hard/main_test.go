package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcatenatedDivisibility(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "case 1",
			nums: []int{3, 12, 45},
			k:    5,
			want: []int{3, 12, 45},
		},
		{
			name: "case 2",
			nums: []int{10, 5},
			k:    10,
			want: []int{5, 10},
		},
		{
			name: "case 3",
			nums: []int{1, 2, 3},
			k:    5,
			want: []int{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := concatenatedDivisibility(test.nums, test.k)
			assert.Equal(t, test.want, actual)
		})
	}
}
