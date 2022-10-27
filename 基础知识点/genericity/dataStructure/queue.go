package main

//基于泛型的队列
type MyQueue[T interface{}] struct {
	elements []T
}

func (q *MyQueue[T]) Push(value T) {
	q.elements = append(q.elements, value)
}

func (q *MyQueue[T]) Pop() (value T, isEmtry bool) {
	if len(q.elements) == 0 {
		return value, true
	}
	value = q.elements[0]
	q.elements = q.elements[1:]
	return value, len(q.elements) == 0
}

// 队列大小
func (q MyQueue[T]) Size() int {
	return len(q.elements)
}
