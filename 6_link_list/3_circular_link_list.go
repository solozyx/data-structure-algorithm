package main

import (
	"fmt"
)

type CatNode struct {
	no   int
	name string
	next *CatNode
}

// 直接在尾部插入节点,不考虑no编号排序
func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	// 判断是否添加第1节点
	// 环形链表head节点要存储真实数据
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		// 构成一个环形
		// TODO - notice 没有构成环形
		// head.next = newCatNode
		head.next = head
		fmt.Println(*newCatNode, " 插入到环形链表")
		return
	}

	// 定义1个临时变量,用于找到环形链表尾部节点
	// temp指针先指向head节点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}

	// 插入到环形链表
	temp.next = newCatNode
	newCatNode.next = head
}

// 遍历环形链表,必须传入头节点,否则无法遍历
func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表:")
	// head节点不能动,辅助临时变量
	temp := head
	if temp.next == nil {
		fmt.Println("the circle link list is empty :(")
		return
	}
	// do-while 非空链表则至少能输出自己
	for {
		fmt.Printf("[id=%d name=%s] ==>", temp.no, temp.name)
		// 环形链表尾部节点
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	fmt.Println()
}

// 环形链表删除1个节点
func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head

	// 空链表
	if temp.next == nil {
		fmt.Println("the link list is empty,not allow delete node :(")
		// 这里没有改变链表head节点 也要返回
		return head
	}

	// 当链表只有1个节点
	if temp.next == head {
		if temp.no == id {
			// head自己指向自己,next指针置空即可
			temp.next = nil
			fmt.Printf("delete CatNode.no = %d\n", id)
		} else {
			fmt.Printf("not find to delete CatNode.no = %d\n", id)
		}
		// 这里没有改变链表head节点 也要返回
		return head
	}

	// 链表有 >=2 个节点
	// 将 指针helper 定位到链表尾部
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	// 链表有 >=2 个节点
	// 默认在for循环能删除节点
	flag := true
	for {
		// 说明已经追到了环形链表尾部节点 [尾部节点还没比较]
		if temp.next == head {
			// 环形链表除尾部节点全部比较完毕,退出循环
			// 这里退出循环后,需要比较链表尾部节点
			break
		}
		// 找到要删除的节点
		if temp.no == id {
			// 要删除的是头节点
			if temp == head {
				head = head.next
			}
			// 删除节点
			helper.next = temp.next
			fmt.Printf("delete CatNode.no = %d\n", id)
			// 删除节点后置为false 防止出到for循环外重复执行删除操作
			flag = false
			break
		}
		// 指针temp后移 [作用:比较]
		temp = temp.next
		// 指针helper后移 [作用:一旦找到要删除的节点 helper指针完成删除操作]
		helper = helper.next
	}

	// [尾部节点还没比较] 这里比较尾部节点
	// 如果flag为true,则上面for循环中没有删除
	if flag {
		if temp.no == id {
			// 删除链表尾部节点
			helper.next = temp.next
			fmt.Printf("delete CatNode.no = %d\n", id)
		} else {
			fmt.Printf("not find to delete CatNode.no = %d\n", id)
		}
	}

	// 链表删除head节点后,链表head节点发生改变
	// 这里返回改变后的链表head节点,供其他函数栈使用
	// 才能操作删除head节点后的此时的链表
	return head
}

func main() {
	// 初始化1个环形链表的头结点
	head := &CatNode{}

	// 创建1只猫
	cat1 := &CatNode{
		no:   1,
		name: "tom1",
		// 指针类型默认值是 nil,无需赋值
	}
	cat2 := &CatNode{no: 2, name: "tom2"}
	cat3 := &CatNode{no: 3, name: "tom3"}

	// 传入指针,引用传递,函数内部对head的操作会影响head的指向
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)

	ListCircleLink(head)

	// 删除链表不存在节点
	// head = DelCatNode(head, 30)
	// ListCircleLink(head)

	// 环形链表执行删除操作后,链表的head节点可能发生变化
	// 为防止死循环,这里接收删除操作后的head节点
	head = DelCatNode(head, 1)
	ListCircleLink(head)
}
