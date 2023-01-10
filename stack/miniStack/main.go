package main

import "math"

type MinStack struct {
	stack    []int
	minStack []int
}

// 初始化
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt32},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.minStack[len(this.minStack)-1]         // 获取辅助栈最新的值
	this.minStack = append(this.minStack, min(x, top)) // 辅助栈中，最上的永远最小
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]          // 弹出最上的
	this.minStack = this.minStack[:len(this.minStack)-1] // 也弹出最上的
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
