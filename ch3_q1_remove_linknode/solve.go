package ch3_q1_remove_linknode

import (
	"github.com/airylinus/basic_algorithm_practise/common"
)

// removeLinkNode
// 1 --> 4 --> 5 --> 6
// If
// we assumed targetVal will never show up at Head node / End node
// declare pointer p, if p's Next value equals target, then remove p's Next
//
//	`p.Next = p.Next.Next`
func removeLinkNodeFirstTry(head *common.LinkNode, targetVal int) (newHead *common.LinkNode) {
	//fmt.Println(head.Print())
	p := head
	for p != nil && p.Next != nil {
		if p.Next.Val == targetVal {
			p.Next = p.Next.Next
		}
		p = p.Next
	}
	//fmt.Println(head.Print())
	return head
}

func removeLinkNode(head *common.LinkNode, targetVal int) (newHead *common.LinkNode) {
	p := head
	// the target will never be the last node
	for p != nil {
		// found target
		if p.Val == targetVal {
			p.Val = p.Next.Val
			p.Next = p.Next.Next
		}
		p = p.Next
	}
	return head
}
