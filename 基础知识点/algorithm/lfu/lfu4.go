package lfu

import (
	"math"
)

type Node struct {
	key        interface{}
	value      interface{}
	prev, next *Node
}

type List struct {
	size int
	head *Node
	tail *Node
}

func InitNode(key, value interface{}) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}
func NewList() *List {
	return &List{
		size: 0,
	}
}

// 在双向链表中的头结点后添加节点
func (l *List) AddHead(node *Node) {
	if l.head == nil {
		l.head = node
		l.tail = node
		l.head.prev = nil
		l.tail.next = nil
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
		l.head.prev = nil
	}
	l.size++
}

// 在双向链表中的头结点后添加节点
func (l *List) AddTail(node *Node) {
	if l.tail == nil {
		l.head = node
		l.tail = node
		l.head.prev = nil
		l.tail.next = nil
	} else {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
		l.tail.next = nil
	}
	l.size++
}

func (l *List) RemoveHead() (node *Node) {

	if l.head == nil {
		return nil
	}
	node = l.head
	if node.next != nil {
		l.head = node.next
		l.head.prev = nil
	} else {
		l.head = nil
		l.tail = nil
	}
	l.size--
	return
}

func (l *List) RemoveTail() (node *Node) {

	if l.tail == nil {
		return nil
	}
	node = l.tail
	if node.prev != nil {
		l.tail = node.prev
		l.tail.next = nil
	} else {
		l.head = nil
		l.tail = nil
	}
	l.size--
	return
}

// 删除某一个双向链表中的节点
func (l *List) remove(node *Node) *Node {
	// 如果node==nil,默认删除尾节点
	if node == nil {
		node = l.tail
	}
	if node == l.tail {
		l.RemoveTail()
	} else if node == l.head {
		l.RemoveHead()
	} else {
		node.next.prev = node.prev
		node.prev.next = node.next
		l.size--
	}
	return node
}

type LFUNode struct {
	freq int
	node *Node
}

func InitLFUNode(key, value interface{}) *LFUNode {
	return &LFUNode{
		freq: 0,
		node: InitNode(key, value),
	}
}

type LFUCache4 struct {
	capacity int
	find     map[interface{}]*LFUNode
	freq_map map[int]*List
	size     int
	count    int
}

func InitLFUCahe(capacity int) *LFUCache4 {
	return &LFUCache4{
		capacity: capacity,
		find:     map[interface{}]*LFUNode{},
		freq_map: map[int]*List{},
	}
}

// 更新节点的频率
func (l *LFUCache4) updateFreq(node *LFUNode) {
	freq := node.freq
	// 删除
	node.node = l.freq_map[freq].remove(node.node)
	if l.freq_map[freq].size == 0 {
		delete(l.freq_map, freq)
	}

	freq++
	node.freq = freq
	if _, ok := l.freq_map[freq]; !ok {
		l.freq_map[freq] = NewList()
	}
	l.freq_map[freq].AddTail(node.node)
}
func findMinNum(fmp map[int]*List) int {
	min := math.MaxInt32
	for key, _ := range fmp {
		min = func(a, b int) int {
			if a > b {
				return b
			}
			return a
		}(min, key)
	}
	return min
}
func (l *LFUCache4) Get(key interface{}) interface{} {
	if _, ok := l.find[key]; !ok {
		return -1
	}
	node := l.find[key]
	l.updateFreq(node)
	return node.node.value
}

func (this *LFUCache4) Put(key, value interface{}) {
	if this.capacity == 0 {
		return
	}
	// 命中缓存
	if _, ok := this.find[key]; ok {
		node := this.find[key]
		node.node.value = value
		this.updateFreq(node)
	} else {
		if this.capacity == this.size {
			// 找到一个最小的频率
			min_freq := findMinNum(this.freq_map)
			node := this.freq_map[min_freq].RemoveHead()
			lfuNode := &LFUNode{
				node: node,
				freq: 1,
			}
			this.find[key] = lfuNode
			delete(this.find, node.key)
			this.size--
		}
		node := InitLFUNode(key, value)
		node.freq = 1
		this.find[key] = node
		if _, ok := this.freq_map[node.freq]; !ok {
			this.freq_map[node.freq] = NewList()
		}
		this.freq_map[node.freq].AddTail(node.node)
		this.size++
	}
}
