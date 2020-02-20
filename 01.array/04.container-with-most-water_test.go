package array

import (
	"testing"
)

//leetcode:11 medium

/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
Find two lines, which together with x-axis forms a container,
such that the container contains the most water.
*/

//暴力求解
//Time:O(n^2), Space:O(1)
func ContainerWithMostWater(height []int) int {
	if len(height) == 0 {
		return 0
	}
	//暴力求解, 任何可能都不放过.
	maxArea := 0 //存放最大面积的变量.
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			//获取最短板的那个数字,也就是最小值的数字
			minHeight := height[i]
			if height[j] < minHeight {
				minHeight = height[j]
			}
			//获取j与i之间的差距离.
			distance := j - i
			//求面积.
			area := minHeight * distance
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
func TestContainerWithMostWater(t *testing.T) {
	tests := []struct {
		data   []int //数据
		expect int   //预期值
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	}
	for index, item := range tests {
		if actual := ContainerWithMostWater(item.data); actual != item.expect {
			index++
			t.Errorf("index:%d, expect:%d, actual:%d\n", index, item.expect, actual)
		}
	}
}

//解法二: 双指针解法
//Time: O(n), Space:O(1)
func ContainerWithMostWaterV2(height []int) int {
	if len(height) == 0 { //边界条件
		return 0
	}
	left, right := 0, len(height)-1 //声明左右指针.
	maxArea := 0                    //声明存放最大面积的变量.
	for left < right {              //左边指针相遇则结束程序
		distance := right - left          //求解左右之间的差距离.可能想象成一个矩形,求最大面积
		minHeight := height[left]         //左右两边的数字求最小, 可以想象成一个桶, 决定桶能盛最多水,取决于最短的那根木板.
		if height[right] < height[left] { //如果右边小于左边,则取右边的木板高度.
			minHeight = height[right]
			right-- //移动右指针继续下一次计算.这是是减减,因为右指针要向左边移动
		} else {
			left++ //反之.左边的指针要向右移动.
		}
		area := distance * minHeight //计算当前的面积
		if area > maxArea {          //如果大于之前计算过的面积则替换之.
			maxArea = area
		}
	}
	//返回最大的面积值.
	return maxArea
}
func TestContainerWithMostWaterV2(t *testing.T) {
	tests := []struct {
		data   []int //数据
		expect int   //预期值
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	}
	for index, item := range tests {
		if actual := ContainerWithMostWaterV4(item.data); actual != item.expect {
			index++
			t.Errorf("index:%d, expect:%d, actual:%d\n", index, item.expect, actual)
		}
	}
}

func ContainerWithMostWaterV3(height []int) int {
	i, j := 0, len(height)-1
	maxArea := 0
	for i < j {
		min := height[i]
		dis := j - i
		if min > height[j] {
			min = height[j]
			j--
		} else {
			i++
		}
		if area := dis * min; area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
func ContainerWithMostWaterV4(height []int) int {
	var (
		maxArea = 0
		head    = 0
		tail    = len(height) - 1
	)
	Min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	Max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for head < tail {
		area := (tail - head) * Min(height[head], height[tail])
		maxArea = Max(maxArea, area)
		if height[head] > height[tail] {
			tail--
		} else {
			head++
		}
	}
	return maxArea
}
