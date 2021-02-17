package _013_roman_to_integer

func romanToInt(s string) int {

	var ri = make(map[string]int)
	ri["I"] = 1
	ri["V"] = 5
	ri["X"] = 10
	ri["L"] = 50
	ri["C"] = 100
	ri["D"] = 500
	ri["M"] = 1000

	var res int

	for i := 0; i < len(s); {

		si := ri[string(s[i])]

		if i+1 < len(s) && si < ri[string(s[i+1])] {
			res += ri[string(s[i+1])] - si
			i += 2
		} else {
			res += si
			i += 1
		}
	}

	return res
}
