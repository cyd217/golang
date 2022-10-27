package main

import "container/list"

type LruCache interface {
	// 设置/添加一个缓存，如果key存在，则用新值覆盖旧值
	Add(key, value interface{})
	// 通过key获取一个缓存值
	Get(key interface{}) (value interface{}, exist bool)
	// 通过key删除一个缓存值
	Remove(key interface{})
	// 删除 '最无用' 的一个缓存值
	RemoveOldest()
	// 获取缓存已存在的元素个数
	Len()
	// 缓存中 元素 已经所占用内存的大小
	Clear()
}

type LruCacheImpl struct {
	//MaxEntries 是 Cache 中实体的最大数量，0 表示没有限制
	MaxEntries int

	Call func(key, value interface{})

	ll *list.List

	mp map[interface{}]*list.Element
}

//一个 entry 包含一个 key 和一个 value，都是任意类型
type entry struct {
	key   interface{}
	value interface{}
}

//创建一个 LRU Cache。maxEntries 为 0 表示缓存没有大小限制
func New(maxEntries int) *LruCacheImpl {
	return &LruCacheImpl{
		MaxEntries: maxEntries,
		ll:         list.New(),
		mp:         make(map[interface{}]*list.Element, maxEntries),
	}
}

//向 Cache 中插入一个 KV
func (c *LruCacheImpl) Add(key, value interface{}) {
	if c.mp == nil {
		c.mp = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}

	if ee, ok := c.mp[key]; ok {
		//如果存在，移动至列表头部  list的每个value 都是一个entry节点
		c.ll.MoveToFront(ee)
		ee.Value.(*entry).value = value
	}
	// 头部插入一个元素并返回该元素
	ele := c.ll.PushFront(&entry{key, value})

	//mp的key是entry节点的key，value是list的节点（可以理解entry节点）
	c.mp[key] = ele

	//超过则删除最老的元素
	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		c.RemoveOldest()
	}
}

//传入一个 key，返回一个是否有该 key 以及对应 value
func (c *LruCacheImpl) Get(key interface{}) (value interface{}, exist bool) {
	if c.mp == nil {
		return
	}
	if ee, ok := c.mp[key]; ok {
		//如果存在，移动至列表头部  list的每个value 都是一个entry节点
		c.ll.MoveToFront(ee)
		return ee.Value.(*entry).value, ok
	}
	return
}

//从 Cache 中删除一个 KV
func (c *LruCacheImpl) Remove(key interface{}) {
	if c.mp == nil {
		return
	}
	if ele, ok := c.mp[key]; ok {
		c.removeElement(ele)
	}
}

//从 Cache 中删除最久未被访问的数据
func (c *LruCacheImpl) RemoveOldest() {
	if c.mp == nil {
		return
	}
	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}

//从 Cache 中删除一个元素，供内部调用
func (c *LruCacheImpl) removeElement(e *list.Element) {
	//先从 list 中删除
	c.ll.Remove(e)

	kv := e.Value.(*entry)
	//再从 map 中删除
	delete(c.mp, kv.key)

	//如果回调函数不为空则调用
	if c.Call != nil {
		c.Call(kv.key, kv.value)
	}
}

//获取 Cache 当前的元素个数
func (c *LruCacheImpl) Len() int {
	if c.mp == nil {
		return 0
	}
	return c.ll.Len()
}

//清空 Cache
func (c *LruCacheImpl) Clear() {
	if c.Call != nil {
		for _, e := range c.mp {
			kv := e.Value.(*entry)
			c.Call(kv.key, kv.value)
		}
	}
	c.ll = list.New()
	c.mp = make(map[interface{}]*list.Element)
}
