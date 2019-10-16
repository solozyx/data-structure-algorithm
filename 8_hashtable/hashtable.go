// 1.使用链表实现哈希表
// 2.链表没有表头,即链表第1个节点就存放雇员信息
// TODO:ERROE 第1个节点插入id=12 此时插入 id<12 的新节点会把链表破坏掉有漏洞
package main

import (
	"fmt"
	"os"
)

// 雇员
type Emp struct {
	Id   int
	Name string
	// 指向下一个Emp节点的指针
	Next *Emp
}

// EmpLink 没有表头 即第1个节点就存放雇员数据
type EmpLink struct {
	// 头指针
	Head *Emp
}

// 哈希表 hashtable 含有1个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

func (e *Emp) ShowMe() {
	fmt.Printf("链表 %d 找到该雇员 %d \n", e.Id%7, e.Id)
}

// 添加员工 保证添加时员工编号从小到大
func (l *EmpLink) Insert(emp *Emp) {
	// 把当前链表的头指针 赋值给 辅助指针
	cur := l.Head
	// 辅助指针 pre 在 cur 前面
	var pre *Emp = nil
	// 如果当前的 EmpLink 是空链表
	if cur == nil {
		l.Head = emp
		return
	}
	// 如果当前的 EmpLink 不是空链表 给 emp 找到对应的位置并插入
	for {
		if cur != nil {
			// 让 cur 和 emp 比较  >= 允许有编号相同的节点同时存在
			if cur.Id >= emp.Id {
				// 找到位置 插入到 cur 前
				break
			}
			// 然后让 pre 保持在 cur 前面
			pre = cur
			cur = cur.Next
		} else {
			// cur 已经找到链表末尾 则把 emp 插入到链表末尾
			break
		}
	}
	// 退出循环 将 emp 添加到链表
	pre.Next = emp
	emp.Next = cur
}

// 显示链表元素
func (l *EmpLink) ShowLink(linkNo int) {
	if l.Head == nil {
		fmt.Printf("链表 %d 为空\n", linkNo)
		return
	}
	// 辅助指针
	cur := l.Head
	for {
		if cur != nil {
			fmt.Printf("链表 %d 雇员id = %d 名字 = %s -->", linkNo, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

func (l *EmpLink) FindById(id int) *Emp {
	cur := l.Head
	// 假设链表存储的雇员id唯一
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

// 给 HashTable 编写 Insert 雇员的方法
func (h *HashTable) Insert(emp *Emp) {
	// 使用散列函数 id name id+name 不同方式进行散列 确定将该雇员添加到哪个链表
	linkNo := h.HashFun(emp.Id)
	// 使用对应的链表添加 emp
	h.LinkArr[linkNo].Insert(emp)
}

// 显示 hashtable 的所有雇员
func (h *HashTable) ShowAll() {
	for i := 0; i < len(h.LinkArr); i++ {
		h.LinkArr[i].ShowLink(i)
	}
}

// 散列 得到对于链表的下标
func (h *HashTable) HashFun(id int) int {
	// 扩展 : 二级链表 取模 再 取模 id % 7 % x
	return id % 7
}

func (h *HashTable) FindById(id int) *Emp {
	// 使用散列函数 确定将该雇员所在哪个链表
	linkNo := h.HashFun(id)
	return h.LinkArr[linkNo].FindById(id)
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashtable HashTable

	for {
		fmt.Println("===============雇员系统菜单============")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show  表示显示雇员")
		fmt.Println("find  表示查找雇员")
		fmt.Println("exit  表示退出系统")

		fmt.Println("请输入你的操作")
		fmt.Scanln(&key)

		switch key {
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{Id: id, Name: name}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Println("请输入id号:")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d 的雇员不存在\n", id)
			} else {
				emp.ShowMe()
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}
	}
}
