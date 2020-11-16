package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(0))
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(52))
	fmt.Println(convertToTitle(701))
	fmt.Println(convertToTitle(17698))
	fmt.Println(convertToTitle(11111))
}

func convertToTitle(n int) string {
	res := ""
	for n > 0 {
		res = string('A'+(n-1)%26) + res
		n = (n - 1) / 26
	}
	return res
}
