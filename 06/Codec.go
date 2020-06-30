package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {

}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	strs := make([]string, 0)
	this.dfs(root, &strs)
	return strings.Join(strs, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) <= 0 {
		return nil
	}
	vals := strings.Split(data, ",")
	root := TreeNode{}
	index := 0
	this.recoverDfs(vals, &index, &root)
	return &root
}

func (this *Codec) dfs(node *TreeNode, str *[]string) {
	if node == nil {
		*str = append(*str, "nil")
		return
	}
	*str = append(*str, strconv.Itoa(node.Val))

	this.dfs(node.Left, str)
	this.dfs(node.Right, str)
}

func (this *Codec) recoverDfs(vals []string, index *int, node *TreeNode) {
	if vals[*index] == "nil" {
		node = nil
		return
	}
	if vals[*index] != "nil" {
		node.Val, _ = strconv.Atoi(vals[*index])
	}
	*index++
	if *index < len(vals) && vals[*index] != "nil" {
		node.Left = &TreeNode{}
		this.recoverDfs(vals, index, node.Left)
	}
	*index++
	if *index < len(vals) && vals[*index] != "nil" {
		node.Right = &TreeNode{}
		this.recoverDfs(vals, index, node.Right)
	}
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */