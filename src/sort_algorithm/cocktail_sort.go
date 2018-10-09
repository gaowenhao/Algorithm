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
	鸡尾酒排序,很久以前无意中看到过一次,今天搜了一下原理,大体有个了解(也不确定对不对)就是冒泡排序,
	但是方向是双向的,多了两个变量分别用于控制方向和边界,然后还要维护这两个字段,我本以为和冒泡差距不大,但是10万的数组冒泡15713ms,而鸡尾酒排序9051ms
*/

func CocktailSort(array []int) {
	bound := 0
	// true 顺序, false 逆序
	direction := true

	length := len(array)
	for i := 0; i < length/2; i++ {
		if direction {
			for j := i + bound; j < length-bound-1; j++ {
				if array[j] > array[j+1] {
					arraytools.Swap(array, j, j+1)
				}
			}
			direction = false
		} else {
			for j := length - 1 - bound; j > 0+bound; j-- {
				if array[j] < array[j-1] {
					arraytools.Swap(array, j, j-1)
				}
			}
			bound++
			direction = true
		}
	}
}
