package main

import (
	"errors"
	"fmt"
	"os"
)

// 环形队列
type CircleQueue struct {
	// 队列元素容量,实际环形队列最大长度是 maxSize-1，为方便程序实现预留tail标志位不存储数据
	maxSize int
	// 数组存储队列元素
	array [5]int
	// 队头元素下标 初始化为0
	head int
	// 队尾元素下标+1  初始化为0
	tail int
}

// 添加元素进队列
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue is already full")
	}
	// tail 是 [队尾元素下标 + 1] 把新值直接存储在此
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

// 从队列弹出元素
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	// 弹出队头 head 是队头元素下标
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}

// 遍历打印队列
func (this *CircleQueue) ListQueue() {
	fmt.Println("环形队列: ")
	size := this.Size()
	if size == 0 {
		fmt.Println("queue is empty")
	}
	// 辅助变量指向 this.head 队头不要动
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("CircleQueue.array[%d] = %d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

// 判断环形队列是否满
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

// 判断环形队列是否空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

// 当前环形队列长度
func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func main() {
	// 初始化1个环形队列
	queue := &CircleQueue{
		maxSize: 5,
		// array是值类型,创建队列时自动初始化,所有元素零值0填充
		head: 0,
		tail: 0,
	}

	var key string
	var val int
	for {
		fmt.Println("输入: push 添加数据到队列; pop 从队列弹出数据; list 显示队列; exit 退出程序")
		fmt.Scanln(&key)
		switch key {
		case "push":
			fmt.Println("输入要 push 进队列数值:")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("新元素push进队列成功:)")
			}
		case "pop":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列弹出数据 = ", val)
			}
		case "list":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
