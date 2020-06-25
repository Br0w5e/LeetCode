package main

type Node struct {
	key int
	value int
	pre *Node
	next *Node
}

type LRUCache struct {
	cap int
	header *Node
	tail *Node
	m map[int]*Node
}

func Constructor1(capacity int) LRUCache {
	//设置一个哑结点
	cache := LRUCache{
		cap: capacity,
		header: &Node{},
		tail: &Node{},
		m : make(map[int]*Node, capacity),
	}
	cache.header.next = cache.tail
	cache.tail.pre = cache.header
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; !ok {
		//不存在直接返回
		return -1
	} else {
		//存在则移动到头部，并返回值
		this.remove(node)
		this.putHead(node)
		return node.value
	}

}

func (this *LRUCache) Put(key int, value int)  {
	if node, ok := this.m[key]; ok {
		//节点存在，更新数据，并将节点移动到头部
		node.value = value
		this.remove(node)
		this.putHead(node)
	} else {
		//节点不存在,已满则移除链表尾部，并加入头部，未满则直接加入头部
		if len(this.m) == this.cap {
			deletekey := this.tail.pre.key
			this.remove(this.tail.pre)
			delete(this.m, deletekey)
		}
		newNode := Node{
			key: key,
			value: value,
		}
		this.putHead(&newNode)
		this.m[key] = &newNode
	}
}

func (this *LRUCache) remove(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) putHead(node *Node) {
	originNext := this.header.next
	this.header.next = node
	node.next = originNext

	originNext.pre = node
	node.pre = this.header
}