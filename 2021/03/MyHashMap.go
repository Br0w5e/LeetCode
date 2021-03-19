package main

//706. 设计哈希映射
type MyHashMap struct {
	Map map[int]int
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{make(map[int]int )}
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
	this.Map[key] = value
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if val, ok := this.Map[key]; ok {
		return val
	}
	return -1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
	delete(this.Map, key)
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
