type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var res string
	for _, st := range strs {
		res += strconv.Itoa(len(st)) + "#" + st
	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
	var res []string
	i := 0
	for i < len(encoded) {
		hashIdx := strings.Index(encoded[i:], "#")
		length, _ := strconv.Atoi(encoded[i : i+hashIdx])

		start := i + hashIdx + 1
		end := start + length

		res = append(res, encoded[start:end])
		i = end
	}
	return res
}
