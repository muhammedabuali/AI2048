package main

import (
	"fmt"
)

type Stack struct {
	data []int
	top  int
}

func make_stack(capacity int) *Stack {
	this := &Stack{}
	this.data = make([]int, capacity, capacity)
	this.top = -1
	return this
}

func (this *Stack) empty() bool {
	return this.top == -1
}

func (this *Stack) full() bool {
	return this.top == (len(this.data) - 1)
}

func (this *Stack) peak() int {
	return this.data[this.top]
}

func (this *Stack) push(value int) {
	this.top++
	this.data[this.top] = value
}

func (this *Stack) pop() int {
	top := this.peak()
	this.data[this.top] = 0
	this.top--
	return top
}

func (this *Stack) display() {
	fmt.Println(this.data)
}
