package main

import (
	"fmt"
)

func main() {
	names := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var heroName = ""
	fmt.Println("请输入要查找的人名:")
	fmt.Scanln(&heroName)

	// 顺序查找:第1种方式
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			fmt.Printf("找到%v , 下标%v \n", heroName, i)
			break
		} else if i == (len(names) - 1) {
			fmt.Printf("没有找到%v \n", heroName)
		}
	}

	// 顺序查找:第2种方式
	fmt.Println("请输入要查找的人名:")
	fmt.Scanln(&heroName)

	index := -1

	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			// 将找到的值对应的下标赋给 index
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v , 下标%v \n", heroName, index)
	} else {
		fmt.Printf("没有找到%v \n", heroName)
	}
}

/*
请输入要查找的人名:
白眉鹰王
找到白眉鹰王 , 下标0
请输入要查找的人名:
青翼蝠王
找到青翼蝠王 , 下标3

Process finished with exit code 0
*/
