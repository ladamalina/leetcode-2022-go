package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinInversionCount(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int64
	}{
		{
			name: "case 1",
			nums: []int{3, 1, 2, 5, 4},
			k:    3,
			want: 0,
		},
		{
			name: "case 2",
			nums: []int{5, 3, 2, 1},
			k:    4,
			want: 6,
		},
		{
			name: "case 3",
			nums: []int{2, 1},
			k:    1,
			want: 0,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.want, minInversionCount(test.nums, test.k))
		})
	}
}
