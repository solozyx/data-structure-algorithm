package main

import (
	"fmt"
)

// 二分查找
// 如要查找的数是 findVal
// 1. arr是一个有序数组,从小到大排序
// 2. 先找到 中间的下标 middle = (leftIndex + rightIndex) / 2, 中间下标的值和findVal进行比较
// 2.1 如果 arr[middle] > findVal ,应向 arr[leftIndex] --- arr[middle - 1]
// 2.2 如果 arr[middle] < findVal ,应向 arr[middel + 1] ---- arr[rightIndex]
// 2.3 如果 arr[middle] == findVal 找到
// 2.4 上面的 2.1 2.2 2.3 逻辑会递归执行
// 3.[分析出退出递归的条件]
// if  leftIndex > rightIndex { return .. // 找不到.. }

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	// 判断leftIndex 是否大于 rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到 findVal")
		return
	}

	// 中间下标
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		// 要查找的数应在  leftIndex --- middel-1
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		// 要查找的数应在  middel+1 --- rightIndex
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		// 找到
		fmt.Printf("找到findVal,下标为%v \n", middle)
	}
}

func main() {
	arr := [6]int{1, 8, 10, 89, 1000, 1234}
	BinaryFind(&arr, 0, len(arr)-1, 10)
}

/*
找到findVal,下标为2

Process finished with exit code 0
*/
