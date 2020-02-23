package array

import (
	"testing"
)

//来自<<剑指offer>> 面试题03. 数组中重复的数字
/*
找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
数组中某些数字是重复的，但不知道有几个数字重复了，
也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3


限制：

2 <= n <= 100000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//解法一: 利用hash
func FindRepeatNumber(nums []int) int {
	if len(nums) == 0 { //边界条件
		return 0
	}
	//跟据题目要求: 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内
	//直接检查是否超出范围.
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 || nums[i] > len(nums)-1 {
			return 0
		}
	}
	//申请一个hash辅助
	hash := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok { //如果hash存在这个值,说明重复.
			return nums[i]
		}
		hash[nums[i]] = struct{}{}
	}
	return 0
}
func TestFindRepeatNumber(t *testing.T) {
	tests := []struct {
		data   []int //数据
		expect []int //预期值
	}{
		{[]int{2, 3, 1, 0, 2, 5, 3}, []int{2, 3}},
	}
	for index, item := range tests {
		if actual := FindRepeatNumber(item.data); !InArray(actual, item.expect) {
			index++
			t.Errorf("index:%d, expect:%v, actual:%d\n", index, item.expect, actual)
		}
	}
}

//解法二: 题目数字只在0~n-1之间的范围, 说明, 数字与下标是一样的.如果出现不一致则说明有重复的.

func FindRepeatNumberV2(nums []int) int {
	if len(nums) == 0 { //boundary condition
		return 0
	}
	//See if the elements is out of bounds
	for i := 0; i < len(nums);i ++ {
		if nums[i] < 0 || nums[i] > len(nums) -1 {
			return 0
		}
	}
	//遍历一遍.
	for i := 0; i < len(nums); i ++ {
		for nums[i] != i { //如果下标与数字不相等,则交换处理
			if nums[i] == nums[nums[i]] { //遇到数字与下标对应的数字相等,则说明有重复.
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return 0
}

func TestFindRepeatNumberV2(t *testing.T) {
	tests := []struct {
		data   []int //数据
		expect []int //预期值
	}{
		{[]int{2, 3, 1, 0, 2, 5, 3}, []int{2, 3}},
	}
	for index, item := range tests {
		if actual := FindRepeatNumberV2(item.data); !InArray(actual, item.expect) {
			index++
			t.Errorf("index:%d, expect:%v, actual:%d\n", index, item.expect, actual)
		}
	}
}

func InArray(a int, arr []int) bool {
	if len(arr) == 0 {
		return false
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == a {
			return true
		}
	}

	return false
}
