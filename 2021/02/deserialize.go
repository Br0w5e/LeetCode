package main

//385. 迷你语法分析器
import "strconv"

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
	if s[0] != '['{
		ni:=NestedInteger{}
		num, _ := strconv.Atoi(s)
		ni.SetInteger(num)
		return &ni
	}

	stack := make([]NestedInteger, 0)
	si := -1
	for ei, c := range s {
		if c == '[' {
			ni := NestedInteger{}
			stack = append(stack, ni)
		} else if c == ']' {
			// [673,123]
			if si != -1 {
				num, _ := strconv.Atoi(s[si:ei])
				ni := NestedInteger{}
				ni.SetInteger(num)
				si = -1
				stack[len(stack)-1].Add(ni)
			}
			// [123,[123]]
			if len(stack) > 1 {
				l := len(stack) - 1
				inner := stack[l]
				stack = stack[0:l]
				stack[l-1].Add(inner)
			}
		}  else if c == ','{
			// [673,123]
			if si != -1 {
				num, _ := strconv.Atoi(s[si:ei])
				ni := NestedInteger{}
				ni.SetInteger(num)
				si = -1
				stack[len(stack)-1].Add(ni)
			}
		} else if si == -1 {
			si = ei
		}
	}
	return &stack[0]
}
