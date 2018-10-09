/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
	Blog   : gaowenhao.cn
-----------------------------------------------------
*/

package sort_algorithm

/*
	这里面提供了两种归并方式,上面那种是算法导论上的写法,下面的一种快一些,区别在第二个归并的注释中有写到.
	归并排序首先先用二分法分割数组,力度到个上之后,使用一个临时数组使分割后的两端数组有序入列,当全部入列后,在把临时数组内容归到原始数组中.
*/

const MaxInt = int(^uint32(0) >> 1)

// 算法导论上的经典写法
func MergeSort(array []int, start int, end int) {
	if start < end {
		middle := (start + end) / 2
		MergeSort(array, start, middle)
		MergeSort(array, middle+1, end)
		Merge(array, start, middle, end)
	}
}

func Merge(array []int, start int, middle int, end int) {
	leftLength := middle - start + 1
	rightLength := end - middle

	leftArray := make([]int, leftLength+1)
	rightArray := make([]int, rightLength+1)

	for i := 0; i < rightLength; i++ {
		rightArray[i] = array[middle+i+1]
	}
	rightArray[rightLength] = MaxInt

	for i := 0; i < leftLength; i++ {
		leftArray[i] = array[start+i]
	}

	leftArray[leftLength] = MaxInt

	x := 0
	y := 0

	// 归并进大数组
	for k := start; k < end+1; k++ {
		if leftArray[x] <= rightArray[y] {
			array[k] = leftArray[x]
			x++
		} else {
			array[k] = rightArray[y]
			y++
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func MergeSort2(array []int, start int, end int) {
	if start < end {
		middle := (start + end) / 2
		MergeSort2(array, start, middle)
		MergeSort2(array, middle+1, end)
		Merge2(array, start, middle, end)
	}
}

/*
	自己理解的一种写法（这种方法比上面的更快些,而且不需要生成那个相对大数,在我眼里更容易理解一些.）
	导论上的写法会构建两个数组,先把大数组内容有序地甩进临时数组内，然后把两个临时数组再归并到大数组内，
	而我这里的写法只创建一个临时数组，用以排序，排序后将这个临时数组的内容归并回去。
	直观地讲，空间复杂度基本是一样大的，我这边的做法在复制过程中临时数组就有序了，归并就是单纯地复制，而上面的做法是复制过程中仅仅是复制，归并过程会有排序
*/
func Merge2(array []int, start int, middle int, end int) {
	tempLength := end - start + 1
	tempArray := make([]int, tempLength)
	left := start
	right := middle + 1
	cursor := 0

	// 归并的排序过程
	for left <= middle && right <= end {
		if array[left] < array[right] {
			tempArray[cursor] = array[left]
			left++
		} else {
			tempArray[cursor] = array[right]
			right++
		}
		cursor++
	}

	// 归并过程
	for left <= middle {
		tempArray[cursor] = array[left]
		cursor++
		left++
	}

	// 归并过程
	for right <= end {
		tempArray[cursor] = array[right]
		cursor++
		right++
	}

	for i := 0; i < tempLength; i++ {
		array[start+i] = tempArray[i]
	}
}
