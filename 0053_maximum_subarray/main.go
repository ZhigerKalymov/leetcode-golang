package main

import "fmt"

func main(){
	fmt.Println("RESULT: ", maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}

func maxSubArray(nums []int) int {
	var n, nm int

	for i,j := range nums {

		n += j

		if j > n {
			n = j
		}

		if n > nm || i == 0{
			nm = n
		}
	}

	return nm
}
