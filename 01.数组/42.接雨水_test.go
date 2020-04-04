package array

import (
	"fmt"
	"testing"
)

// 42. 接雨水
/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trapping-rain-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Trap(height []int) int {

	if len(height) < 1 {
		return 0
	}
	ans := 0
	for i := 1; i < len(height)-1; i++ {
		left, right := 0, 0 // 定义求左边最高值, 右边最高值
		for x := i; x >= 0; x-- {
			left = max(height[x], left)
		}
		for y := i; y < len(height)-1; y++ {
			right = max(height[y], right)
		}
		ans += min(left, right) - height[i]
	}
	return ans
}

func TestTrap(t *testing.T) {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(Trap(arr))
}

// -------------------------------------------------------------
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
