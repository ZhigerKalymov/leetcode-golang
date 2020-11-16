package main

import "fmt"

func main(){
	fmt.Printf("result - 5: %d\n", removeElement([]int{0,1,2,2,3,0,4,2}, 2))
}

func removeElement(nums []int, val int) int {

	for i:=0; i<len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	return len(nums)
}