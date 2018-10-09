/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
    Blog   : gaowenhao.cn
-----------------------------------------------------
*/

package arraytools

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成数组内容区间0-RandomBetween
const RandomBetween int = 100000

// 普通交换
func Swap(array []int, x int, y int) {
	array[x], array[y] = array[y], array[x]
}

// 异或交换
func SwapXor(array []int, x int, y int) {
	array[x] = array[x] ^ array[y]
	array[y] = array[x] ^ array[y]
	array[x] = array[y] ^ array[x]
}

// 代数交换
func SwapMath(array []int, x int, y int) {
	array[x] = array[x] + array[y]
	array[y] = array[x] - array[y]
	array[x] = array[x] - array[y]
}

// 生成数组
func GenerateArray(length int) []int {
	var array []int
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; {
		array = append(array, rand.Intn(RandomBetween))
		i++
	}
	return array
}

// 打印数组
func PrintArray(array []int) {
	length := len(array)
	if length < 20 {
		fmt.Println("前20个数字: ", array[:])
		fmt.Println("后20个数字: ", array[:])
		fmt.Println()
	} else {
		fmt.Println("前20个数字: ", array[0:20])
		fmt.Println("后20个数字: ", array[len(array)-20:])
		fmt.Println()
	}

}
