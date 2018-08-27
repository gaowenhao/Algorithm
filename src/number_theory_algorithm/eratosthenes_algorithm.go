/*
-----------------------------------------------------
    Author : 高文豪
    Github : https://github.com/gaowenhao
-----------------------------------------------------
*/

/*
	埃拉托色尼筛选法,简称埃氏筛法,写了两种方式在1000000 和 10000000 两个数值上筛选，我的那种方式快一些,我感觉好理解,
	但是我那种方式求素函数不太好,但却还是比传统方法快.

	这个地方有一个数学概念非常重要,温习一下:
		我们定义一个数他是 num 且 num > 2,则num必定有一个素数因子小于等于 int(sqrt(num)).具体的定义我忘了,我这里说的是个大概.
*/

package number_theory_algorithm

import (
	"fmt"
	"math"
	"time"
)

/*
	我理解的一种方式,和原版做个对比用.
*/
func Erato(max int) []int {
	var primeArray = []int{2}
	var result = []int{2}
	sqrt := int(math.Sqrt(float64(max)))

Loop:
	for i := 3; i < max; i += 2 {
		for _, primeNumber := range primeArray {
			if i%primeNumber == 0 {
				continue Loop
			} else {
				continue
			}
		}
		result = append(result, i)
		if i > sqrt {
			continue
		} else {
			primeArray = append(primeArray, i)
		}

	}
	return result
}

func Erato2(max int) []int {
	sqrt := int(math.Sqrt(float64(max)))
	primeArray := GetPrimeNumberArray(sqrt)
	var result []int

Loop:
	for i := 3; i < max; i++ {
		for _, primerNumber := range primeArray {
			if i%primerNumber == 0 {
				continue Loop
			} else {
				continue
			}
		}
		result = append(result, i)
	}

	result = append(primeArray, result[:]...)

	return result
}

// 返回0-n的数组
func MakeArray(max int) []int {
	var result []int

	for i := 2; i < max; i++ {
		result = append(result, i)
	}
	return result
}

func GetPrimeNumberArray(sqrt int) []int {
	var result []int
	for i := 2; i <= sqrt; i++ {
		if IsPrimeNumber(i) {
			result = append(result, i)
		}
	}
	return result
}

func IsPrimeNumber(val int) bool {
	for x := 2; x < val; x++ {
		if val%x == 0 {
			return false
		}
	}
	return true
}

func TestErato() {
	start := time.Now().UnixNano() / 1e6
	fmt.Print(Erato2(100))
	fmt.Println()
	end := time.Now().UnixNano() / 1e6

	fmt.Println("耗时 : ", end-start, "ms")
}
