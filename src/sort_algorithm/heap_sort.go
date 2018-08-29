/*
-----------------------------------------------------
    Author : 高文豪
    Github : https://github.com/gaowenhao
-----------------------------------------------------
*/

package sort_algorithm

import "tools/arraytools"

func HeapSort(array []int) {
	length := len(array)

	for x := length - 1; x > 0; x-- {
		arraytools.Swap(array, 0, x)
		MaxHeapfy(array, x, 0)
	}
}

// 构建最大堆
func BuildMaxHeap(array []int, length int) {
	// >> 1 相当于除以2
	startIndex := length>>1 - 1
	for x := startIndex; x > -1; x-- {
		MaxHeapfy(array, length, x)
	}
}

func MaxHeapfy(array []int, depth int, index int) {
	leftChildIndex := index << 1
	rightChildIndex := leftChildIndex + 1

	largest := index

	if leftChildIndex < depth && array[leftChildIndex] > array[largest] {
		largest = leftChildIndex
	}

	if rightChildIndex < depth && array[rightChildIndex] > array[largest] {
		largest = rightChildIndex
	}

	if largest != index {
		arraytools.Swap(array, index, largest)
		MaxHeapfy(array, depth, largest)
	}
}
