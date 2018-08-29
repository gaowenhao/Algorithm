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
	"tools/mathtools"
)

/*
	皮尔逊相关系数
*/

func Pearson(v1 []float64, v2 []float64) float64 {
	v1Length := len(v1)
	v2Length := len(v2)

	if v1Length != v2Length {
		return -1
	}

	sum1 := mathtools.Sum(v1)
	sum2 := mathtools.Sum(v2)

	sumPow1 := mathtools.SumPow(v1, 2)
	sumPow2 := mathtools.SumPow(v2, 2)

	sumPlus := SumPlus(v1, v2)

	num := sumPlus - sum1*sum2/float64(v1Length)
	den := math.Sqrt((sumPow1 - math.Pow(sum1, 2)/float64(v1Length)) * (sumPow2 - math.Pow(sum2, 2)/float64(v1Length)))

	return num / den
}

func SumPlus(v1 []float64, v2 []float64) float64 {
	result := 0.0

	for index, val := range v1 {
		result += val * v2[index]
	}

	return result
}

func TestPearson() {
	v1 := []float64{10, 1, 5}
	v2 := []float64{10, 3, 5}
	fmt.Print(Pearson(v1, v2))
}
