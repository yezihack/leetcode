/*
 * @Author: 百里
 * @Date: 2020-02-10 22:49:50
 * @LastEditTime : 2020-02-11 16:17:52
 * @LastEditors  : 百里
 * @Description: 单链表操作.
 * @FilePath: \leetcode\02.链表\01.单链表.go
 * @https://github.com/yezihack
 */
package linked

import "fmt"

type SNode struct {
	data int    //数据存储
	next *SNode //后继指针
}

type ISingleLinked interface {
	AddHead(data int) *SNode     //插入头节点
	Append(data int) *SNode      //追加节点,即添加到尾部.
	RemoveHead() bool            //删除头节点
	RemoveTail() bool            //删除尾节点
	RemoveNode(node *SNode) bool //删除节点
	Print() string               //打印链表
}

//声明单链表结构体.
type SLinked struct {
	count int    //length node
	head  *SNode //head node
	tail  *SNode //tail node
}

//实例一下
func NewSLinked() *SLinked {
	return &SLinked{}
}

//添加头节点
func (s *SLinked) AddHead(data int) *SNode {
	node := &SNode{data: data} //创建一个新节点.
	if s.head == nil {         //如果头指针为空的话.
		s.head = node //node指向头指针
		s.tail = node //node也指向尾指针. 即头尾指针指向同一个节点.
	} else {
		//实现头插法.
		node.next = s.head //新的节点node后继指针指向头节点head
		s.head = node      //将node设置为新的头指针head.
	}
	s.count++
	return node
}

//追加操作.
func (s *SLinked) Append(data int) *SNode {
	node := &SNode{data: data} //创建一个新节点.
	if s.tail == nil {         //如果尾指针为空的话.
		s.head = node //node指向头指针
		s.tail = node //node也指向尾指针. 即头尾指向同一个节点.
	} else {
		//实现尾插法
		s.tail.next = node //尾指针tail的后继指针指向新节点node
		s.tail = node      //将新节点node设置为新的尾指针tail
	}
	s.count++
	return node
}

//移除头节点.
func (s *SLinked) RemoveHead() bool {
	if s.count == 0 { //如果无节点则返回False
		return false
	}
	s.head = s.head.next //将头指针的后继指针设置为新的头指针
	s.count--
	return true
}

func (s *SLinked) RemoveTail() bool {
	if s.count == 0 { //链表长度为0的话.则返回False
		return false
	} else if s.count == 1 { //链表长度为1的话,设置为空
		s.tail = nil
		s.head = nil
	} else {
		tailPre := s.head                //定义一个指针
		for i := 0; i < s.count-2; i++ { //循环count-2
			tailPre = tailPre.next
		}
		tailPre.next = nil
		s.tail = tailPre
	}
	s.count--
	return true
}

func (s *SLinked) RemoveNode(node *SNode) bool {
	if node == s.head {
		s.head = s.head.next
	} else if node == s.tail {
		tailPre := s.head
		for i := 0; i < s.count-2; i++ {
			tailPre = tailPre.next
		}
		tailPre.next = nil
		s.tail = tailPre
	} else {
		l := s.head
		var pre *SNode
		for i := 0; i < s.count-1; i++ {
			if l == node {
				break
			}
			pre = l
			l = l.next
		}
		if pre.next != nil {
			pre.next = pre.next.next
		}
	}
	return true
}

func (s *SLinked) Print() string {
	str := ""
	head := s.head
	for head != nil {
		str += fmt.Sprint(head.data)
		head = head.next
		if head != nil {
			str += "=>"
		}
	}
	return str
}
