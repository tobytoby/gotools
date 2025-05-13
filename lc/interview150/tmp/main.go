package main

import (
	"fmt"
)

// 插入排序函数
func insertionSort(arr []int) {
	// 从第1个元素开始，逐个将元素插入到已排序序列中
	for i := 1; i < len(arr); i++ {
		// 记录要插入的元素
		key := arr[i]
		// 在已排序序列中从后向前扫描
		j := i - 1
		// 将大于 key 的元素向后移动
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		// 将 key 插入到正确的位置
		arr[j+1] = key
	}
}

func main() {
	// 示例数组
	arr := []int{12, 11, 13, 5, 6}
	fmt.Println("排序前的数组:", arr)
	insertionSort(arr)
	fmt.Println("排序后的数组:", arr)
}
