package leetcode

type MyLinkedList struct {
	head *MyNode
	tail *MyNode
	size int
}

type MyNode struct {
	Val  int
	Next *MyNode
	Prev *MyNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Get 获取链表中第 index 个节点的值。如果索引无效，则返回-1。
func (this *MyLinkedList) Get(index int) int {
	// 检查索引范围
	if index < 0 || index >= this.size {
		return -1
	}
	curr := this.head
	for i := 0; i < index && curr != nil; i++ {
		curr = curr.Next
	}
	if curr != nil {
		return curr.Val
	} else { // 索引无效
		return -1
	}
}

// AddAtHead 在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
func (this *MyLinkedList) AddAtHead(val int) {
	node := &MyNode{Val: val}
	// 头插法
	node.Next = this.head
	if this.head != nil {
		this.head.Prev = node
	}
	// 更新头节点
	this.head = node
	// 更新尾节点
	if this.tail == nil {
		this.tail = node
	}
	this.size++
}

// AddAtTail 将值为 val 的节点追加到链表的最后一个元素。
func (this *MyLinkedList) AddAtTail(val int) {
	node := &MyNode{Val: val}
	// 尾插法
	node.Prev = this.tail
	if this.tail != nil {
		this.tail.Next = node
	}
	// 更新尾节点
	this.tail = node
	// 更新头节点
	if this.head == nil {
		this.head = node
	}
	this.size++
}

// AddAtIndex 在链表中的第 index 个节点之前添加值为 val  的节点。
// 如果 index 等于链表的长度，则该节点将附加到链表的末尾。
// 如果 index 大于链表长度，则不会插入节点。
// 如果index小于0，则在头部插入节点。
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == this.size {
		this.AddAtTail(val)
		return
	}
	if index > this.size {
		return
	}
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	node := &MyNode{Val: val}
	curr := this.head
	for i := 0; i < index-1 && curr != nil; i++ {
		curr = curr.Next
	}
	if curr != nil {
		node.Prev = curr
		node.Next = curr.Next
		if node.Next != nil {
			node.Next.Prev = node
		}
		//node.Next.Prev = node
		curr.Next = node
		this.size++
	}
}

// DeleteAtIndex 如果索引 index 有效，则删除链表中的第 index 个节点。
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size || this.size == 0 {
		return
	}

	if this.size == 1 {
		this.head = nil
		this.tail = nil
		this.size--
		return
	}

	if index == 0 { // 删除头结点
		this.head = this.head.Next
		if this.head != nil {
			this.head.Prev = nil
		}
	} else if index == this.size-1 { // 删除尾节点
		this.tail = this.tail.Prev
		if this.tail != nil {
			this.tail.Next = nil
		}
	} else { // 删除中间节点
		curr := this.head
		for i := 0; i < index-1 && curr != nil; i++ {
			curr = curr.Next
		}
		if curr != nil && curr.Next != nil {
			curr.Next = curr.Next.Next
			if curr.Next != nil {
				curr.Next.Prev = curr
			}
			//curr.Next.Prev = curr
		}
	}
	this.size--
}
