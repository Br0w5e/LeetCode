package main

import "container/list"

//面试题 03.03. 堆盘子
type StackOfPlates struct {
	arr []*list.List
	cap int
}

func Constructor(cap int) StackOfPlates {
	return StackOfPlates{make([]*list.List, 0, 100), cap}
}

func (this *StackOfPlates) Push(val int) {
	if this.cap == 0 {
		return
	}
	last := len(this.arr) - 1
	if last < 0 || this.arr[last].Len() == this.cap {
		this.arr = append(this.arr, list.New())
		last++
	}
	this.arr[last].PushBack(val)
}

func (this *StackOfPlates) Pop() int {
	last := len(this.arr) - 1
	if last < 0 {
		return -1
	}
	l := this.arr[last]
	res := l.Remove(l.Back()).(int)
	if l.Len() == 0 {
		for last--; last >= 0 && this.arr[last] == nil; last-- {
		}
		this.arr = this.arr[:last+1]
	}
	return res
}

func (this *StackOfPlates) PopAt(index int) int {
	if index >= len(this.arr) {
		return -1
	}
	l := this.arr[index]
	if l == nil {
		return -1
	}
	res := l.Remove(l.Back()).(int)
	if l.Len() == 0 {
		this.arr[index] = nil
		for i := index + 1; i < len(this.arr); i++ {
			this.arr[i-1] = this.arr[i]
		}
		this.arr = this.arr[:len(this.arr)-1]
	}
	return res
}


/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */
