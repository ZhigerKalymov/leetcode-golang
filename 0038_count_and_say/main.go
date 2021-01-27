package _038_count_and_say

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	k := countAndSay(n-1)
	cnt := 1
	var res string

	for i := 1; i < len(k) + 1; i++ {
		if i < len(k) && k[i] == k[i-1] {
			cnt += 1
		} else {
			res += strconv.Itoa(cnt) + string(k[i-1])
			cnt = 1
		}
	}
	return res
}
