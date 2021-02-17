package _732_find_the_highest_altitude

func largestAltitude(gain []int) int {
	var res, cur = 0, 0
	for _, j := range gain {
		cur += j
		if cur > res { res = cur }
	}
	return res
}
