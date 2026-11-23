func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	total := m + n
	half := (total + 1) / 2

	low, high := 0, m

	for low <= high {
		i := low + (high-low)/2
		j := half - i

		al := math.MinInt
		if i > 0 {
			al = nums1[i-1]
		}
		ar := math.MaxInt
		if i < m {
			ar = nums1[i]
		}

		bl := math.MinInt
		if j > 0 {
			bl = nums2[j-1]
		}
		br := math.MaxInt
		if j < n {
			br = nums2[j]
		}

		if al <= br && ar >= bl {
			if total%2 == 1 {
				return float64(max(al, bl))
			}
			return float64(max(al, bl)+min(ar, br)) / 2.0
		}

		if al > br {
			high = i - 1
		} else {
			low = i + 1
		}
	}

	return 0.0
}