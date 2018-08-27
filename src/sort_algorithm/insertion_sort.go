/*
-----------------------------------------------------
    Author : 高文豪
    Github : https://github.com/gaowenhao
-----------------------------------------------------
*/

package sort_algorithm

/*
	插入排序也相对容易理解,从头遍历(这也可以从尾)为每个遍历过的值找到他应该所属的位置,然后插入即可.
	1: [5,3,2,4]
	2: [3,5,2,4]
	3: [2,3,5,4]
	4: [2,3,4,5]
*/

func InsertionSort(array []int) {
	length := len(array)

	for x := 1; x < length; x++ {
		temp := array[x]
		y := x
		for ; y > 0 && array[y-1] > temp; y-- {
			array[y] = array[y-1]
		}
		array[y] = temp
	}
}
