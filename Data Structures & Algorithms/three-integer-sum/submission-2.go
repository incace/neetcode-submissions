func threeSum(nums []int) [][]int {
	var res [][]int
	hash := make(map[int]int)
	n := len(nums)
	sort.Ints(nums)

	for i, num := range nums {
		hash[num] = i
	}

	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			target := 0 - nums[i] - nums[j]
			if target < nums[j] {
				break
			}

			if val, ok := hash[target]; ok && val > j {
				res = append(res, []int{nums[i], nums[j], target})
			}
		}
	}
	return res
}
