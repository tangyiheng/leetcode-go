package leetcode

const Len int = 10000

type HashNode struct {
	key  int
	val  int
	next *HashNode
}

func (N *HashNode) Put(key int, value int) {
	// 覆盖
	if N.key == key {
		N.val = value
		return
	}
	// 追加到链表末尾
	if N.next == nil {
		N.next = &HashNode{key, value, nil}
		return
	}
	N.next.Put(key, value)
}

func (N *HashNode) Get(key int) int {
	if N.key == key {
		return N.val
	}
	// 不存在
	if N.next == nil {
		return -1
	}
	return N.next.Get(key)
}

func (N *HashNode) Remove(key int) *HashNode {
	if N.key == key {
		p := N.next
		N.next = nil
		return p
	}
	// 在链表上查找
	if N.next != nil {
		N.next = N.next.Remove(key)
	}
	return N
}

type MyHashMap struct {
	content [Len]*HashNode
}

func Constructor706() MyHashMap {
	return MyHashMap{}
}

// func Constructor() MyHashMap {
//
// }

func (this *MyHashMap) Hash(key int) int {
	return key % Len
}

func (this *MyHashMap) Put(key int, value int) {
	node := this.content[this.Hash(key)]
	if node == nil {
		this.content[this.Hash(key)] = &HashNode{key: key, val: value, next: nil}
		return
	}
	node.Put(key, value)
}

func (this *MyHashMap) Get(key int) int {
	node := this.content[this.Hash(key)]
	if node == nil {
		return -1
	}
	return node.Get(key)
}

func (this *MyHashMap) Remove(key int) {
	node := this.content[this.Hash(key)]
	if node == nil {
		return
	}
	this.content[this.Hash(key)] = node.Remove(key)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
