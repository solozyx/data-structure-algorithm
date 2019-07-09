package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	// 数组存储队列数据
	array [5]int
	// 队头元素下标 - 1
	front int
	// 队尾元素下标
	rear int
}

func (this *Queue) AddQueue(val int) (err error) {
	// 判断队列是否已满
	// rear 是队尾元素下标 (即 rear指向队尾元素)
	if this.rear == this.maxSize-1 {
		return errors.New("queue is already full")
	}
	// rear 后移
	this.rear++
	this.array[this.rear] = val
	return
}

func (this *Queue) GetQueue() (val int, err error) {
	// 判断队列是否为空 front追到和rear相等 queue空
	if this.rear == this.front {
		return -1, errors.New("queue is empty")
	}
	this.front++
	val = this.array[this.front]
	// 这里可以不做清零操作
	this.array[this.front] = 0
	return val, err
}

// 显示队列,找到队头,然后到遍历到队尾
func (this *Queue) ShowQueue() {
	fmt.Println("队列当前元素:")
	// this.front 不包含队头元素,front是队头元素下标 - 1
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("Queue.array[%d]=%d\t", i, this.array[i])
	}
	fmt.Println()
}

func main() {
	// 创建队列
	queue := &Queue{
		maxSize: 5,
		// array是值类型,创建队列时自动初始化,所有元素零值0填充
		front: -1,
		rear:  -1,
	}

	// 初始化2个变量
	var key string
	var val int

	for {
		fmt.Println("输入: add 添加数据到队列; get 从队列获取数据; show 显示队列; exit 退出程序")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入要入队列数值:")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println("数据入队列失败 :( " + err.Error())
			} else {
				fmt.Println("数据入队列成功 :)")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("队列取出数值 = ", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
