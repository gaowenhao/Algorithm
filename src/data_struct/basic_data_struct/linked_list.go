/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
    Blog   : gaowenhao.cn
-----------------------------------------------------
*/

package basic_data_struct

/*
	双向链表简单的增删改查,完全是根据自己的理解码的,并没有经过太多测试,如有疏漏还请指教.
*/

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

// -----------------------------LinkedList元素------------------------------

type LinkedNode struct {
	prev *LinkedNode
	next *LinkedNode
	data Object
}

// 打印Node
func (node *LinkedNode) Show() {
	fmt.Println(node.data)
}

// 创建一个Node
func NewLinkedNode(prev *LinkedNode, next *LinkedNode, data Object) *LinkedNode {
	return &LinkedNode{prev, next, data}
}

//---------------------------LinkedList--------------------------------

type LinkedList struct {
	firstNode *LinkedNode
	lastNode  *LinkedNode
	size      int
}

func NewLinkedList(firstNode *LinkedNode, lastNode *LinkedNode, size int) *LinkedList {
	return &LinkedList{firstNode, lastNode, size}
}

// 遍历
func (lst *LinkedList) Show() {
	current := lst.firstNode
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

// 插入一个值
func (lst *LinkedList) Insert(data Object) bool {
	firstNode := lst.firstNode
	newNode := NewLinkedNode(nil, firstNode, data)
	lst.firstNode = newNode

	if firstNode == nil {
		lst.lastNode = newNode
	} else {
		firstNode.prev = newNode
	}
	lst.size++
	return true
}

// 插入一个值到末尾
func (lst *LinkedList) InsertLast(data Object) bool {
	lastNode := lst.lastNode
	newNode := NewLinkedNode(lastNode, nil, data)
	lst.lastNode = newNode

	if lastNode == nil {
		lst.firstNode = newNode
	} else {
		lastNode.next = newNode
	}
	lst.size++
	return true
}

// 插入一个值到index
func (lst *LinkedList) InsertIndex(index int, data Object) bool {
	if index > lst.size {
		return false
	} else if index == lst.size { // 如果恰好插入的位置是最后位
		lst.size++
		return lst.InsertLast(data)
	}

	cursor := 0
	currentNode := lst.firstNode

	for ; currentNode != nil; cursor++ {
		if cursor == index {
			newNode := NewLinkedNode(currentNode.prev, currentNode, data)

			if currentNode.prev != nil {
				currentNode.prev.next = newNode
			} else {
				lst.firstNode = newNode
			}
			currentNode.prev = newNode

			lst.size++

			return true
		} else {
			currentNode = currentNode.next
		}
	}
	return false
}

// 插入一个值到index的后面
func (lst *LinkedList) InsertAfterIndex(index int, data Object) bool {
	if (index + 1) > lst.size {
		return false
	}

	cursor := 0
	currentNode := lst.firstNode

	for ; currentNode != nil; cursor++ {
		if cursor == index {
			newNode := NewLinkedNode(currentNode, currentNode.next, data)

			if currentNode.next != nil {
				currentNode.next.prev = newNode
			} else {
				lst.lastNode = newNode
			}
			currentNode.next = newNode

			lst.size++

			return true
		} else {
			currentNode = currentNode.next
		}
	}
	return false
}

// 删除一个值
func (lst *LinkedList) Remove(data Object) {
	currentNode := lst.firstNode
	for currentNode != nil {
		if currentNode.data == data {
			beforeNode := currentNode.prev
			afterNode := currentNode.next

			// 这地方if else太多了,有点傻,后面考虑重构一下.
			if lst.size < 3 {
				if beforeNode == nil {
					lst.firstNode = afterNode
					lst.lastNode = afterNode
				} else {
					lst.firstNode = beforeNode
					lst.lastNode = beforeNode
				}
			} else {
				if currentNode == lst.firstNode {
					lst.firstNode = afterNode
					afterNode.prev = nil
				} else if currentNode == lst.lastNode {
					lst.lastNode = beforeNode
					beforeNode.next = nil
				} else {
					beforeNode.next = afterNode
					afterNode.prev = beforeNode
				}
			}

			// 销毁当前对象,golang中似乎没有什么好的办法摧毁一个对象
			currentNode.prev = nil
			currentNode.next = nil
			currentNode.data = nil

			lst.size--

			break
		} else {
			currentNode = currentNode.next
		}
	}
}

// 依照下标设置一个值
func (lst *LinkedList) Set(index int, data Object) bool {
	if (index + 1) > lst.size {
		return false
	}

	cursor := 0
	currentNode := lst.firstNode
	for ; cursor < lst.size; cursor++ {
		if cursor == index {
			currentNode.data = data
			return true
		} else {
			currentNode = currentNode.next
		}
	}

	return false
}

// 获取对象下标,如果不存在则返回-1
func (lst *LinkedList) Index(data Object) int {
	currentNode := lst.firstNode

	for cursor := 0; currentNode != nil; cursor++ {
		if currentNode.data == data {
			return cursor
		} else {
			currentNode = currentNode.next
		}
	}
	return -1
}

// 获取长度
func (lst *LinkedList) GetSize() int {
	return lst.size
}

// 翻转链表, 双向链表翻转实在没啥好说的,单向链表也是循环最好理解
func (lst *LinkedList) Reverse() {
	currentNode := lst.firstNode
	lst.firstNode = lst.lastNode
	lst.lastNode = currentNode // 首位互换

	// 所有节点的next prev互换
	for currentNode != nil {
		next := currentNode.next
		currentNode.next = currentNode.prev
		currentNode.prev = next
		currentNode = next
	}
}

//-----------------------------------------------------------

func TestLinkedList() {
	ll := NewLinkedList(nil, nil, 0)
	p1 := &Person{"张三", 15}
	p2 := &Person{"李四", 18}
	p3 := &Person{"王老五", 21}
	ll.Insert(p1)
	ll.InsertIndex(0, p2)
	ll.InsertLast(p3)

	ll.Show()

	fmt.Println("-----------")

	ll.Reverse()

	ll.Show()
}
