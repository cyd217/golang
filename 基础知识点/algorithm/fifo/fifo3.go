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

type FIFOCache3 struct {
	capacity int
	size     int
	list     *List
	find     map[interface{}]*Node
}

func InitFIFO3(capacity int) *FIFOCache3 {
	return &FIFOCache3{
		capacity: capacity,
		find:     map[interface{}]*Node{},
		list:     NewList(),
	}
}

func (l *FIFOCache3) Put(key, value interface{}) {
	if l.capacity == 0 {
		return
	}
	if value, ok := l.find[key]; ok {
		node := value
		l.list.remove(node)
		node.value = value
		l.list.AddTail(node)
	} else {
		if l.size == l.capacity {
			node := l.list.RemoveHead()
			delete(l.find, node.key)
			l.size--
		}
		node := InitNode(key, value)
		l.list.AddTail(node)
		l.find[key] = node
		l.size++
	}

}
