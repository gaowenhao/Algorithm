/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
	Blog   : gaowenhao.cn
-----------------------------------------------------
*/

package balance_binary_tree

/*
	AVL树,实现插入,删除操作.
*/
import "fmt"

type BalanceBinaryTreeNode struct {
	left   *BalanceBinaryTreeNode
	right  *BalanceBinaryTreeNode
	data   Object
	height int
}

// 创建一个Node
func NewBalanceBinaryTreeNode(left *BalanceBinaryTreeNode, right *BalanceBinaryTreeNode, data Object, height int) *BalanceBinaryTreeNode {
	return &BalanceBinaryTreeNode{left, right, data, height}
}

// -------------------------------BalanceTree-------------------------------------------

type BalanceBinaryTree struct {
	root *BalanceBinaryTreeNode
}

// 构件一个树
func NewBalanceBinaryTree(data Object) *BalanceBinaryTree {
	return &BalanceBinaryTree{NewBalanceBinaryTreeNode(nil, nil, data, 1)}
}

// 插入一个元素
func (tree *BalanceBinaryTree) InsertNode(data Object) {
	tree.root = Insert(data, tree.root)
}

func (tree *BalanceBinaryTree) DeleteNode(data Object) {
	tree.root = Delete(data, tree.root)
}

func Insert(data Object, node *BalanceBinaryTreeNode) *BalanceBinaryTreeNode {
	if node == nil {
		newNode := NewBalanceBinaryTreeNode(nil, nil, data, 1)
		return newNode
	}
	if data.(int) < node.data.(int) {
		node.left = Insert(data, node.left)
		node = Rebalance(node)
	} else if data.(int) > node.data.(int) {
		node.right = Insert(data, node.right)
		node = Rebalance(node)
	} else {
		fmt.Print("同数")
		return nil
	}
	node.height = Max(GetHeight(node.left), GetHeight(node.right)) + 1
	return node
}

func Rebalance(node *BalanceBinaryTreeNode) *BalanceBinaryTreeNode {
	right := node.right
	left := node.left
	if GetHeight(right)-GetHeight(left) == 2 {
		if GetHeight(right.right) > GetHeight(right.left) {
			//左旋
			node = LeftRotate(node)
		} else {
			node = RightThenLeftRotate(node)
		}
	} else if GetHeight(left)-GetHeight(right) == 2 {
		if GetHeight(left.left) > GetHeight(left.right) {
			//右旋
			node = RightRotate(node)
		} else {
			//左旋后右旋
			node = LeftThenRightRotate(node)
		}
	}
	return node
}

func Delete(data Object, node *BalanceBinaryTreeNode) *BalanceBinaryTreeNode {
	if node == nil {
		return nil
	}

	if node.data.(int) == data.(int) {
		if node.left == nil && node.right == nil {
			return nil
		} else if node.left == nil || node.right == nil {
			if node.left != nil {
				return node.left
			} else {
				return node.right
			}
		} else {
			var n *BalanceBinaryTreeNode
			n = node.left
			for {
				if n.right == nil {
					break
				}
				n = n.right
			}

			n.data, node.data = node.data, n.data
			node.left = Delete(n.data, node.left)
		}
	} else if node.data.(int) > data.(int) {
		node.left = Delete(data, node.left)
	} else {
		node.right = Delete(data, node.right)
	}
	node.height = Max(GetHeight(node.left), GetHeight(node.right)) + 1
	node = Rebalance(node)
	return node
}

// 左旋
func LeftRotate(node *BalanceBinaryTreeNode) *BalanceBinaryTreeNode {
	newNode := node.right
	node.right = newNode.left
	newNode.left = node
	node.height = Max(GetHeight(node.left), GetHeight(node.right)) + 1
	newNode.height = Max(GetHeight(newNode.left), GetHeight(newNode.right)) + 1
	return newNode
}

// 右旋
func RightRotate(node *BalanceBinaryTreeNode) *BalanceBinaryTreeNode {
	newNode := node.left
	node.left = newNode.right
	newNode.right = node
	node.height = Max(GetHeight(node.left), GetHeight(node.right)) + 1
	newNode.height = Max(GetHeight(node), GetHeight(newNode.left)) + 1
	return newNode
}

// 左旋后右旋
func LeftThenRightRotate(node *BalanceBinaryTreeNode) *BalanceBinaryTreeNode {
	newNode := LeftRotate(node.left) //左旋转
	node.left = newNode
	return RightRotate(node)

}

// 右旋后左旋
func RightThenLeftRotate(node *BalanceBinaryTreeNode) *BalanceBinaryTreeNode {
	newNode := RightRotate(node.right)
	node.right = newNode
	node.right = newNode
	return LeftRotate(node)
}

func GetHeight(node *BalanceBinaryTreeNode) int {
	if node == nil {
		return 0
	} else {
		return node.height
	}
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func TestBalanceBinaryTree() {
	tree := NewBalanceBinaryTree(1)
	tree.InsertNode(3)
	tree.InsertNode(4)
	tree.InsertNode(2)

	tree.DeleteNode(3)
	fmt.Print("1")
}
