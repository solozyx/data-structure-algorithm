package main

import (
	"fmt"
)

type HeroNode struct {
	no       int
	name     string
	nickname string
	// 指向下一个HeroNode结点
	next *HeroNode
}

// 给链表插入1个节点
// 第1种插入方法,在单链表尾部插入
// head节点非常重要,没有head节点则链表无法访问
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	// 1.先找到链表尾部结点
	// 2.head节点不能动,创建1个辅助结点
	temp := head
	for {
		// 链表尾部节点next是nil
		if temp.next == nil {
			break
		}
		// temp不断的指向下1个节点
		temp = temp.next
	}
	// 3.将newHeroNode插入到链表最后
	temp.next = newHeroNode
}

// 第2种插入,根据 HeroNode.no 编号从小到大插入
// 场景:不走数据库,没有 order by
// 先插入到数据库,再查询数据库,order by 假设10万个人 10万个insert语句 10万个select语句
// 走内存SQL语句没有,提高效率
// 现在没人在乎内存,有的是,内存不值钱
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	// 1.找到适当节点
	// 2.创建1个辅助结点
	temp := head
	// 能否插入标志 默认true能插入是大多数情况
	flag := true
	// 待插入节点.no 和 temp.next.no 比较
	// 只要有1个条件成立 则break 认为找到了插入位置 或者不能插入
	for {
		// 已经找到链表尾部,再往后走则空指针错误要break
		if temp.next == nil {
			break
			// TODO - notice
			// } else if temp.next.no >= newHeroNode.no {
		} else if temp.next.no > newHeroNode.no {
			// 说明 newHeroNode 应该插入到temp后面
			// 相同no能插入则判断条件 >= 相同no不能插入则判断条件 >
			break
		} else if temp.next.no == newHeroNode.no {
			// 说明链表中已经有这个no 不能插入
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("the link list is already exists no = ", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

// 删除1个节点
func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	// 找到要删除节点.no 和 temp.next.no 比较
	for {
		// 到链表尾部
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			// 找到
			flag = true
			break
		}
		temp = temp.next
	}
	// 找到删除
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Printf("link list not exists id = %d\n", id)
	}
}

// 显示链表所有节点
func ListHeroNode(head *HeroNode) {
	// 1.head节点不能动,创建1个辅助节点
	temp := head
	// 2.判断链表是否空链表
	if temp.next == nil {
		fmt.Println("link list is empty")
		return
	}
	//3.遍历链表
	for {
		fmt.Printf("[%d , %s , %s]==>", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		// 判断是否链表尾部
		if temp.next == nil {
			break
		}
	}
}

func main() {
	// 1.先创建一个头结点 不用赋值 默认值填充即可
	head := &HeroNode{}

	// 2. 创建一个新的HeroNode节点
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
		// next 当前节点的下一个节点要在链表结构中确定 这里无法指定
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}

	// InsertHeroNode(head,hero3)
	// InsertHeroNode(head,hero1)
	// InsertHeroNode(head,hero2)
	// ListHeroNode(head)

	hero4 := &HeroNode{
		no:       3,
		name:     "吴用",
		nickname: "智多星",
	}

	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero4)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero2)
	ListHeroNode(head)

	fmt.Println()
	DelHeroNode(head, 1)
	DelHeroNode(head, 3)
	ListHeroNode(head)
}

/*
the link list is already exists no =  3
[1 , 宋江 , 及时雨]==>[2 , 卢俊义 , 玉麒麟]==>[3 , 林冲 , 豹子头]==>
[2 , 卢俊义 , 玉麒麟]==>
Process finished with exit code 0
*/
