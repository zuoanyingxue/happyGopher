package main

import "fmt"

type Node struct {
	Val  int   // 节点值
	Next *Node // 下一个节点
}

func main() {
	head := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node7 := &Node{Val: 7}
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	fmt.Printf("before delete node, list is: \n")
	showListNodeInfo(head)

	newList, info := removeNthAndMthNode(head, 1, 1)

	fmt.Printf("\nafter delete node, list is: \n")
	showListNodeInfo(newList)
	fmt.Printf("info is %v\n", info)
}

func removeNthAndMthNode(head *Node, n, m int) (*Node, string) {
	tmpHead := &Node{}
	tmpHead.Next = head
	fast1 := tmpHead
	fast2 := tmpHead
	nNode := tmpHead
	mNode := tmpHead

	// 保持n值为较大数
	if m > n {
		n, m = m, n
	}

	// fast节点先走m步
	havaWalk := 0
	for ; m > 0 && fast1.Next != nil; m-- {
		fast1 = fast1.Next
	}
	if m > 0 {
		return tmpHead.Next, "m node is not exists." // 说明倒数第m个节点不存在
	}
	havaWalk = m

	// m节点、fast1节点同时走len(queue)-m 步
	for fast1.Next != nil {
		fast1 = fast1.Next
		mNode = mNode.Next

		havaWalk++
		if havaWalk <= n-1 {
			fast2 = fast2.Next
		} else if havaWalk > n {
			fast2 = fast2.Next
			nNode = nNode.Next
		}
	}
	if mNode.Next != nil {
		mNode.Next = mNode.Next.Next
	}

	if havaWalk < n {
		return tmpHead.Next, "n node is not exists." // 说明倒数第m个节点不存在
	}

	// n节点、fast2节点同时走len(queue)-n 步
	for fast2.Next != nil {
		fast2 = fast2.Next
		nNode = nNode.Next
	}

	if nNode.Next != nil {
		nNode.Next = nNode.Next.Next
	}

	return tmpHead.Next, "has deleted m node and n node, happy!"
}

func showListNodeInfo(head *Node) {
	for head != nil {
		fmt.Printf("%v->", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}
