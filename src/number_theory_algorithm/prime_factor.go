/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
	Blog   : gaowenhao.cn
-----------------------------------------------------
*/

package number_theory_algorithm

/*
	分解质因数,这个很简单,从2往上判断是否%即可,如果已经是遍历过质数的倍数(合数)那么就会被自动剔除.
*/

func PrimeFactor(val int) []int {
	num := 2
	var result []int

	for num <= val {
		if val%num == 0 {
			result = append(result, num)
			val /= num
		} else {
			num++
		}
	}
	return result
}
