/*
-----------------------------------------------------
    Author : 高文豪
    Github : https://github.com/gaowenhao
-----------------------------------------------------
*/

package correlation

import (
	"fmt"
	"math"
)

/*
	欧氏距离是,机器学习求相关性的最基础算法.

    因为数学不是自然科学,所以数学中的的空间维度可以无限,故而有N个元素的向量可以适用欧氏距离,算法很简单,公式也很简单.

	sqrt((x1 - y1)^2 + (x2 - y2)^2 + (x3 - y3)^2 ...)
*/

func EuclideanDistance(v1 []float64, v2 []float64) float64 {
	v1Length := len(v1)
	v2Length := len(v2)

	if v1Length != v2Length {
		return -1
	}

	powResult := 0.0

	for index, v1val := range v1 {
		powResult += math.Pow(float64(v1val-v2[index]), 2)
	}

	return math.Sqrt(powResult)
}

func TestEuclideanDistance() {
	v1 := []float64{3, 2, 7, 6, 7}
	v2 := []float64{4, 3, 7, 13, 7}
	fmt.Print(EuclideanDistance(v1, v2))
}
