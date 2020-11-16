package main

import (
	"math"
)

func reverse(x int) int {
	var res int

	for ; x != 0; x /= 10  {
		res *= 10
		res += x - ((x / 10) * 10)
	}

	if math.MaxInt32 < res || res < math.MinInt32 {
		return 0
	}

	return res
}