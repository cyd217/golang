package main

//基于泛型的队列
type MyStack[T interface{}] struct {
	elements []T
}

func (q *MyStack[T]) Push(value T) {
	q.elements = append(q.elements, value)
}

func (q *MyStack[T]) Pull() (value T, isEmtry bool) {
	if len(q.elements) == 0 {
		return value, true
	}
	value = q.elements[len(q.elements)-1]
	q.elements = q.elements[:len(q.elements)-1]
	return value, len(q.elements) == 0
}

// 队列大小
func (q MyStack[T]) Size() int {
	return len(q.elements)
}
