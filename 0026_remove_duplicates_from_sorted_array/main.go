package main

import "fmt"


func main(){

	req := []int{0,0,1,1,1,2,2,3,3,4}

	fmt.Printf("result: %d\n", removeDuplicates(req))
}

func removeDuplicates(nums []int) int {

	var ol = len(nums)

	for i:=0;i<len(nums);{
		for j:=i+1;j<len(nums);{
			if nums[i] == nums[j] {
				nums = append(nums[:j], nums[j+1:]...)
				break
			}
			j++
		}

		if len(nums) < ol {
			ol = len(nums)
		} else {
			i++
		}
	}

	return len(nums)
}