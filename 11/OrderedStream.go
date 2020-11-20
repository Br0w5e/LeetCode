package main

import "math"

//5601. 设计有序流
type OrderedStream struct {
	ptr int
	queue []string
}


func Constructor(n int) OrderedStream {
	return OrderedStream{
		ptr: 1,
		queue: make([]string, n+1),
	}
}


func (this *OrderedStream) Insert(id int, value string) []string {
	this.queue[id] = value
	res := make([]string, 0)
	for this.ptr < len(this.queue) && this.queue[this.ptr] != "" {
		res = append(res, this.queue[this.ptr])
		this.ptr++
	}
	return res
}


/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(id,value);
 */
