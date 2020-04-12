package _4_递归

import (
	"fmt"
	"sort"
	"testing"
)

/*
你正在使用一堆木板建造跳水板。有两种类型的木板，其中长度较短的木板长度为shorter，长度较长的木板长度为longer。你必须正好使用k块木板。编写一个方法，生成跳水板所有可能的长度。

返回的长度需要从小到大排列。

示例：

输入：
shorter = 1
longer = 2
k = 3
输出： {3,4,5,6}
提示：

0 < shorter <= longer
0 <= k <= 100000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diving-board-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 思路:
// 题目的意思就是求一个k的组合, 利用给出的长短木板进行组合.
// 边界条件考虑:
// 如果k = 0, 也就是不需要组合.返回空数组.
// 如果长短目录都相等, 最长的目录组就是 k * 长度(shorter, longer)

func LivingBoard(shorter, longer, k int) []int {
	if k == 0 { // 也就是不需要组合.返回空数组.
		return nil
	}
	if shorter == longer { // 如果长短目录都相等, 最长的目录组就是 k * 长度(shorter, longer)
		return []int{shorter * k}
	}
	group := make([]int, k+1) // 为什么k+1, 因为在组合中我可以使用全是长木板也可以使用全是短木板.
	for i := 0; i < k+1; i++ {
		fmt.Printf("i:%d, k-i:%d\n", i, k-i)
		// 当i=0时,不采用短木板,也就是k-i=k, 表示全使用长木板.
		// 当i=k时,采用全是短木板,而k-k=0,表示不采用长木板
		group[i] = i*shorter + (k-i)*longer
	}
	// 进行排序一下.从低到高的顺序.
	sort.Ints(group)
	return group
}
func TestLivingBoard(t *testing.T) {
	result := LivingBoard(1, 2, 3)
	fmt.Println(result)
}
