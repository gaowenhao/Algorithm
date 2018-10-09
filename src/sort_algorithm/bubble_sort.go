/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
    Blog   : gaowenhao.cn
-----------------------------------------------------
*/

package sort_algorithm

import "tools/arraytools"

/*
	经典冒泡排序算法,这个很简单,每次外层遍历就是靠着两两对比,把最大的值冒上去,下一次外层循环忽略掉最后一个值然后便找出了第二大的,如此循环.
*/
func BubbleSort(array []int) {
	length := len(array)

	for x := 0; x < length; x++ {
		for y := 0; y < (length - x - 1); y++ {
			if array[y] > array[y+1] {
				arraytools.Swap(array, y, y+1)
			}
		}
	}
}
