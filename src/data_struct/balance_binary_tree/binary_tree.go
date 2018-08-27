/*
-----------------------------------------------------
    Author : 高文豪
    Github : https://github.com/gaowenhao
-----------------------------------------------------
*/

package balance_binary_tree

import (
	"fmt"
)

/*
	二叉树的实现, 自带二叉查找树功能.
*/

type BinaryTreeNode struct {
	left  *BinaryTreeNode
	right *BinaryTreeNode
	data  Object
}

// 打印Node
func (node *BinaryTreeNode) Show() {
	fmt.Println(node.data)
}

// 创建一个Node
func NewBinaryTreeNode(left *BinaryTreeNode, right *BinaryTreeNode, data Object) *BinaryTreeNode {
	return &BinaryTreeNode{left, right, data}
}

// -----------------------------------------------------------------------

type BinaryTree struct {
	root *BinaryTreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{nil}
}

// 堆树构建,测试遍历.
func (binaryTree *BinaryTree) BuildBinaryTreeByArray(array []Object, node *BinaryTreeNode, index int) *BinaryTreeNode {
	var newNode *BinaryTreeNode

	if node == nil {
		leftIndex := 2*index + 1
		rightIndex := 2*index + 2
		newNode = NewBinaryTreeNode(binaryTree.BuildBinaryTreeByArray(array, NewBinaryTreeNode(nil, nil, array[leftIndex]), leftIndex),
			binaryTree.BuildBinaryTreeByArray(array, NewBinaryTreeNode(nil, nil, array[rightIndex]), rightIndex), array[index])
		binaryTree.root = newNode
	} else {
		length := len(array)
		if index >= length>>1 {
			return node
		} else {
			leftIndex := 2*index + 1
			rightIndex := 2*index + 2

			if rightIndex >= length {
				left := binaryTree.BuildBinaryTreeByArray(array, NewBinaryTreeNode(nil, nil, array[leftIndex]), leftIndex)
				return NewBinaryTreeNode(left, nil, array[index])
			} else {
				left := binaryTree.BuildBinaryTreeByArray(array, NewBinaryTreeNode(nil, nil, array[leftIndex]), leftIndex)

				right := binaryTree.BuildBinaryTreeByArray(array, NewBinaryTreeNode(nil, nil, array[rightIndex]), rightIndex)
				return NewBinaryTreeNode(left, right, array[index])
			}
		}
	}
	return newNode
}

// 遍历二叉树
func (binaryTree *BinaryTree) Show() {
	binaryTree._Show(binaryTree.root)
}

func (binaryTree *BinaryTree) _Show(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	fmt.Print(root.data, "-")

	binaryTree._Show(root.left)
	binaryTree._Show(root.right)
}

func (binaryTree *BinaryTree) InsertFindTree(data Object) {
	binaryTree._InsertFindTree(data, binaryTree.root)
}

// 二叉查找树 插入算法
func (binaryTree *BinaryTree) _InsertFindTree(data Object, node *BinaryTreeNode) bool {
	if node == nil {
		binaryTree.root = NewBinaryTreeNode(nil, nil, data)
		return true
	}

	currentNode := node
	if data.(int) <= currentNode.data.(int) {
		if node.left == nil {
			node.left = NewBinaryTreeNode(nil, nil, data)
		} else {
			binaryTree._InsertFindTree(data, node.left)
		}
	} else {
		if node.right == nil {
			node.right = NewBinaryTreeNode(nil, nil, data)
		} else {
			binaryTree._InsertFindTree(data, node.right)
		}
	}

	return true
}

// 二叉查找树 搜索
func (binaryTree *BinaryTree) SearchFindTree(data Object, node *BinaryTreeNode) *BinaryTreeNode {
	if binaryTree.root == nil {
		return nil
	}

	if data.(int) < (node.data).(int) {
		node = binaryTree.SearchFindTree(data, node.left)
	} else if data.(int) > node.data.(int) {
		node = binaryTree.SearchFindTree(data, node.right)
	}
	return node
}

func TestBinaryTree() {
	tree := NewBinaryTree()

	tree.InsertFindTree(1)

	tree.InsertFindTree(2)

	tree.InsertFindTree(2)

	tree.InsertFindTree(6)

	tree.InsertFindTree(3)

	tree.Show()

}
