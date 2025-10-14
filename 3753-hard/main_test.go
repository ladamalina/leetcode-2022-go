package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWaviness(t *testing.T) {
	tests := []struct {
		name string
		num  int64
		want int64
	}{
		{
			name: "case 1",
			num:  63,
			want: 0,
		},
		{
			name: "case 2",
			num:  100,
			want: 0,
		},
		{
			name: "case 3",
			num:  101,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, waviness(test.num))
		})
	}
}

func TestWavinessLte(t *testing.T) {
	tests := []struct {
		name string
		num  int64
		want int64
	}{
		{
			name: "case 1",
			num:  63,
			want: 0,
		},
		{
			name: "case 2",
			num:  100,
			want: 0,
		},
		{
			name: "case 3",
			num:  101,
			want: 1,
		},
		{
			name: "case 4",
			num:  120,
			want: 10,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, wavinessLte(test.num))
		})
	}
}

func TestTotalWaviness(t *testing.T) {
	tests := []struct {
		name string
		num1 int64
		num2 int64
		want int64
	}{
		{
			name: "case 1",
			num1: 120,
			num2: 130,
			want: 3,
		},
		{
			name: "case 2",
			num1: 198,
			num2: 202,
			want: 3,
		},
		{
			name: "case 3",
			num1: 4848,
			num2: 4848,
			want: 2,
		},
		{
			name: "case 4",
			num1: 101,
			num2: 120,
			want: 10,
		},
		{
			name: "case 5",
			num1: 107,
			num2: 120,
			want: 4,
		},
		{
			name: "case 566",
			num1: 63,
			num2: 101,
			want: 1,
		},
		{
			name: "case 673",
			num1: 8900,
			num2: 9532,
			want: 794,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, totalWaviness(test.num1, test.num2))
		})
	}
}
