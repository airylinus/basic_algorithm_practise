package common

import "fmt"

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func (ln *LinkNode) Print() string {
	output := ""
	p := ln
	for p != nil {
		if output == "" {
			output += fmt.Sprintf("%d", p.Val)
		} else {
			output += fmt.Sprintf(" -> %d", p.Val)
		}
		p = p.Next
	}
	return output
}

func MakeLinkNode(in []int) *LinkNode {
	var end *LinkNode
	var head *LinkNode
	for i := len(in) - 1; i >= 0; i-- {
		head = &LinkNode{Val: in[i]}
		head.Next = end
		end = head
	}
	return head
}
