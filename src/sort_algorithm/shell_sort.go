/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
	Blog   : gaowenhao.cn
-----------------------------------------------------
*/

package sort_algorithm

/*
	希尔排序,选定一个gap值,gap为第一遍历,用这个值去跟0位值比,如果小于则调换,下一次内循环因为y>=0所以会直接退出,
	第二次的值会跟数组前半段第二个比,当如此循环一个外遍历后,相距离gap那么远的值就是有序的,接下来gap /=2缩小gap值,如此循环.

	[5,3,4,2]
	[4,3,5,2]
	[4,2,5,3]
	[2,4,5,3]
	[2,3,4,5]
*/

func ShellSort(array []int) {
	length := len(array)
	gap := length >> 1

	for 1 <= gap {
		for x := gap; x < length; x++ {
			temp := array[x]
			y := x - gap
			for y >= 0 && array[y] > temp {
				array[x] = array[y]
				y -= gap
			}
			array[y+gap] = temp
		}
		gap = gap >> 1
	}
}
