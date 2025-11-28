package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumMusicPlaylists(t *testing.T) {
	tests := []struct {
		name string
		n    int
		g    int
		k    int
		want int
	}{
		{
			name: "case 1",
			n:    3,
			g:    3,
			k:    1,
			want: 6,
		},
		{
			name: "case 2",
			n:    2,
			g:    3,
			k:    0,
			want: 6,
		},
		{
			name: "case 2",
			n:    2,
			g:    3,
			k:    1,
			want: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, numMusicPlaylists(test.n, test.g, test.k))
		})
	}
}
