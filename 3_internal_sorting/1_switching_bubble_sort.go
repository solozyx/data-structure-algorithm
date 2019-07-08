package main

import (
	"fmt"
)

// 交换式排序 - 冒泡排序

func BubbleSort(arr *[5]int) {
	fmt.Println("排序前arr=", (*arr))
	// 临时变量(用于做交换)
	temp := 0
	// 冒泡排序
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			// 从小到大
			if (*arr)[j] > (*arr)[j+1] {
				// 交换
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后arr=", (*arr))
}

func main() {
	// 定义数组
	arr := [5]int{24, 69, 80, 57, 13}
	// 将数组指针传递给冒泡函数,完成排序
	BubbleSort(&arr)
	fmt.Println("main arr = ", arr) // 有序
}

/*
排序前arr= [24 69 80 57 13]
排序后arr= [13 24 57 69 80]
main arr =  [13 24 57 69 80]

Process finished with exit code 0
*/
