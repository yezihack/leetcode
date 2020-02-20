package array

import (
	"fmt"
	"testing"
)

//leetcode:27 simple
/*
给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

示例 1:

给定 nums = [3,2,2,3], val = 3,

函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。[2, 2]

你不需要考虑数组中超出新长度后面的元素。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//题目的意思是: 返回重复元素的个数, 并将除非目标元素的元素向前移动.
func RemoveElement(nums []int, val int) int {
	ans := 0 //计算长度.
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ { //遍历整个数组
		if nums[i] != val { //数组元素与目标值比较, 如果不相等,则进行操作.
			nums[ans] = nums[i] //将数组元素填充到数组第一个位置
			ans++               //然后进行下一次
		}
	}
	fmt.Print(nums)
	return ans
}
func TestRemoveElement(t *testing.T) {
	tests := []struct {
		data   []int //数据
		val    int
		expect int //预期值
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{1, 2, 3, 1, 2, 3}, 2, 4},
	}
	for index, item := range tests {
		if actual := RemoveElementV3(item.data, item.val); actual != item.expect {
			index++
			t.Errorf("index:%d, expect:%d, actual:%d\n", index, item.expect, actual)
		}
	}
}

func RemoveElementV2(nums []int, val int) int {
	i := 0
	length := len(nums)
	fmt.Println(val, nums)
	for i < length {
		if nums[i] == val {
			nums[i] = nums[length-1]
			length--
		} else {
			i++
		}
	}
	fmt.Println(nums)
	return length
}
func RemoveElementV3(nums []int, val int) int {
	fmt.Println("after", nums)
	count := 0 //重复的元素个数.
	for i := 0; i < len(nums); i++ {
		if nums[i] != val { //与目标值不相等则向前移动.
			nums[count] = nums[i]
			count++
		}
	}
	fmt.Println("after", nums)
	return count
}
