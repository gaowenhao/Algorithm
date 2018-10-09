/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
	Blog   : gaowenhao.cn
-----------------------------------------------------
*/

package basic_data_struct

import "fmt"

/**
队列,使用切片实现简单队列.
*/

// 初始化长度
const QueueInitSize int = 20

// 是否自动扩展长度,自动扩展就和Java中的ArrayList差不多,只不过我这里的扩展相对简单,Golang的切片扩展起来相对方便,不需要手动复制数组了.
const QueueIsAutomaticExpand bool = true

type Queue struct {
	data []Object
	size int
}

// 构建新栈
func NewQueue() *Queue {
	return &Queue{make([]Object, 0), StackInitSize}
}

// 推入数据
func (queue *Queue) Push(data Object) bool {
	if queue.IsFull() {
		if QueueIsAutomaticExpand {
			queue.size += queue.size >> 1
		} else {
			return false
		}
	}

	queue.data = append(queue.data, data)
	return true
}

// 取出数据
func (queue *Queue) Pop() Object {
	ret := queue.data[0]
	queue.data = append(queue.data[:0], queue.data[1:]...)
	return ret
}

// 判断队列是否已满
func (queue *Queue) IsFull() bool {
	return cap(queue.data) == queue.size
}

func TestQueue() {
	q := Queue{}
	q.Push("aaa")
	q.Push("bbb")
	fmt.Print(q.Pop())
	q.Push("ccc")
	fmt.Print(q.Pop())
}
