/*
-----------------------------------------------------
    Author : 高文豪
    Github : https://github.com/gaowenhao
-----------------------------------------------------
*/

package number_theory_algorithm

import "math"

/*
	判断素数朴素方法
*/

// 获取
func GetPrimeNumberArray(val int) []int {
	var result []int
	for i := 2; i <= val; i++ {
		if IsPrimeNumber(i) {
			result = append(result, i)
		}
	}
	return result
}

// 判断一个数字是否是素数
func IsPrimeNumber(val int) bool {
	sqrt := int(math.Sqrt(float64(val))) + 1
	for x := 2; x < sqrt; x++ {
		if val%x == 0 {
			return false
		}
	}
	return true
}
