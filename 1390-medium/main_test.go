package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumFourDivisors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "case 1",
			nums: []int{21, 4, 7},
			want: 32,
		},
		{
			name: "case 2",
			nums: []int{21, 21},
			want: 64,
		},
		{
			name: "case 3",
			nums: []int{1, 2, 3, 4, 5},
			want: 0,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.want, sumFourDivisors(test.nums))
		})
	}
}
