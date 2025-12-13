package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTilingRectangle(t *testing.T) {
	tests := []struct {
		name string
		n    int
		m    int
		want int
	}{
		{
			name: "case 1",
			n:    2,
			m:    3,
			want: 3,
		},
		{
			name: "case 2",
			n:    5,
			m:    8,
			want: 5,
		},
		{
			name: "case 3",
			n:    11,
			m:    13,
			want: 6,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tilingRectangle(test.n, test.m), test.want)
		})
	}
}
