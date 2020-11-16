package main

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for k := i+1; k <= len(nums)-1; k++ {
			if v + nums[k] == target {
				return []int{i,k}
			}
		}
	}
	return []int{}
}
