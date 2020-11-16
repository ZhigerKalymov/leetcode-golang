package main

import "fmt"

func main(){
	fmt.Println(solution([]int{1,2,2,3,3,3,4,5,5}, 1))
}

func solution(data []int, n int) []int {
	var res []int
	var dbl = make(map[int]int)

	for i:=0; i<len(data); i++{
		dbl[data[i]] += 1
	}

	fmt.Println("map: ", dbl)

	for k,v := range dbl {
		if v <= n {
			res = append(res, k)
		}
	}

	return res

}
