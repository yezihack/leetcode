/*
 * @Author: your name
 * @Date: 2020-02-18 10:59:20
 * @LastEditTime: 2020-02-18 11:16:14
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: \leetcode\13.设计\01.LRU.go
 */
package design

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */

type LRUCache struct {
	capacity int //存储大小
	cache    map[int]*LinkedNode
}
func Constructor(capacity int) LRUCache {

}

func (this *LRUCache) Get(key int) int {

}

func (this *LRUCache) Put(key int, value int) {

}

func (this *LRUCache) addHead(key, value int) {

}


type Linkeder interface {
	AddHead(key, value int) bool
	Append(key, value int) bool
	Insert(index, key, value int) bool
	RemoveTail() bool
}

type LinkedNode struct {
	key   int         //key
	value int         //value
	next  *LinkedNode //next pointer
	prev  *LinkedNode //prev pointer
}

type Linked struct {
	head *LinkedNode
	tail *LinkedNode
}

func (l *Linked) AddHead(key, value int) bool {
	var node *LinkedNode
	if l.head == nil {
		node = &LinkedNode{
			key:key,
			value:value,
		}
	}
	head := l.head
	node.next = head
	head.prev = node
	head = node


}

func (l *Linked) RemoveTail() bool {
	panic("implement me")
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
