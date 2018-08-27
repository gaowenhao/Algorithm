package btree

/*
-----------------------------------------------------
    Author : 高文豪
    Github : https://github.com/gaowenhao
-----------------------------------------------------
*/

/*
	实现2-3树 为了方便仅实现了 插入-调节功能 ,这种数据结构情况略多所以最后重构了一下代码,看上去还算舒服.
*/
import "fmt"

// 2-3树
type _23Node struct {
	parent        *_23Node
	left          *_23Node
	middle        *_23Node
	right         *_23Node
	array         []int
	leftTempNode  *_23Node
	rightTempNode *_23Node
}

func Insert23Node(node *_23Node, data int) *_23Node {
	length := len(node.array)
	if length == 0 {
		node.array = append(node.array, data)
		return node
	} else if !HasAnyChild(node) {
		if data > node.array[0] { // 值在右面
			node.array = append(node.array, data)
		} else { // 值在左面
			node.array = append([]int{data}, node.array[:]...)
		}
	} else {
		if data < node.array[0] {
			node.left = Insert23Node(node.left, data)
		} else if data > node.array[0] {
			if length == 2 && node.array[1] >= data {
				node.middle = Insert23Node(node.middle, data)
			} else {
				node.right = Insert23Node(node.right, data)
			}
		}
	}

	Adjust(node)
	return node
}

func Adjust(node *_23Node) {
	length := len(node.array)
	if length > 2 {
		//根节点的分裂
		if node.parent == nil {
			// 如果根节点没有中心节点类似的东西
			if node.leftTempNode == nil && node.rightTempNode == nil {
				leftChild := New23Node(node, nil, nil, nil, []int{node.array[0]})
				rightChild := New23Node(node, nil, nil, nil, []int{node.array[2]})

				node.array = append(node.array[0:0], node.array[1])

				node.left = leftChild
				node.right = rightChild
			} else {
				// 如果根节点被设置了值
				leftChild := New23Node(node, node.left, nil, node.leftTempNode, []int{node.array[0]})
				rightChild := New23Node(node, node.rightTempNode, nil, node.right, []int{node.array[2]})

				// 左子赋值
				node.left.parent = leftChild
				node.leftTempNode.parent = leftChild

				// 右子赋值
				node.rightTempNode.parent = rightChild
				node.right.parent = rightChild

				//node 置空
				node.leftTempNode = nil
				node.rightTempNode = nil
				node.middle = nil
				node.parent = nil

				// node设置值
				node.array = append(node.array[0:0], node.array[1])
				node.left = leftChild
				node.right = rightChild
			}
		} else {
			// 如果父节点只有一个元素
			if len(node.parent.array) == 1 {
				if node == node.parent.left {
					// 插入左面
					node.parent.array = append(node.parent.array[0:0], []int{node.array[1], node.parent.array[0]}...)
					middle := New23Node(node.parent, node.rightTempNode, nil, node.right, []int{node.array[2]})
					node.array = append(node.array[0:0], node.array[0])
					node.right = node.leftTempNode
					node.parent.middle = middle

				} else {
					node.parent.array = append(node.parent.array, node.array[1])
					middle := New23Node(node.parent, node.left, nil, node.leftTempNode, []int{node.array[0]})

					node.array = append(node.array[0:0], node.array[2])
					node.left = node.rightTempNode

					node.parent.middle = middle
				}
			} else if len(node.parent.array) == 2 {
				// 如果父节点有两个元素
				if node == node.parent.left {
					node.parent.array = append(node.parent.array[0:0], []int{node.array[1], node.parent.array[0], node.parent.array[1]}...)
					right := New23Node(node.parent, node.rightTempNode, nil, node.right, []int{node.array[2]})

					node.array = append(node.array[0:0], node.array[0])
					node.right = node.leftTempNode
					node.parent.leftTempNode = right
					node.parent.rightTempNode = node.parent.middle
					node.parent.middle = nil
				} else {
					node.parent.array = append(node.parent.array, node.array[1])
					left := New23Node(node.parent, node.left, nil, node.leftTempNode, []int{node.array[0]})
					node.array = append(node.array[0:0], node.array[2])
					node.left = node.rightTempNode
					node.parent.leftTempNode = node.parent.middle
					node.parent.rightTempNode = left
					node.parent.middle = nil
				}
			}
		}
	}
}

// 判断是否有子节点.
func HasAnyChild(node *_23Node) bool {
	if node.left == nil && node.middle == nil && node.right == nil {
		return false
	} else {
		return true
	}
}

func New23Node(parent *_23Node, left *_23Node, middle *_23Node, right *_23Node, array []int) *_23Node {
	return &_23Node{parent, left, middle, right, array, nil, nil}
}

func Test23Tree() {
	node := New23Node(nil, nil, nil, nil, []int{15})
	Insert23Node(node, 14)
	Insert23Node(node, 13)
	Insert23Node(node, 12)
	Insert23Node(node, 11)
	Insert23Node(node, 10)
	Insert23Node(node, 9)
	Insert23Node(node, 8)
	Insert23Node(node, 7)
	Insert23Node(node, 6)
	Insert23Node(node, 5)
	Insert23Node(node, 4)
	Insert23Node(node, 3)
	Insert23Node(node, 2)
	Insert23Node(node, 1)
	fmt.Print(node)
}
