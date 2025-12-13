package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZigZagArrays(t *testing.T) {
	tests := []struct {
		name    string
		n, l, r int
		want    int
	}{
		{
			name: "case 1",
			n:    3,
			l:    4,
			r:    5,
			want: 2,
		},
		{
			name: "case 2",
			n:    3,
			l:    1,
			r:    3,
			want: 10,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.want, zigZagArrays(test.n, test.l, test.r))
		})
	}
}
