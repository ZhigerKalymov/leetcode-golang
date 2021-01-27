package main

import (
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test_123",
			args{123},
			321,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkReverse(input int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse(input)
	}
}

func BenchmarkReverse123(b *testing.B)         { benchmarkReverse(123, b) }
func BenchmarkReverseBig(b *testing.B) { benchmarkReverse(2147483648, b) }
func BenchmarkReverseNegativeOne(b *testing.B)   { benchmarkReverse(-1, b) }