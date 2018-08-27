/*
-----------------------------------------------------
    Author : 高文豪
    Github : https://github.com/gaowenhao
-----------------------------------------------------
*/

package sort_algorithm

/*
	快排也非常容易理解,先二分,然后选择一个对比值,从二分后左面数组中找出大于对比值的,右面数组中找出小于对比值的,交换.而后用分割后的两段数组递归此过程.
*/

import "arraytools"

func QuickSort(array []int, left int, right int) {
	if right < left {
		return
	}

	start := left
	end := right
	temp := array[start]

	for start != end {
		for end > start && array[end] >= temp {
			end--
		}
		for start < end && array[start] <= temp {
			start++
		}
		arraytools.Swap(array, start, end)
	}

	array[left] = array[start]
	array[start] = temp

	QuickSort(array, left, start-1)
	QuickSort(array, start+1, right)
}
