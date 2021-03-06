package main

func isPalindrome(x int) bool {
	if x<0 || (x%10==0 && x!=0) {
		return false
	}

	rn := 0
	for ; x>rn; {
		rn = rn * 10 + x % 10
		x /= 10
	}

	return x==rn || x == rn/10
}
