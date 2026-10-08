func isAnagram(s string, t string) bool {
	smap1 := make(map[byte]int)
	smap2 := make(map[byte]int)
	
	if len(s) != len(t){
		return false
	}

	n := len(s)

	for i := 0; i < n; i++{
		smap1[s[i]]++
		smap2[t[i]]++
	}

	if len(smap1) != len(smap2){
		return false
	}

	for k, v := range smap1{
		if v != smap2[k]{
			return false
		}
	}
	return true
}
