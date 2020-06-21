package main

import (
	"fmt"
	"container/list"
)

//描述：
//从尾到头反过来打印出每个结点的值。


type ListNode struct {
	data int
	next *ListNode

}

//尾插法创建呆头节点链表
func createLinkedList(data []int)*ListNode {
	var headNode *ListNode = &ListNode{-1,nil}
	var p *ListNode
	for _,e :=range data{
		newNode := &ListNode{e,nil}
		if headNode.next == nil{
			headNode.next = newNode
		}else{
			p.next = newNode
		}
		p = newNode
	}
	return headNode
}

//使用递归

func printReverselyWithRecursiveWay(node *ListNode){
	if node != nil{
		printReverselyWithRecursiveWay(node.next)
		fmt.Printf("%d,",node.data)
	}
}

// 使用头插入法构建逆序列表

func printReverselyWithReverseList(originalListNode *ListNode){
	listNode := originalListNode
	headNode := ListNode{-1,nil}
	for listNode != nil{
		//首先要将下一个节点的地址保存起来
		p := listNode.next
		listNode.next = headNode.next
		headNode.next = listNode
		listNode = p
	}
	// print
	p := headNode.next
	for p != nil{
		fmt.Printf("%v,",p.data)
		p = p.next
	}
}

//使用栈

func printReverselyWithStackWay(listNode *ListNode){
	myList := list.New()
	for listNode != nil{
		myList.PushBack(listNode.data)
		listNode = listNode.next
	}
	for e:= myList.Back();e!=nil;e=e.Prev(){
		fmt.Printf("%v,",e.Value)
	}
}

//方法2,3在测试时候需要注释掉一个，因为任何一个方法操作完后，原来的指针位置发生了变化
func main(){
	A := []int{1,2,3,4,5}
	linkedNode := createLinkedList(A)
	//1递归法打印
	fmt.Println("RecursiveWay")
	printReverselyWithRecursiveWay(linkedNode.next)
	fmt.Println()
	//2逆序链表打印
	//fmt.Println("ReverseList")
	//printReverselyWithReverseList(linkedNode.next)
	//fmt.Println()
	//3栈方式打印
	fmt.Println("StackWay")
	printReverselyWithStackWay(linkedNode.next)
}

