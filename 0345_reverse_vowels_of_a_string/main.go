package main

import (
	"fmt"
)

func main() {

	e1 := "hello"
	e2 := "leetcode"


	fmt.Println("result hello: ", reverseVowels(e1))
	fmt.Println("result leetcode: ", reverseVowels(e2))
}

func reverseVowels(s string) string {
	res := []byte(s)
	for i, j := 0, len(s)-1; i<len(s) && j>i ; i++ {
		if vow(s[i]) {
			for j>=i && !vow(s[j]) {
				j--
			}
			res[i], res[j] = res[j], res[i]
			j--
		}
	}
	return string(res)
}

func vow(c byte) bool {
	switch c {
		case 'a', 'e', 'i', 'o' , 'u', 'A', 'E', 'I', 'O', 'U' :
			return true
		default:
			return false
	}
}
