package main

type StackImpl struct {
	data []int
	top  int
}

func make_stack(capacity int) Stack {
	this := &StackImpl{}
	this.data = make([]int, capacity, capacity)
	this.top = -1
	return this
}

func (this *StackImpl) empty() bool {
	return this.top == -1
}

func (this *StackImpl) full() bool {
	return this.top == (len(this.data) - 1)
}

func (this *StackImpl) peak() int {
	return this.data[this.top]
}

func (this *StackImpl) push(value int) {
	this.top++
	this.data[this.top] = value
}

func (this *StackImpl) pop() int {
	top := this.peak()
	this.top--
	return top
}
