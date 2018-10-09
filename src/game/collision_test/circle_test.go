/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
	Blog   : gaowenhao.cn
-----------------------------------------------------
*/

package collision_test

import "math"

/*
	判断两个圆形是否碰撞,这个其实很简单就是判断两个圆的欧氏距离（二维的欧氏距离就是x和y）是否小于两个圆的半径和(而两个圆半径和就是两个圆最近的距离)
*/

type CircleObject struct {
	x      int // 圆心x
	y      int // 圆心y
	radius int // 半径
}

func CircleTest(obj1 CircleObject, obj2 CircleObject) bool {
	return int(math.Sqrt(math.Pow(float64(obj1.x-obj2.x), 2)+math.Pow(float64(obj1.y-obj2.y), 2))) < obj1.radius+obj2.radius
}
