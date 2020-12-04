package main

type FrontMiddleBackQueue struct {
	queue []int
	size int

}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue {
		make([]int, 0),
		0,
	}
}

func (this *FrontMiddleBackQueue) PushFront(val int)  {
	tmp := make([]int, this.size+1)
	copy(tmp[1:], this.queue)
	tmp[0] = val
	this.queue = tmp
	this.size++
}


func (this *FrontMiddleBackQueue) PushMiddle(val int)  {
	tmp := make([]int, this.size+1)
	idx := this.size / 2
	copy(tmp[:idx], this.queue[:idx])
	tmp[idx] = val
	copy(tmp[idx+1:], this.queue[idx:])
	this.queue = tmp
	this.size++
}


func (this *FrontMiddleBackQueue) PushBack(val int)  {
	this.queue = append(this.queue, val)
	this.size++
}


func (this *FrontMiddleBackQueue) PopFront() int {
	if this.size == 0 {
		return -1
	}
	res := this.queue[0]
	this.queue = this.queue[1:]
	this.size--
	return res
}


func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.size == 0 {
		return -1
	}
	mid := (this.size-1)/2
	res := this.queue[mid]
	tmp := make([]int, len(this.queue)-1)
	copy(tmp[:mid], this.queue[:mid])
	copy(tmp[mid:], this.queue[mid+1:])
	this.queue = tmp
	this.size--
	return res
}


func (this *FrontMiddleBackQueue) PopBack() int {
	if this.size == 0 {
		return -1
	}
	res := this.queue[this.size-1]
	this.queue = this.queue[:this.size-1]
	this.size--
	return res
}


/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
