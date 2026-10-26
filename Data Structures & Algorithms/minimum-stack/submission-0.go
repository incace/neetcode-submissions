type MinStack struct {
    mainStack []int
    minStack  []int
    n         int
}

func Constructor() MinStack {
    return MinStack{
        mainStack: make([]int, 0),
        minStack:  make([]int, 0),
        n:         0,
    }
}

func (this *MinStack) Push(val int) {
	if this.n == 0 {
		this.mainStack = append(this.mainStack, val)
		this.minStack = append(this.minStack, val)
		this.n++
		return
	}

	this.mainStack = append(this.mainStack, val)

	if this.minStack[this.n-1] > val {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, this.minStack[this.n-1])
	}
	this.n++
}

func (this *MinStack) Pop() {
	this.mainStack = this.mainStack[:this.n-1]
	this.minStack = this.minStack[:this.n-1]
	this.n--
}

func (this *MinStack) Top() int {
	return this.mainStack[this.n-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[this.n-1]
}
