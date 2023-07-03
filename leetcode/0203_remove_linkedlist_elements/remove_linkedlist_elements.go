package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	// 虚拟头节点
	dummyHead := &ListNode{}
	dummyHead.Next = head
	// 目标节点的前继节点
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			// 删除目标节点
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
