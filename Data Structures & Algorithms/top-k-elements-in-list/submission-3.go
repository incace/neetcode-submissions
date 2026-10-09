func topKFrequent(nums []int, k int) []int {
	n := len(nums)
	freq_cnt := make(map[int]int)
	bucket_arr := make([][]int, n+1)
	var res []int

	for _, num := range nums{
		freq_cnt[num]++
	}

	for key, val := range freq_cnt{
		bucket_arr[val] = append(bucket_arr[val], key)
	}

	for i:=n; i > 0 && k > 0; i--{
		if len(bucket_arr[i]) <= k{
			k -= len(bucket_arr[i])
			res = append(res, bucket_arr[i]...)
		} else {
			res = append(res, bucket_arr[i][:k]...)
			k = 0
		}
	}
	return res
}
