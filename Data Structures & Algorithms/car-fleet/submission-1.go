type Pair struct{
	position int
	speed int
}

func carFleet(target int, position []int, speed []int) int {
	pairs := make([]Pair, len(position))
	time := make([]float64, len(position))
	stack := make([]float64, len(position))
	tos := 0

	for i := 0; i < len(position); i++ {
		pairs[i] = Pair{position[i], speed[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
    return pairs[i].position > pairs[j].position
	})

	for i := 0; i < len(position); i++ {
		time[i] = float64(target - pairs[i].position) / float64(pairs[i].speed)
	}

	stack[0] = time[0]

	for i := 1; i < len(position); i++{
		if time[i] <= stack[tos]{
			continue
		} else{
			tos++
			stack[tos] = time[i]
		}
	}

	return tos + 1
}
