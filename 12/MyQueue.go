package main
//232. 用栈实现队列


//用栈实现队列
type MyQueue struct {
	queue []int
}
//初始化
func Constructor() MyQueue {
    return MyQueue{queue:[]int{}} 
}

//加入队尾
func (this *MyQueue) Push(x int)  {
    this.queue = append(this.queue, x)
}

//从队头移除元素
func (this *MyQueue) Pop() int {
    res := this.queue[0]
    this.queue = this.queue[1:len(this.queue)]
    return res 
}

//获取队列的第一个元素
func (this *MyQueue) Peek() int {
    return this.queue[0]
}

//判断队列是否为空
func (this *MyQueue) Empty() bool {
    return len(this.queue) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */