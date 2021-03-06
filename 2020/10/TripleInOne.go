package main


type TripleInOne struct {
	stackSize int
	stack     []int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		stackSize: stackSize,
		stack:     make([]int, (stackSize+1)*3),
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	if len(this.stack) < (this.stackSize+1)*stackNum {
		return
	}
	sidx := (this.stackSize + 1) * stackNum
	num := this.stack[sidx]
	if num >= this.stackSize {
		return
	}
	this.stack[sidx+num+1] = value
	this.stack[sidx]++
}

func (this *TripleInOne) Pop(stackNum int) int {
	if len(this.stack) < (this.stackSize+1)*stackNum {
		return -1
	}
	sidx := (this.stackSize + 1) * stackNum
	num := this.stack[sidx]
	if num == 0 {
		return -1
	}
	v := this.stack[sidx+num]
	this.stack[sidx]--
	this.stack[sidx+num] = 0
	return v
}

func (this *TripleInOne) Peek(stackNum int) int {
	if len(this.stack) < (this.stackSize+1)*stackNum {
		return -1
	}
	sidx := (this.stackSize + 1) * stackNum
	num := this.stack[sidx]
	if num == 0 {
		return -1
	}
	return this.stack[sidx+num]
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	if len(this.stack) < (this.stackSize+1)*stackNum {
		return true
	}
	sidx := (this.stackSize + 1) * stackNum
	num := this.stack[sidx]
	if num == 0 {
		return true
	}
	return false
}


/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */

