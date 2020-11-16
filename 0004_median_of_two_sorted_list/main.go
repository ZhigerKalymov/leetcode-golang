package _004_median_of_two_sorted_list

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	s := append(nums1, nums2)

	var sum int = 0

	for n := range s {
		sum += n
	}

	return sum/len(s)
}
