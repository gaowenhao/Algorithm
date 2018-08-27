/*
-----------------------------------------------------
    Author : 高文豪
    Github : https://github.com/gaowenhao
-----------------------------------------------------
*/

package sort_algorithm

import "arraytools"

/*
	选择排序,个人认为这便是最好理解的那种,从头遍历(也可以尾遍历,这里就说从头遍历的这种把.)对比所有的数字,如果小于自己则立刻交换,第二次外层循环便可忽略第一个值了,因为其值肯定为全场最小.
*/

func SelectionSort(array []int) {
	length := len(array)

	for x := 0; x < length; x++ {
		for y := x + 1; y < length; y++ {
			if array[y] < array[x] {
				arraytools.Swap(array, x, y)
			}
		}
	}
}
