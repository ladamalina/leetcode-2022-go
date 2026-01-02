package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBestClosingTime(t *testing.T) {
	tests := []struct {
		name      string
		customers string
		want      int
	}{
		{
			name:      "case 1",
			customers: "YYNY",
			want:      2,
		},
		{
			name:      "case 2",
			customers: "NNNNN",
			want:      0,
		},
		{
			name:      "case 3",
			customers: "YYYY",
			want:      4,
		},
		{
			name:      "case 31",
			customers: "NYNNNYYN",
			want:      0,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, bestClosingTime(test.customers), test.want)
		})
	}
}
