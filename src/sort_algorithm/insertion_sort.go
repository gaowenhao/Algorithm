/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
	Blog   : gaowenhao.cn
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
		// 指向当前元素
		temp := array[x]
		y := x
		// y下标递减,如果下标为y-1的值大于当前元素, 则把y-1位置的元素,给y
		for ; y > 0 && array[y-1] > temp; y-- {
			array[y] = array[y-1]
		}
		// 跳出循环后, 插入最开始指向的当前元素
		array[y] = temp
	}
}
