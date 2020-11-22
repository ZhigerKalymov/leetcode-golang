package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("result: ", solution("4"))
	fmt.Println("result: ", solution("15"))
	fmt.Println("result: ", solution("16"))
	fmt.Println("result: ", solution("36"))
	fmt.Println("result: ", solution("56"))
	fmt.Println("result: ", solution("57"))
	fmt.Println("result: ", solution("59876545678938247234823418297614259876545678938247224070104841062065987654567893824723463462235342526598765456789382472414929650159876545678938247259876545678938247254387842951559876545678938247235235168758962395987654567893824724567893824723482341829761425987654567893824722407010484106206598765484639214302"))
}


func solution(n string) int {

	//TODO big.int

	i, err := strconv.Atoi(n)
	if err != nil {
		return 0
	}
	j := 0

	for ; i > 1; {
		if i & 1 == 1 {
			if i % 4 == 1 || i == 3 {
				i -= 1
			} else {
				i += 1
			}
		} else {
			i = i >> 1
		}
		j += 1
	}

	return j
}