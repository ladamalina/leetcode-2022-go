package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSumDivThree(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "case 1",
			nums: []int{3, 6, 5, 1, 8},
			want: 18,
		},
		{
			name: "case 2",
			nums: []int{4},
			want: 0,
		},
		{
			name: "case 3",
			nums: []int{1, 2, 3, 4, 4},
			want: 12,
		},
		{
			name: "case 42",
			nums: []int{4, 1, 5, 3, 1},
			want: 12,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, maxSumDivThree(test.nums), test.want)
		})
	}
}
