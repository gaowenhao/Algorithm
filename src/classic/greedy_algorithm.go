/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
	Blog   : gaowenhao.cn
-----------------------------------------------------
*/

package classic

import (
	"fmt"
	"strconv"
)

/*
	贪心算法模拟找钱
*/

// 面值,也不知道FaceValue这个词用的对不,总之这里不能用const很难受.
var FaceValue = [6]int{1, 5, 10, 20, 50, 100}

func Greedy(mount int) {
	var result string

Loop:
	for mount != 0 {
		for i := 5; i >= 0; i-- {
			if mount/FaceValue[i] >= 1 {
				count := int(mount / FaceValue[i])
				mount = mount - count*FaceValue[i]
				result += "需要" + strconv.Itoa(count) + "张 " + strconv.Itoa(FaceValue[i]) + "块 \n"
				continue Loop
			}
		}
	}

	fmt.Print(result)
}
