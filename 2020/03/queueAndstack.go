//两个栈实现队列
type CQueue struct {
	in []int
	out []int
    
}


func Constructor() CQueue {
    return CQueue {
		make([]int, 0, 5),
		make([]int, 0, 5),
	}
}


func (this *CQueue) AppendTail(value int)  {
    if value < 1 && value > 10000 {
		return
	} else {
		this.in = append(this.in[:], value)
	}
}


func (this *CQueue) DeleteHead() int {
    //无人入栈
	if len(this.in) == 0 {
		return -1
	}
	// 出入相等
	if len(this.in) == len(this.out) {
		return -1
	}
	//“out栈的长度”刚好等于“要输出的数在in栈中的位置”
	out := this.in[len(this.out)]
	this.out = append(this.out[:], this.in[len(this.out)])

    return out
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */












//队列实现栈
type MyStack struct {
    list []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{list:make([]int,0)}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.list = append(this.list, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    val := this.Top()
    this.list = this.list[:len(this.list)-1]
    return val
}


/** Get the top element. */
func (this *MyStack) Top() int {
    return this.list[len(this.list)-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return len(this.list)==0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */