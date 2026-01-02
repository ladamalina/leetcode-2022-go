package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCollisions(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "case 1",
			s:    "RLRSLL",
			want: 5,
		},
		{
			name: "case 2",
			s:    "LLRR",
			want: 0,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.want, countCollisions(test.s))
		})
	}
}
