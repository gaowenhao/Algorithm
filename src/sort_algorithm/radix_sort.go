/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
    Blog   : gaowenhao.cn
-----------------------------------------------------
*/

package sort_algorithm

import (
	"math"
)

/*
	基数排序:我印象里基数排序的速度是快于快速排序的,但是在golang里却慢很多,上次印象用的是python写的程序,目前原因还并没想太好.
	这个排序很好理解,这里选择基数为10,那么也就是先让整个数组的个位有序,然后十位有序,依次类推到最后一位，最后数组则必然有序.
*/

// 默认基数
const Radix = 100

func RadixSort(array []int) {
	max := GetMaxNumber(array)

	k := int(math.Ceil(math.Log10(float64(max))))

	for i := 1; i < k+1; i++ {
		var bucket [Radix][]int

		for _, val := range array {
			position := GetPosition(val, i)
			bucket[position] = append(bucket[position], val)
		}

		// 清空切片
		array = array[:0]

		for _, val := range bucket {
			array = append(array, val...)
		}
	}
}

func GetMaxNumber(array []int) int {
	max := 0

	for _, val := range array {
		if val > max {
			max = val
		}
	}

	return max
}

// 计算是第几位的序
func GetPosition(val int, index int) int {
	return val % int(math.Pow(float64(Radix), float64(index))) / int(math.Pow(float64(Radix), float64(index-1)))
}
