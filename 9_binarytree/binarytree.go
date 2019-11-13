package main

import (
	"fmt"
)

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

// 前序遍历 [root节点 -> 左子树 -> 右子树]
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		// 递归
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

// 中序遍历 [root左子树 -> root -> root右子树]
func InfixOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}

func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
	}
}

func main() {
	// 构建一个二叉树
	// 根节点 left right 默认为空 没有形成二叉树
	root := &Hero{No: 1, Name: "宋江"}
	left1 := &Hero{No: 2, Name: "吴用"}

	node10 := &Hero{No: 10, Name: "tom"}
	node12 := &Hero{No: 12, Name: "jack"}
	left1.Left = node10
	left1.Right = node12

	right1 := &Hero{No: 3, Name: "卢俊义"}

	// 根节点 左边子树
	root.Left = left1
	// 根节点 右边子树
	root.Right = right1

	right2 := &Hero{No: 4, Name: "林冲"}
	right1.Right = right2

	fmt.Println("前序遍历")
	PreOrder(root)
	fmt.Println("中序遍历")
	InfixOrder(root)
	fmt.Println("后序遍历")
	PostOrder(root)
}

/*
		    1
		/       \
      2          3
   /   \          \
 10    12          4

go run binarytree.go
前序遍历
no=1 name=宋江
no=2 name=吴用
no=10 name=tom
no=12 name=jack
no=3 name=卢俊义
no=4 name=林冲

中序遍历
no=10 name=tom
no=2 name=吴用
no=12 name=jack
no=1 name=宋江
no=3 name=卢俊义
no=4 name=林冲

后序遍历
no=10 name=tom
no=12 name=jack
no=2 name=吴用
no=4 name=林冲
no=3 name=卢俊义
no=1 name=宋江
*/
