type pair struct {
	value     string
	timestamp int
}

type TimeMap struct {
	data map[string][]pair
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][]pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.data[key]; !ok{
		this.data[key] = make([]pair, 0)
	}

	this.data[key] = append(this.data[key], pair{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.data[key]; !ok{
		return ""
	}

	data := this.data[key]
	l, r := 0, len(data) -1

	for l <= r {
		m := l + (r-l)/2
		if data[m].timestamp == timestamp {
			return data[m].value
		} else if data[m].timestamp > timestamp{
			r = m - 1
		} else{
			l = m + 1
		}
	}

	if r < 0 {
		return ""
	}

	return data[r].value
}
