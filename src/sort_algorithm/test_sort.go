/*
-----------------------------------------------------
    Author : 高文豪
    Github : github.com/gaowenhao
	Blog   : gaowenhao.cn
-----------------------------------------------------
*/

package sort_algorithm

/*
测试快速排序:
耗时 :  8 ms
--------------------------------------------------------------------------------------------------

测试归并排序:
耗时 :  18 ms
--------------------------------------------------------------------------------------------------

测试堆排序:
耗时 :  18 ms
--------------------------------------------------------------------------------------------------

测试希尔排序:
耗时 :  18 ms
--------------------------------------------------------------------------------------------------

测试基数排序:
耗时 :  43 ms
--------------------------------------------------------------------------------------------------

测试插入排序:
耗时 :  1877 ms
--------------------------------------------------------------------------------------------------

测试鸡尾酒排序:
耗时 :  9051 ms
--------------------------------------------------------------------------------------------------


测试冒泡排序:
耗时 :  15713 ms
--------------------------------------------------------------------------------------------------

测试选择排序:
耗时 :  15520 ms
--------------------------------------------------------------------------------------------------
*/

import (
	"fmt"
	"time"
	"tools/arraytools"
)

const Print bool = true

func TestBubbleSort(length int) {
	fmt.Println("测试冒泡排序:")

	array := arraytools.GenerateArray(length)

	if Print {
		fmt.Println("排序前:")
		arraytools.PrintArray(array)
	}

	start := time.Now().UnixNano() / 1e6

	BubbleSort(array)

	end := time.Now().UnixNano() / 1e6

	if Print {
		fmt.Println("排序后:")
		arraytools.PrintArray(array)
	}

	fmt.Println("耗时 : ", end-start, "ms")

	fmt.Println("--------------------------------------------------------------------------------------------------")

	fmt.Println()
}

func TestSelectionSort(length int) {
	fmt.Println("测试选择排序:")

	array := arraytools.GenerateArray(length)

	if Print {
		fmt.Println("排序前:")
		arraytools.PrintArray(array)
	}

	start := time.Now().UnixNano() / 1e6

	SelectionSort(array)

	end := time.Now().UnixNano() / 1e6

	if Print {
		fmt.Println("排序后:")
		arraytools.PrintArray(array)
	}

	fmt.Println("耗时 : ", end-start, "ms")

	fmt.Println("--------------------------------------------------------------------------------------------------")

	fmt.Println()
}

func TestQuickSort(length int) {
	fmt.Println("测试快速排序:")

	array := arraytools.GenerateArray(length)

	if Print {
		fmt.Println("排序前:")
		arraytools.PrintArray(array)
	}

	start := time.Now().UnixNano() / 1e6

	QuickSort(array, 0, length-1)

	end := time.Now().UnixNano() / 1e6

	if Print {
		fmt.Println("排序后:")
		arraytools.PrintArray(array)
	}

	fmt.Println("耗时 : ", end-start, "ms")

	fmt.Println("--------------------------------------------------------------------------------------------------")

	fmt.Println()
}

func TestMergeSort(length int) {
	fmt.Println("测试归并排序:")

	array := arraytools.GenerateArray(length)

	if Print {
		fmt.Println("排序前:")
		arraytools.PrintArray(array)
	}

	start := time.Now().UnixNano() / 1e6

	MergeSort(array, 0, length-1)

	end := time.Now().UnixNano() / 1e6

	if Print {
		fmt.Println("排序后:")
		arraytools.PrintArray(array)
	}

	fmt.Println("耗时 : ", end-start, "ms")

	fmt.Println("--------------------------------------------------------------------------------------------------")

	fmt.Println()
}

func TestMergeSort2(length int) {
	fmt.Println("测试归并排序:")

	array := arraytools.GenerateArray(length)

	if Print {
		fmt.Println("排序前:")
		arraytools.PrintArray(array)
	}

	start := time.Now().UnixNano() / 1e6

	MergeSort2(array, 0, length-1)

	end := time.Now().UnixNano() / 1e6

	if Print {
		fmt.Println("排序后:")
		arraytools.PrintArray(array)
	}

	fmt.Println("耗时 : ", end-start, "ms")

	fmt.Println("--------------------------------------------------------------------------------------------------")

	fmt.Println()
}

func TestInsertionSort(length int) {
	fmt.Println("测试插入排序:")

	array := arraytools.GenerateArray(length)

	if Print {
		fmt.Println("排序前:")
		arraytools.PrintArray(array)
	}

	start := time.Now().UnixNano() / 1e6

	InsertionSort(array)

	end := time.Now().UnixNano() / 1e6

	if Print {
		fmt.Println("排序后:")
		arraytools.PrintArray(array)
	}

	fmt.Println("耗时 : ", end-start, "ms")

	fmt.Println("--------------------------------------------------------------------------------------------------")

	fmt.Println()
}

func TestShellSort(length int) {
	fmt.Println("测试希尔排序:")

	array := arraytools.GenerateArray(length)

	if Print {
		fmt.Println("排序前:")
		arraytools.PrintArray(array)
	}

	start := time.Now().UnixNano() / 1e6

	ShellSort(array)

	end := time.Now().UnixNano() / 1e6

	if Print {
		fmt.Println("排序后:")
		arraytools.PrintArray(array)
	}

	fmt.Println("耗时 : ", end-start, "ms")

	fmt.Println("--------------------------------------------------------------------------------------------------")

	fmt.Println()
}

func TestRadixSort(length int) {
	fmt.Println("测试基数排序:")

	array := arraytools.GenerateArray(length)

	if Print {
		fmt.Println("排序前:")
		arraytools.PrintArray(array)
	}

	start := time.Now().UnixNano() / 1e6

	RadixSort(array)

	end := time.Now().UnixNano() / 1e6

	if Print {
		fmt.Println("排序后:")
		arraytools.PrintArray(array)
	}

	fmt.Println("耗时 : ", end-start, "ms")

	fmt.Println("--------------------------------------------------------------------------------------------------")

	fmt.Println()
}

func TestHeapSort(length int) {
	fmt.Println("测试堆排序:")

	array := arraytools.GenerateArray(length)

	if Print {
		fmt.Println("排序前:")
		arraytools.PrintArray(array)
	}

	start := time.Now().UnixNano() / 1e6

	BuildMaxHeap(array, length)

	HeapSort(array)

	end := time.Now().UnixNano() / 1e6

	if Print {
		fmt.Println("排序后:")
		arraytools.PrintArray(array)
	}

	fmt.Println("耗时 : ", end-start, "ms")

	fmt.Println("--------------------------------------------------------------------------------------------------")

	fmt.Println()
}

func TestCocktailSort(length int) {
	fmt.Println("测试鸡尾酒排序:")

	array := arraytools.GenerateArray(length)

	if Print {
		fmt.Println("排序前:")
		arraytools.PrintArray(array)
	}

	start := time.Now().UnixNano() / 1e6

	CocktailSort(array)

	end := time.Now().UnixNano() / 1e6

	if Print {
		fmt.Println("排序后:")
		arraytools.PrintArray(array)
	}

	fmt.Println("耗时 : ", end-start, "ms")

	fmt.Println("--------------------------------------------------------------------------------------------------")

	fmt.Println()
}

func TestMonkeySort(length int) {
	fmt.Println("测试猴子排序:")

	array := arraytools.GenerateArray(length)

	if Print {
		fmt.Println("排序前:")
		arraytools.PrintArray(array)
	}

	start := time.Now().UnixNano() / 1e6

	MonkeySort(array)

	end := time.Now().UnixNano() / 1e6

	if Print {
		fmt.Println("排序后:")
		arraytools.PrintArray(array)
	}

	fmt.Println("耗时 : ", end-start, "ms")

	fmt.Println("--------------------------------------------------------------------------------------------------")

	fmt.Println()
}

// 测试所有的数组
func TestAllSort(length int) {
	TestQuickSort(length)
	TestMergeSort(length)
	TestHeapSort(length)
	TestShellSort(length)
	TestRadixSort(length)
	TestInsertionSort(length)
	TestCocktailSort(length)
	TestBubbleSort(length)
	TestSelectionSort(length)
}
