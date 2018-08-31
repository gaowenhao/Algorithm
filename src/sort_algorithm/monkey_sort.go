/*
-----------------------------------------------------
    Author : 高文豪
    Github : https://github.com/gaowenhao
-----------------------------------------------------
*/

package sort_algorithm

import (
	"fmt"
	"math/rand"
	"time"
	"tools/arraytools"
)

/*
	这个算法我并不知道怎么实现,总之就是瞎实现就可以了,其实这个不一定是猴子排序,希望不要误导后来人,这不一定是猴子排序.
	至于算法复杂度,是 O(n!)? 这地方就是个排列数.
*/

func MonkeySort(array []int) {
	length := len(array)
	for index, _ := range array {
		rand.Seed(time.Now().UnixNano())
		randomVal := rand.Intn(length - 1)
		arraytools.Swap(array, index, randomVal)
	}

	if IsSorted(array) {
		fmt.Println("排好了.")
		fmt.Print(array)
	} else {
		fmt.Println("没排好,重来.")
		MonkeySort(array)
	}
}

// 判断是否有序,这个算法复杂度是O(n)
func IsSorted(array []int) bool {
	length := len(array) - 1
	for i := 0; i < length; i++ {
		if array[i] > array[i+1] {
			return false
		} else {
			continue
		}
	}

	return true
}
