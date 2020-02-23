package array

import (
	"fmt"
	"testing"
)

//26. 删除排序数组中的重复项
/*
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

示例 1:

给定数组 nums = [1,1,2],

函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

你不需要考虑数组中超出新长度后面的元素。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//题目意思是: 移除相同的元素.并将唯一的元素向前放的一个数组.
func RemoveDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	count := 0 //统计重复出现的次数
	fmt.Println("before", nums)
	for i := 1; i < len(nums); i++ { //循环整个数组
		if nums[i] == nums[i-1] { //数组的前与后的元素进行对比.如果相等则,记录count
			count++
		} else { //如果不相等则交换位置,注意这里是使用i-count,count是记录出现重复的次数.
			nums[i-count] = nums[i]
		}
		fmt.Printf("i:%d, count:%d, %v\n", i, count, nums)
	}
	fmt.Println("after", nums)
	return len(nums) - count //返回整个数组长度减去重复的count次数,则等于不相等的数组长度.
}
func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		data   []int //数据
		expect int   //预期值
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}
	for index, item := range tests {
		if actual := RemoveDuplicates(item.data); actual != item.expect {
			index++
			t.Errorf("index:%d, expect:%d, actual:%d\n", index, item.expect, actual)
		}
	}
}

func RemoveDuplicatesV2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	fmt.Println("before", nums)
	count := 0                       //记录出现重复的元素的个数.
	for i := 1; i < len(nums); i++ { //遍历整个数组, 注意len(nums)-1.因为下面使用了i+1操作.
		if nums[i] != nums[count] { //当前元素与count位置的元素不相等则交换.
			count++
			nums[count] = nums[i]
		}
	}
	fmt.Println("after", nums)
	return count + 1
}
func RemoveDuplicatesFunc(nums []int) int {
	size := len(nums)
	if size <= 1 {
		return size
	}
	count := 0
	for i := 1; i < size; i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			nums[i-count] = nums[i]
		}
	}
	fmt.Println(nums[0 : size-count])
	//nums = nums[0 : size-count]
	newArr := make([]int, size-count)
	copy(newArr, nums)
	nums = nil
	return size - count
}
func TestRemoveDuplicatesFunc(t *testing.T) {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	count := RemoveDuplicatesFunc(arr)
	fmt.Println(arr, count)
}

func TestSlice(t *testing.T) {
	arr := []int{1, 2, 3}
	fmt.Printf("before:%p\n", arr)
	del := func(a []int) {
		a[0] = 100
		a = append(a[:0], a[0:1]...)
		fmt.Printf("a:%p, %v\n", a, a)
	}
	del(arr)
	fmt.Printf("before:%p\n", arr)
	fmt.Println(arr)
}
