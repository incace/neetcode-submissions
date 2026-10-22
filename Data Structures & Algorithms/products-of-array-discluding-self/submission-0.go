func productExceptSelf(nums []int) []int {
	var res []int;
	pref, postfix := 1,1
	n := len(nums)

	for i := 0 ; i < n ; i++ {
		res = append(res, pref)
		pref *= nums[i]
	}


	for i := n-1 ; i >= 0 ; i-- {
		res[i] = res[i] * postfix
		postfix *= nums[i]
	}

	return res

}
