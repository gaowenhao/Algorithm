/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
    Blog   : gaowenhao.cn
-----------------------------------------------------
*/

/*
	斐波那契序列,这个就不介绍了,这里用闭包实现的.
*/

package number_theory_algorithm

import "fmt"

func Fib() func() int {
	num1, num2 := 0, 1
	return func() int {
		num1, num2 = num2, num1+num2
		return num1
	}
}

func TestFib() {
	fun := Fib()
	for i := 0; i < 15; i++ {
		fmt.Print(fun(), ",")
	}
}
