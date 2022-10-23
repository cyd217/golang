package main

import "container/list"

type FifoCache interface {
	//Push 入队
	Push(data interface{})

	//Pop 出队
	Pop() interface{}

	//Size 当前队列大小
	Len() int
}

//一个 entry 包含一个 key 和一个 value，都是任意类型
type entry struct {
	key   interface{}
	value interface{}
}

// TODO: 定义fifo结构体
type FifoCacheImpl struct {
	// 缓存最大容量
	MaxEntries int
	// 双链表
	ll *list.List
}

// TODO: 构造函数，创建一个新 Cache，如果 maxBytes 是0，则表示没有容量限制
func NewFifoCache(maxEntries int) *FifoCacheImpl {
	return &FifoCacheImpl{
		MaxEntries: maxEntries,
		ll:         list.New(),
	}
}

// TODO: 通过 Set 方法往 Cache 头部增加一个元素（如果已经存在，则移到头部，并修改值）
func (c *FifoCacheImpl) Push(key, value interface{}) {
	element := &entry{key, value}
	c.ll.PushFront(element)
	// 如果超出内存长度，则删除队首的节点
	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		c.Pop()
	}
}

func (c *FifoCacheImpl) Pop() (value interface{}) {
	if c.ll == nil {
		return
	}
	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
	return ele
}

//从 Cache 中删除一个元素，供内部调用
func (c *FifoCacheImpl) removeElement(e *list.Element) {
	//先从 list 中删除
	c.ll.Remove(e)
}

// TODO: 缓存中元素个数
func (f *FifoCacheImpl) Len() int {
	return f.ll.Len()
}
