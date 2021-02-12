package main

//1381. 设计一个支持增量操作的栈

//设置个size进行记录最大值
type CustomStack struct {
	stack []int
	size int
}


func Constructor(maxSize int) CustomStack {
	return CustomStack{
		make([]int, 0),
		maxSize,
	}
}


func (this *CustomStack) Push(x int)  {
	if len(this.stack) < this.size {
		this.stack = append(this.stack, x)
	}
}


func (this *CustomStack) Pop() int {
	if len(this.stack) > 0 {
		res := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		return res
	}
	return -1
}


func (this *CustomStack) Increment(k int, val int)  {
	if k > len(this.stack) {
		k = len(this.stack)
	}
	for i := 0; i < k; i++ {
		this.stack[i] += val
	}
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
