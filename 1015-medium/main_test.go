package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestRepunitDivByK(t *testing.T) {
	tests := []struct {
		name string
		k    int
		want int
	}{
		{
			name: "case 1",
			k:    1,
			want: 1,
		},
		{
			name: "case 2",
			k:    2,
			want: -1,
		},
		{
			name: "case 3",
			k:    3,
			want: 3,
		},
		{
			name: "case 10^5",
			k:    100000,
			want: -1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, smallestRepunitDivByK(test.k))
		})
	}
}
