package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBinaryPalindromes(t *testing.T) {
	tests := []struct {
		name string
		n    int64
		want int
	}{
		{
			name: "case 1",
			n:    9,
			want: 6,
		},
		{
			name: "case 2",
			n:    0,
			want: 1,
		},
		{
			name: "case 320",
			n:    8796093022207,
			want: 6291455,
		},
		{
			name: "case 322",
			n:    1099511627776,
			want: 2097151,
		},
		{
			name: "case 341",
			n:    1,
			want: 2,
		},
		{
			name: "case 343",
			n:    17592186044417,
			want: 8388608,
		},
		{
			name: "case 363",
			n:    431829740748048,
			want: 42516273,
		},
		{
			name: "case 10^15",
			n:    1000000000000000,
			want: 63356754,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, countBinaryPalindromes(test.n))
		})
	}
}
