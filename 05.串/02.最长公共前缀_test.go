package string_test

import (
	"fmt"
	"strings"
	"testing"
)

//leetcode 14. 最长公共前缀
/*
输入: ["flower","flow","flight"]
输出: "fl"

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
*/
//先找出最短的字符串, 然后再与数组里的字符串一一比较. 如果遇到不相等的则
func LongestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0 { //边界条件
		return ""
	}
	minStr := strs[0]
	//从数组里查找到最小字符串.
	for i := 1; i < l; i++ {
		if len(minStr) > len(strs[i]) { //只要存在比第一个字符还短的则进行赋值操作.
			minStr = strs[i]
		}
	}
	//如果最短字符串长度为0则,返回空
	if len(minStr) == 0 { //边界条件 .
		return ""
	}
	//因为我们只求前缀.将最短字符串比较即可.
	for i := 0; i < len(minStr); i++ { //循环最短字符串. 然后与数组里每一个字符串按挨个字符进行比较.只要遇到不相等则截取最短字符串.
		for j := 0; j < l; j++ { //循环字符串数组.
			fmt.Printf("i:%d,j:%d, str;%s, s:%c, v:%c\n", i, j, strs[j], strs[j][i], minStr[i])
			if strs[j][i] != minStr[i] { //挨个字符进行比较.
				minStr = strs[j][:i] //只要不相等,则截取数组里的字符最短前缀.
				return minStr
			}
		}
	}
	return minStr
}
func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		data   []string //数据
		expect string   //预期值
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"abcd", "bbb", "aa"}, ""},
	}
	for index, item := range tests {
		if actual := longestCommonPrefixV4(item.data); actual != item.expect {
			index++
			t.Errorf("index:%d, expect:%s, actual:%s\n", index, item.expect, actual)
		}
	}
}

func longestCommonPrefixV2(strs []string) string {
	l := len(strs)
	if l == 0 { //边界条件
		return ""
	}
	minPrefix := strs[0]             //假设第一个字符是最小字符串.
	for i := 0; i < len(strs); i++ { //遍历整个数组
		for strings.Index(strs[i], minPrefix) != 0 {
			//循环处理, 只要minPrefix不是子集则继续循环.
			// 每次对minPrefix减小一个字符进行循环比较
			if len(strs[i]) == 0 { //边界条件, 如果数组里有空的元素,则直接返回为空
				return ""
			}
			minPrefix = minPrefix[0 : len(minPrefix)-1] //每次减小一个字符进行下一轮比较.
			fmt.Printf("str:%s, min:%s\n", strs[i], minPrefix)
		}
	}
	return minPrefix
}

func TestIndex(t *testing.T) {
	str := "flower"
	fmt.Println(strings.Index("flower", str))
	fmt.Println(strings.Index("flow", str))
	fmt.Println(strings.Index("flight", str))
	fmt.Println(strings.Index("flower1", str))
}

func longestCommonPrefixV3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	min := strs[0]
	for i := 0; i < len(strs); i++ {
		for strings.Index(strs[i], min) != 0 {
			if strs[i] == "" {
				return ""
			}
			min = min[:len(min)-1]
		}
	}
	return min
}
func longestCommonPrefixV4(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	min := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(min) {
			min = strs[i]
		}
	}
	for i := 0; i < len(min); i++ {
		for j := 0; j < len(strs); j++ {
			if min[i] != strs[j][i] {
				return strs[j][:i]
			}
		}
	}
	return min
}
