package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountEffective(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "case 1",
			nums: []int{1, 2, 3},
			want: 3,
		},
		{
			name: "case 2",
			nums: []int{7, 4, 6},
			want: 4,
		},
		{
			name: "case 3",
			nums: []int{8, 8},
			want: 1,
		},
		{
			name: "case 4",
			nums: []int{2, 2, 1},
			want: 5,
		},
		{
			name: "case 26",
			nums: []int{10, 12},
			want: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, countEffective(test.nums))
		})
	}
}
