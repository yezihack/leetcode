/*
 * @Author: 百里
 * @Date: 2020-02-10 22:49:50
 * @LastEditTime : 2020-02-11 16:17:52
 * @LastEditors  : 百里
 * @Description:
 * @FilePath: \leetcode\02.链表\01.单链表.go
 * @https://github.com/yezihack
 */
package linked

type ISingleLinked interface {
	AddHead(data int) //插入头节点
	Append(data int)  //追加节点,即添加到尾部.
	RemoveHead() bool //删除头节点
	RemoveTail() bool //删除尾节点
}

