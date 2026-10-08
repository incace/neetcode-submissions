func twoSum(nums []int, target int) []int {
    ex_num := make(map[int]int)

	for ind, num := range nums{
		if _, f := ex_num[target - num]; f{
			return []int{ex_num[target - num], ind}
		}
		ex_num[num] = ind
	}
	return nil
}
