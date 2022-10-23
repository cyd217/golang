package lfu

type Node2 struct {
	Key   int
	Value int
	freq  int
	pre   *Node2
	next  *Node2
}

type LFUCache struct {
	limit   int
	HashMap map[int]*Node2
	head    *Node2
	end     *Node2
}

func LFUConstructor(capacity int) LFUCache {
	lfuCache := LFUCache{}
	lfuCache.limit = capacity
	lfuCache.HashMap = make(map[int]*Node2, capacity)
	return lfuCache
}

func (l *LFUCache) Get(key int) int {
	if value, ok := l.HashMap[key]; ok {
		value.freq++
		l.refreshNode(value)
		return value.Value
	} else {
		return -1
	}
}

func (l *LFUCache) Put(key, value int) {
	if v, ok := l.HashMap[key]; !ok {
		if len(l.HashMap) >= l.limit {
			oldKey := l.removeNode(l.head)
			delete(l.HashMap, oldKey)
		}
		node := Node2{Key: key, Value: value, freq: 1}
		l.addNode(&node)
		l.HashMap[key] = &node
	} else {
		v.Value = value
		v.freq++
		l.refreshNode(v)
	}
}

func (l *LFUCache) refreshNode(node *Node2) {
	if node == l.end {
		return
	}
	l.removeNode(node)
	l.addNode(node)
}

func (l *LFUCache) removeNode(node *Node2) int {
	if node == l.end {
		l.end = l.end.pre
		l.end.next = nil
	} else if node == l.head {
		l.head = l.head.next
		l.head.pre = nil
	} else {
		node.pre.next = node.next
		node.next.pre = node.pre
	}
	return node.Key
}

func (l *LFUCache) addNode(node *Node2) {

	if l.head == nil && l.end == nil {
		l.head = node
		l.end = node
		return
	}
	head := l.head
	for head != nil && node.freq >= head.freq {
		head = head.next
	}
	if head == nil {
		l.end.next = node
		node.pre = l.end
		l.end = node
	}
	if head != nil {
		head.pre.next = node
		node.pre = head.pre
		head.pre = node
		node.next = head
	}

	l.head.pre = nil
	l.end.next = nil
	return
}
