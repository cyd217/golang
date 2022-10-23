package lfu

import "container/list"

type LFUCache2 struct {
	kItems map[int]*list.Element
	fItems map[int]*list.List
	cap    int
	minFre int
}
type entry2 struct {
	key, value int
	freq       int
}

func Constructor(capacity int) LFUCache2 {
	lfu := LFUCache2{
		kItems: make(map[int]*list.Element),
		fItems: make(map[int]*list.List),
		cap:    capacity,
		minFre: 1,
	}
	lfu.fItems[lfu.minFre] = list.New() //这个频率是一定会用到的，提前申请好
	return lfu
}

/***
  kItems 是否存在
	if 存在
		更新kItems中键值
		更新频率列表fItems nodeExec() ,
		更新最小频率
	if 不存在
		检查是否已满
		if 满
		      删除最近最少使用 removeOld()
		创建节点 频率为1
		kItems集合添加节点
		fItems集合列表添加节点
		更新最小频率为1
*/
func (l *LFUCache2) Put(key int, value int) {
	if l.cap <= 0 {
		return
	}

	node, ok := l.kItems[key]
	//存在
	if ok {
		node.Value.(*entry2).value = value
		l.nodeExec(node)
		return
	}
	//该键值不存在

	if len(l.kItems) >= l.cap { //如果lfu满了
		l.removeOld()
	}
	kv := &entry2{key: key, value: value, freq: 1}
	node = l.fItems[kv.freq].PushFront(kv)
	l.kItems[key] = node
	l.minFre = 1
}

//删除最近最少使用节点
func (l *LFUCache2) removeOld() {
	list := l.fItems[l.minFre]
	node := list.Back()
	list.Remove(node)
	delete(l.kItems, node.Value.(*entry2).key)
}

/***
kItems是否存在key
if 存在
	新频率列表fItems nodeExec()
if 不存在
	-1
*/
func (l *LFUCache2) Get(key int) int {
	if len(l.kItems) == 0 {
		return -1
	}
	node, ok := l.kItems[key]
	if !ok {
		return -1
	}
	value := node.Value.(*entry2).value
	l.nodeExec(node)
	return value
}

func (l *LFUCache2) nodeExec(node *list.Element) {
	//原频率中删除
	kv := node.Value.(*entry2)
	oldList := l.fItems[kv.freq]
	oldList.Remove(node)

	//更新minfreq
	if oldList.Len() == 0 && l.minFre == kv.freq {
		l.minFre++
	}
	//放入新的频率链表
	kv.freq++
	if _, ok := l.fItems[kv.freq]; !ok {
		l.fItems[kv.freq] = list.New()
	}
	newList := l.fItems[kv.freq]
	node = newList.PushFront(kv)
	l.kItems[kv.key] = node
}

// type LRUCache struct {
// 	capacity   int
// 	m          map[int]*Node
// 	head, tail *Node
// }

// type Node struct {
// 	Key       int
// 	Value     int
// 	Pre, Next *Node
// }

// func (this *LRUCache) Get(key int) int {
// 	if v, ok := this.m[key]; ok {
// 		this.moveToHead(v)
// 		return v.Value
// 	}
// 	return -1
// }

// func (this *LRUCache) moveToHead(node *Node) {
// 	this.deleteNode(node)
// 	this.addToHead(node)
// }

// func (this *LRUCache) deleteNode(node *Node) {
// 	node.Pre.Next = node.Next
// 	node.Next.Pre = node.Pre
// }

// func (this *LRUCache) removeTail() int {
// 	node := this.tail.Pre
// 	this.deleteNode(node)
// 	return node.Key
// }

// func (this *LRUCache) addToHead(node *Node) {
// 	this.head.Next.Pre = node
// 	node.Next = this.head.Next
// 	node.Pre = this.head
// 	this.head.Next = node
// }

// func (this *LRUCache) Put(key int, value int) {
// 	if v, ok := this.m[key]; ok {
// 		v.Value = value
// 		this.moveToHead(v)
// 		return
// 	}

// 	if this.capacity == len(this.m) {
// 		rmKey := this.removeTail()
// 		delete(this.m, rmKey)
// 	}

// 	newNode := &Node{Key: key, Value: value}
// 	this.addToHead(newNode)
// 	this.m[key] = newNode
// }

// func Constructor(capacity int) LRUCache {
// 	head, tail := &Node{}, &Node{}
// 	head.Next = tail
// 	tail.Pre = head
// 	return LRUCache{
// 		capacity: capacity,
// 		m:        map[int]*Node{},
// 		head:     head,
// 		tail:     tail,
// 	}
// }
