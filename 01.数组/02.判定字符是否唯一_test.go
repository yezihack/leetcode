package array

import (
	"fmt"
	"testing"
)
/*
实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

示例 1：
输入: s = "leetcode"
输出: false

示例 2：
输入: s = "abc"
输出: true

限制：
0 <= len(s) <= 100
如果你不使用额外的数据结构，会很加分。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/is-unique-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


//解法一: 使用ascii码,构建一个122+1的数组
//然后循环astr, 存在则加1, 再查看整个数组,是否有大于1的,则说明有重复的.
//time:O(n), space:O(n)
func IsUnique(astr string) bool {
	if len(astr) <= 1 { //边界条件
		return true
	}
	chars := make([]int, 122+1) //申请1个123长度大小的数组.
	for i := 'a'; i <= 'z'; i ++ {
		chars[i] = 0 //初使化每个元素出现0次.
	}
	for i := 0; i < len(astr); i ++ { //遍历整个字符
		chars[astr[i]] ++ //出现一次则加1
		if chars[astr[i]] > 1 { //出现出现次数大于1,则说明出现2次以上.
			return false
		}
	}
	return true
}

func TestIsUnique(t *testing.T) {
	//查看golang的A~Z之间的ASCII码值
	fmt.Println('a', 'z')//96, 122
	tests := []struct{
		data string //数据
		expect bool //预期值
	}{
		{"leetcode", false},
		{"abc", true},
		{"abcdc", false},
		{"abcdefghigka", false},
	}
	for index, item := range tests {
		if actual := IsUnique(item.data); actual != item.expect {
			index ++
			t.Errorf("index:%d, expect:%v, actual:%v\n", index, item.expect, actual)
		}
	}
}
//解法二: 位运算求解.
//复习一下与,或运算
// 与运算: 0 & 1 = 0, 1 & 1 = 0, 0 & 0 = 0 (双1则1,否则0)
// 或运算: 0 | 1 = 0, 1 | 1 = 1, 0 | 0 = 0 (有1则1, 否则为0)
//先计算
func IsUniqueV2(astr string) bool {
	if len(astr) <= 1 { //边界条件
		return true
	}
	mask := 0
	for i := 0; i < len(astr); i ++ {
		gap := astr[i] - 'a' //计算字符与'a'之间的差.
		bit := 1 << gap //1向左位移gap位.
		if (mask & bit) == 0 { //与运算结果还是0的位,则说明没有出现重复的. 如00 & 10 = 00, 010 & 010 = 010
			mask |= bit //或运算,则原是自身bit. 如00 | 10 = 10, 010 | 001 = 011
		} else {
			return false
		}
	}
	return true
}
func TestIsUniqueV2(t *testing.T) {
	//查看golang的A~Z之间的ASCII码值
	fmt.Println('a', 'z')//96, 122
	tests := []struct{
		data string //数据
		expect bool //预期值
	}{
		{"leetcode", false},
		{"abc", true},
		{"abcdc", false},
		{"abcdefghigka", false},
	}
	for index, item := range tests {
		if actual := IsUniqueV2(item.data); actual != item.expect {
			index ++
			t.Errorf("index:%d, expect:%v, actual:%v\n", index, item.expect, actual)
		}
	}
}