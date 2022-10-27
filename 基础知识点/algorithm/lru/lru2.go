package main

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

type LRUCache2 struct {
	capacity int
	find     map[interface{}]*Node
	list     *List
	k        int
	count    int
}

func InitLRU(capacity int) *LRUCache2 {
	return &LRUCache2{
		capacity: capacity,
		list:     NewList(),
		find:     make(map[interface{}]*Node),
	}
}

func (l *LRUCache2) Get(key interface{}) interface{} {
	if value, ok := l.find[key]; ok {
		node := value
		l.list.remove(node)
		l.list.AddHead(node)
		return node.value
	} else {
		return -1
	}
}

func (l *LRUCache2) Put(key, value interface{}) {
	if v, ok := l.find[key]; ok {
		node := v
		l.list.remove(node)
		node.value = value
		l.list.AddHead(node)
	} else {
		node := InitNode(key, value)
		// 缓存已经满了
		if l.list.size >= l.capacity {
			oldNode := l.list.remove(nil)
			delete(l.find, oldNode.value)
		}
		l.list.AddHead(node)
		l.find[key] = node
	}
}
