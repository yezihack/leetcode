package array

import (
	"testing"
)

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:
输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/

//暴力求解
//Time:O(n^2), space:O(1)
func MaxSubArray(nums []int) int {
	count := len(nums)
	if count == 0 {
		return 0
	}
	max := -1 << 32 //默认最小值.
	for i := 0; i < count; i++ {
		sum := 0
		for j := i; j < count; j++ {
			sum += nums[j] //累加操作
			if sum > max { //如果大于max则替换掉.
				max = sum
			}
		}
	}
	return max
}
func TestMaxSubArray(t *testing.T) {
	if MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) != 6 {
		t.Error("err")
	}
	if MaxSubArray([]int{-1}) != -1 {
		t.Error("err")
	}
}

//贪心求解
//Time:O(n), space:O(1)
func MaxSubArrayV2(nums []int) int {
	if len(nums) == 0 { //边界条件
		return 0
	}
	var (
		max = nums[0] //默认第一个元素最大.
		sum = nums[0] //默认第一个元素最大.
	)
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] > sum { //如果当前数组的值大于sum,则替换
			sum = nums[i]
		}
		if sum > max { //如果大于max,则替换
			max = sum
		}
	}
	return max
}
func TestMaxSubArrayV2(t *testing.T) {
	if MaxSubArrayV2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) != 6 {
		t.Error("err")
	}
	if MaxSubArrayV2([]int{-1}) != -1 {
		t.Error("err")
	}
}

//动态求解
//Time:O(n), space:O(n)
//递推公式: F(n) = Max(f(n-1), nums[n]),
func MaxSubArrayV3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//申请一个dp数组.
	dp := make([]int, len(nums))
	//默认第一个元素最大.
	dp[0] = nums[0]
	//max function
	Max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	result := 0
	for i := 1; i < len(nums); i++ {
		dp[i] = Max(dp[i-1]+nums[i], nums[i]) //递推公式
		result = Max(dp[i], result)           //求最大值.
	}
	return result
}
func TestMaxSubArrayV3(t *testing.T) {
	if MaxSubArrayV3([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) != 6 {
		t.Error("err")
	}
	if MaxSubArrayV3([]int{-1}) != -1 {
		t.Error("err2")
	}
}
