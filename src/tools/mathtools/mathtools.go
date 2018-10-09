/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
	Blog   : gaowenhao.cn
-----------------------------------------------------
*/

package mathtools

/*
	一些统计学常用函数,好怀念python,就这些玩意我大python基本上都是一行撸的.
*/

import (
	"math"
)

// 求数组和
func Sum(array []float64) float64 {
	result := 0.0
	for _, val := range array {
		result += val
	}

	return result
}

// 求N次方和
func SumPow(array []float64, pow float64) float64 {
	result := 0.0
	for _, val := range array {
		result += math.Pow(val, pow)
	}
	return result
}

// 求平均值
func Average(array []float64) float64 {
	length := len(array)
	sum := 0.0
	for _, val := range array {
		sum += val
	}
	return sum / float64(length)
}

// 求标准差
func StandardDeviation(array []float64) float64 {
	length := len(array)
	avg := Average(array)
	sum := 0.0
	for _, val := range array {
		sum += math.Pow(val-avg, 2)
	}

	result := math.Sqrt(sum / float64(length))
	return result
}

// 求样本标准差
func SampleStandardDeviation(array []float64) float64 {
	length := len(array)
	avg := Average(array)
	sum := 0.0
	for _, val := range array {
		sum += math.Pow(val-avg, 2)
	}

	result := math.Sqrt(sum / (float64(length) - 1))
	return result
}

// 求方差
func Variance(array []float64) float64 {
	length := len(array)
	avg := Average(array)
	sum := 0.0
	for _, val := range array {
		sum += math.Pow(val-avg, 2)
	}

	result := sum / float64(length)
	return result
}

// 求样本方差
func SampleVariance(array []float64) float64 {
	length := len(array)
	avg := Average(array)
	sum := 0.0
	for _, val := range array {
		sum += math.Pow(val-avg, 2)
	}

	result := sum / (float64(length) - 1)
	return result
}
