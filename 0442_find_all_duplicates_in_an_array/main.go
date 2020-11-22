package main

import (
	"math"
)

func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for _,j := range nums {
		j = int(math.Abs(float64(j)))

		if nums[j-1] > 0 {
			nums[j-1] = nums[j-1] * -1
		} else {
			res = append(res, j)
		}
	}
	return res
}