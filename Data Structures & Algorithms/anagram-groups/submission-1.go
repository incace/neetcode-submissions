func groupAnagrams(strs []string) [][]string {
	fanswer := make(map[[26]int][]string)
	var result [][]string

	for _, str := range strs{
		var nums [26]int
		for _, bytte := range str{
			nums[bytte - 'a']++
		}
		fanswer[nums] = append(fanswer[nums], str)
	}

	for _, val := range fanswer{
		result = append(result, val)
	}

	return result

}

