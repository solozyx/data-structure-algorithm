package main

import (
	"errors"
	"fmt"
)

// 数组模拟栈
type Stack struct {
	// 栈容量
	MaxTop int
	// 栈顶
	Top int
	// 栈底固定, Top=-1	即表示空栈,没必要设计 Bottom成员
	// 数组存储栈元素
	arr [5]int
}

// 入栈
func (this *Stack) Push(val int) (err error) {
	// 判断是否栈满
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack is full")
		return errors.New("stack is full")
	}
	// Top初始值为-1 所以先自增
	this.Top++
	// 再插入数据
	this.arr[this.Top] = val
	return
}

// 出栈
func (this *Stack) Pop() (val int, err error) {
	// 判断是否栈空
	if this.Top == -1 {
		fmt.Println("stack is empty")
		return 0, errors.New("stack is empty")
	}
	// 先取值 再Top--
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

// 遍历栈,注意从栈顶开始遍历
func (this *Stack) List() {
	// 判断是否栈空
	if this.Top == -1 {
		fmt.Println("stack is empty")
		return
	}
	fmt.Println("栈:")
	for currTop := this.Top; currTop >= 0; currTop-- {
		fmt.Printf("Stack.arr[%d]=%d\n", currTop, this.arr[currTop])
	}
}

func main() {
	stack := &Stack{
		// 栈容量 = 5,最多存放5个数到栈中
		MaxTop: 5,
		// 栈顶 = -1 表示栈空
		Top: -1,
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	// 栈满则不能继续插入数据
	// stack.Push(6)

	stack.List()

	val, _ := stack.Pop() // 5
	fmt.Println("出栈val = ", val)
	stack.List()
	fmt.Println()

	val, _ = stack.Pop() // 4
	val, _ = stack.Pop() // 3
	val, _ = stack.Pop() // 2
	val, _ = stack.Pop() // 1

	val, _ = stack.Pop() // stack is empty

	stack.List()
}
