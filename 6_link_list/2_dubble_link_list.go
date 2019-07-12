package main

import (
	"fmt"
)

type Node struct {
	no       int
	name     string
	nickname string
	// 指向前1个节点
	pre *Node
	// 指向下1个节点
	next *Node
}

// 第1种插入,在双向链表尾部插入新数据
func PushNodeOnTail(head *Node, newNode *Node) {
	// 1.先找到链表尾部节点
	// 2.head节点不能动,创建1个辅助节点
	temp := head
	for {
		// 定位到链表尾部
		if temp.next == nil {
			break
		}
		// 让temp不断指向下1个节点
		temp = temp.next
	}
	// 3.将newNode插入到链表最后
	temp.next = newNode
	newNode.pre = temp
}

// 第2种插入,给双向链表插入1个节点,根据no编号从小到大插入
func PushNodeOrder(head *Node, newNode *Node) {
	// 1.找到适当节点
	// 2.head节点不能动,创建1个辅助节点
	temp := head
	flag := true
	// 让待插入节点.no 和 temp.next.no 比较
	for {
		// 到链表尾部
		if temp.next == nil {
			break
		} else if temp.next.no > newNode.no {
			// newNode 应该插入到temp后面
			break
		} else if temp.next.no == newNode.no {
			// 在链表中已经有这个no,就不能重复插入
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("the link list is already exists no = ", newNode.no)
		return
	} else {
		newNode.next = temp.next
		newNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newNode
		}
		temp.next = newNode
	}
}

// 双向链表删除1个节点
func DelNode(head *Node, id int) {
	temp := head
	flag := false
	// 找到要删除节点.no 和 temp.next.no 比较
	for {
		// 到链表尾部节点
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			// 找到
			flag = true
			break
		}
		temp = temp.next
	}
	// 删除
	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Printf("link list not exists id = %d to delete. \n", id)
	}
}

// 这里仍然使用单向链表的遍历方式
func ListNodeOrder(head *Node) {
	// 1.head节点不能动,创建1个辅助节点
	temp := head
	// 先判断该链表是否空链表
	if temp.next == nil {
		fmt.Println("the link list is empty :(")
		return
	}
	// 2.遍历链表
	for {
		fmt.Printf("[%d , %s , %s]==>", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		// 判断是否链表尾部
		if temp.next == nil {
			break
		}
	}
}

// 逆序遍历链表,证明链表是双向的
func ListNodeReverse(head *Node) {
	// 1.head节点不能动,创建1个辅助节点
	temp := head
	// 先判断该链表是否空链表
	if temp.next == nil {
		fmt.Println("the link list is empty :(")
		return
	}
	// 2.让temp定位到双向链表尾部节点
	for {
		if temp.next == nil {
			// 找到链表尾部则跳出for循环
			break
		}
		temp = temp.next
	}

	//3.逆序遍历链表
	for {
		// do-while循环,打出尾部节点,注意 temp.next.no 则空指针错误
		fmt.Printf("[%d , %s , %s]==>", temp.no, temp.name, temp.nickname)
		temp = temp.pre
		// 判断是否链表头部节点
		if temp.pre == nil {
			// 不跳出for则空指针错误
			break
		}
	}
}

func main() {
	// 1.创建1个头结点
	head := &Node{}

	// 2.创建1个新Node
	hero1 := &Node{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &Node{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero3 := &Node{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}

	//PushNodeOnTail(head, hero1)
	//PushNodeOnTail(head, hero2)
	//PushNodeOnTail(head, hero3)

	PushNodeOrder(head, hero3)
	PushNodeOrder(head, hero1)
	PushNodeOrder(head, hero2)

	fmt.Println("顺序遍历链表:")
	ListNodeOrder(head)
	fmt.Println()
	fmt.Println("逆序遍历链表:")
	ListNodeReverse(head)

	DelNode(head, 3)
	fmt.Println()
	fmt.Println("顺序遍历链表:")
	ListNodeOrder(head)
}
