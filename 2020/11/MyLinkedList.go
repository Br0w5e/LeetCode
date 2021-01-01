package main

//707. 设计链表
//链表的增删改查
type Node struct{
	Next *Node
	Val int
}

type MyLinkedList struct {
	Head *Node
	Tail *Node
	Length int
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	head := &Node{}
	tail := head
	return MyLinkedList{Head:head, Tail:tail, Length:0}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Length{
		return -1
	}
	t := this.Head.Next
	for i := 0; i < index; i++{
		t = t.Next
	}
	return t.Val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	this.AddAtIndex(0, val)
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	this.AddAtIndex(this.Length, val)
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index > this.Length{
		index = this.Length
	}
	t := this.Head
	for i := 0; i < index; i++{
		t = t.Next
	}
	n := &Node{Val:val}
	n.Next = t.Next
	t.Next = n
	this.Length++
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index >= this.Length{
		return
	}
	t := this.Head
	for i := 0; i < index; i++{
		t = t.Next
	}
	t.Next = t.Next.Next
	this.Length--
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
