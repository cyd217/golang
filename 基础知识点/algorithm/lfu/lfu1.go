package lfu

import "container/heap"

type entry struct {
	key    string
	value  interface{}
	weight int //访问次数
	index  int //queue索引

}

//二叉堆实现的队列
type queue []*entry

func (q queue) Len() int {
	return len(q)
}

// '<' 是最小堆，'>' 是最大堆
func (q queue) Less(i, j int) bool {
	return q[i].weight < q[j].weight
}

// 交换元素
func (q queue) Swap(i, j int) {
	// 交换元素
	q[i], q[j] = q[j], q[i]
	// 索引不用交换
	q[i].index = i
	q[j].index = j
}

// append ，*q = oldQue[:n-1] 会导致频繁的内存拷贝
// 实际上，如果使用 LFU算法，处于性能考虑，可以将最大内存限制修改为最大记录数限制
// 这样提前分配好 queue 的容量，再使用交换索引和限制索引的方式来实现 Pop 方法，可以免去频繁的内存拷贝，极大提高性能
func (q *queue) Push(v interface{}) {
	n := q.Len()
	en := v.(*entry)
	en.index = n
	*q = append(*q, en) // 这里会重新分配内存，并拷贝数据
}

func (q *queue) Pop() interface{} {
	oldQue := *q
	n := len(oldQue)
	en := oldQue[n-1]
	oldQue[n-1] = nil // 将不再使用的对象置为nil，加快垃圾回收，避免内存泄漏
	*q = oldQue[:n-1] // 这里会重新分配内存，并拷贝数据
	return en
}

// weight更新后，要重新排序，时间复杂度为 O(logN)
func (q *queue) update(en *entry, val interface{}, weight int) {
	en.value = val
	en.weight = weight
	(*q)[en.index] = en
	heap.Fix(q, en.index)
}

type lfuCache struct {
	// 缓存最大容量，
	maxEntity int
	// 已使用的
	usedEntity int
	// 最小堆实现的队列
	queue *queue
	// map的key是字符串，value是entry
	cache map[string]*entry
}

func NewLfuCache(maxEntity int) *lfuCache {
	queue := make(queue, 0)
	return &lfuCache{
		queue:     &queue,
		cache:     make(map[string]*entry),
		maxEntity: maxEntity,
	}
}

// 通过 Set 方法往 Cache 头部增加一个元素，如果存在则更新值
func (l *lfuCache) Set(key string, value interface{}) {
	if en, ok := l.cache[key]; ok {
		l.queue.update(en, value, en.weight+1)
	} else {
		en := &entry{
			key:   key,
			value: value,
		}

		heap.Push(l.queue, en) // 插入queue 并重新排序为堆
		l.cache[key] = en      // 插入 map
		l.usedEntity++

		// 如果超出内存长度，则删除最 '无用' 的元素，0表示无内存限制
		for l.maxEntity > 0 && l.usedEntity >= l.maxEntity {
			l.DelOldest()
		}
	}
}

// 获取指定元素,访问次数加1
func (l *lfuCache) Get(key string) interface{} {
	if en, ok := l.cache[key]; ok {
		l.queue.update(en, en.value, en.weight+1)
		return en.value
	}
	return nil
}

// 删除指定元素（删除queue和map中的val）
func (l *lfuCache) Del(key string) {
	if en, ok := l.cache[key]; ok {
		heap.Remove(l.queue, en.index)
		l.removeElement(en)
	}
}

// 删除最 '无用' 元素（删除queue和map中的val）
func (l *lfuCache) DelOldest() {
	if l.queue.Len() == 0 {
		return
	}
	val := heap.Pop(l.queue)
	l.removeElement(val)
}

// 删除元素并更新内存占用大小
func (l *lfuCache) removeElement(v interface{}) {
	if v == nil {
		return
	}
	en := v.(*entry)
	delete(l.cache, en.key)
	l.usedEntity--
}
