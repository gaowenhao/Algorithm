/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
	Blog   : gaowenhao.cn
-----------------------------------------------------
*/

/*
	埃拉托色尼筛选法,简称埃氏筛法,思路：产生一个数组，与指定数字从0到n相互对应,先定所有数都为素数,0,1,发现一个素数后标记所有他的倍数也为素数
*/

package number_theory_algorithm

import (
	"fmt"
	"time"
)

func Erato(max int) []int {
	boolArray := MakeTrueArray(max + 1)

	var result []int
	boolArray[0] = false
	boolArray[1] = false

	for i := 2; i < max; i++ {
		if !boolArray[i] {
			continue
		} else {
			result = append(result, i)
			for j := i * i; j < max; j += i {
				boolArray[j] = false
			}
		}
	}
	return result
}

func MakeTrueArray(max int) []bool {
	boolArray := make([]bool, max)

	for index, _ := range boolArray {
		boolArray[index] = true
	}
	return boolArray
}

func TestErato() {
	start := time.Now().UnixNano() / 1e6
	fmt.Print(Erato(1000000))
	fmt.Println()
	end := time.Now().UnixNano() / 1e6

	fmt.Println("耗时 : ", end-start, "ms")
}
