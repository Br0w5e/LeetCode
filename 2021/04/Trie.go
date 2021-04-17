package main

//208. 实现 Trie (前缀树)

//map去实现
import (
	"fmt"
	"strings"
	"time"
)

type Trie struct {
	m map[string]bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
	m := make(map[string]bool)
	return Trie{
		m,
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	this.m[word] = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return this.m[word] == true
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if this.m[prefix] {
		return true
	}
	for k, _ := range this.m {
		if strings.HasPrefix(k, prefix) {
			return true
		}
	}
	return false
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

//前缀树
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

