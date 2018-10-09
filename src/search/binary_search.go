/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
    Blog   : gaowenhao.cn
-----------------------------------------------------
*/

package search

/*
	二分查找,这没什么好说的O(lgn) 2为底
*/
func BinarySearch(array []int, val int) int {
	start := 0
	end := len(array) - 1

	for start <= end {
		middle := int((end-start)/2) + start
		if array[middle] > val {
			end = middle
		} else if array[middle] < val {
			start = middle + 1
		} else {
			return middle
		}
	}

	return -1
}
