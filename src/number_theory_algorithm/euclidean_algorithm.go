/*
-----------------------------------------------------
    Author : 高文豪
    Github : https://github.com/gaowenhao
-----------------------------------------------------
*/
package number_theory_algorithm

import "fmt"

/*
	欧几里德碾除法,最最大公约数.
*/

func Gcd(num1 int, num2 int) int {
	if num2 == 0 {
		return num1
	} else {
		return Gcd(num2, num1%num2)
	}
}

func TestGcd() {
	fmt.Print(Gcd(20, 6))
}
