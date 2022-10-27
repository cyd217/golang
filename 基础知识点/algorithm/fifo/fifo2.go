package main

//队列的概念：队列是一种特殊的线性表，特殊之处在于它只允许在表的前端（front）进行删除操作，
//而在表的后端（rear）进行插入操作，和栈一样，队列是一种操作受限制的线性表。
//进行插入操作的端称为队尾，进行删除操作的端称为队头。

type IQueue interface {
	//Push 入队 队尾
	Push(data interface{})

	//Pop 出队 队头
	Pop() interface{}

	//Get 获取队列中 index 位子的data
	Get(index int) interface{}

	//Remove 设置队列中 index 位子的data
	Remove(index int) interface{}

	//RemoveAll 移除队列中的所有数据
	RemoveAll()

	//Size 当前队列大小
	Size() int
}

type BQueue struct {
	MaxEntries int
	Datas      []interface{}
}

//NewBQueue 初始化队列
func NewBQueue(maxEntries int) *BQueue {
	return &BQueue{
		Datas: make([]interface{}, maxEntries),
	}
}

func (b *BQueue) Push(data interface{}) {

	if b.MaxEntries <= 0 {
		return
	}

	if len(b.Datas) == b.MaxEntries {
		b.Pop()
	}
	b.Datas = append(b.Datas, data)
}

func (b *BQueue) Pop() interface{} {

	if len(b.Datas) <= 0 {
		return nil
	}

	var data = b.Datas[0]
	b.Datas = b.Datas[1:]
	return data
}

func (b *BQueue) Get(index int) interface{} {

	if len(b.Datas) <= 0 {
		return nil
	}
	if index < 0 || index >= len(b.Datas) {
		return nil
	}
	return b.Datas[index]
}

func (b *BQueue) RemoveAll() {

	b.Datas = make([]interface{}, 0)
}

func (b *BQueue) Size() int {

	return len(b.Datas)
}
