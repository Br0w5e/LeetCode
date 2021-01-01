package main

//705. 设计哈希集合
// 不要使用内建哈希表
type MyHashSet struct {
	bitset []uint64
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{bitset: []uint64{}}
}


func (s *MyHashSet) Add(key int)  {
	bit := key % 64
	length := key /64
	for i := len(s.bitset); i <= length; i++ {
		s.bitset = append(s.bitset, 0)
	}
	s.bitset[length] = s.bitset[length] | (1 << bit)
}


func (s *MyHashSet) Remove(key int)  {
	bit := key % 64
	length := key /64
	if length >= len(s.bitset) {
		return
	}
	s.bitset[length] = s.bitset[length] & ^(1 << bit)
}


/** Returns true if this set contains the specified element */
func (s *MyHashSet) Contains(key int) bool {
	bit := key % 64
	length := key /64
	if length >= len(s.bitset) {
		return false
	}
	return s.bitset[length] & (1 << bit) != 0
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
