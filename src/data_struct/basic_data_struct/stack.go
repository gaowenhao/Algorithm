/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
	Blog   : gaowenhao.cn
-----------------------------------------------------
*/

package basic_data_struct

import "fmt"

/*
	栈的实现(使用数组实现了,后面的数据结构都使用数组实现 有几个特殊的例外)
*/

// 初始化长度
const StackInitSize int = 20

// 是否自动扩展长度,自动扩展就和Java中的ArrayList差不多,只不过我这里的扩展相对简单,Golang的切片扩展起来相对方便,不需要手动复制数组了.
const StackIsAutomaticExpand bool = true

type Stack struct {
	data []Object
	size int
	top  int
}

// 构建新栈
func NewStack() *Stack {
	return &Stack{make([]Object, 0), StackInitSize, -1}
}

// 进栈
func (stack *Stack) Push(data Object) bool {
	if stack.IsFull() {
		if StackIsAutomaticExpand {
			stack.size += stack.size >> 1
		} else {
			return false
		}
	}

	stack.data = append(stack.data, data)
	stack.top++
	return true
}

// 出栈
func (stack *Stack) Pop() Object {
	if stack.top < 0 || stack.top > stack.size {
		return nil
	}

	ret := stack.data[stack.top]
	stack.data = append(stack.data[:stack.top], stack.data[stack.top+1:]...)
	stack.top--
	return ret
}

// 判断栈是否已满
func (stack *Stack) IsFull() bool {
	return stack.top == stack.size-1
}

func TestStack() {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Print(stack.Pop())
	fmt.Print(stack.Pop())
	fmt.Print(stack.Pop())
	fmt.Print(stack.Pop())
	stack.Push(5)
	fmt.Print(stack.Pop())
}
