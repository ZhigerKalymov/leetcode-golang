package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main(){
	fmt.Println("result: ", solution("4"))
	fmt.Println("result: ", solution("15"))
	fmt.Println("result: ", solution("16"))
	fmt.Println("result: ", solution("32"))
	fmt.Println("result: ", solution("33"))
	fmt.Println("result: ", solution("31"))
	fmt.Println("result: ", solution("12448057941136394342297748548545082997815840357634948550739612798732309232806852458769500556143453462345234543634645635342536228376971224578111829761428033242424070104841062064840113262840137456279456255812346346223534252646680414929650102954654149991876543878429515708804712300982523523516875896239934729753"))
}


func solution(n string) int {

	b := new(big.Int)
	b, ok := b.SetString(n, 10)
	if !ok {
		return 0
	}

	one := big.NewInt(1)

	for ; b > one ; {

	}



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