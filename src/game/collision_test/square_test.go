/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
    Blog   : gaowenhao.cn
-----------------------------------------------------
*/

package collision_test

/*
	判断2D物体两个方形是否碰撞.
*/

type SquareObject struct {
	x      int //  起点x
	y      int //  起点y
	width  int //  宽
	height int //  高
}

func SquareTest(obj1 SquareObject, obj2 SquareObject) bool {
	return obj1.x < obj2.x+obj2.width && obj1.x+obj1.width > obj2.x && obj1.y < obj2.y+obj2.height && obj1.y+obj1.height > obj2.y
}
