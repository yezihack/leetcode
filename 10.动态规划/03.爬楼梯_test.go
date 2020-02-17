/*
 * @Author: 百里
 * @Date: 2020-02-16 09:29:53
 * @LastEditTime : 2020-02-16 21:22:54
 * @LastEditors: Please set LastEditors
 * @Description:
 * @FilePath: \leetcode\10.动态规划\03.爬楼梯_test.go
 * @https://github.com/yezihack
 */
package dp

import (
	"fmt"
	"testing"
)

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//方法一: 动态规划
//爬1层,则是一种方法.爬2层,可以每次爬1层,也可以一次爬2层.
func ClimbStairs(n int) int {
	if n < 0 { //边界条件判断
		return 0
	}
	if n == 1 { //如果只有1层楼梯则只有1种方法.
		return n
	}
	//定义一个DP数组.存储每爬一层楼梯累枳的可能的方法
	//一种自下而上的思路
	dp := make([]int, n)
	dp[0] = 1                //如果楼梯只有1层,则只有1种爬法.
	dp[1] = 2                //如果楼梯有2层,则有2种
	for i := 2; i < n; i++ { //从第3层楼梯开始爬
		dp[i] = dp[i-1] + dp[i-2] //每爬一层,即是前面n-1层爬法和n-2层的爬法.
	}
	fmt.Println(dp)

	return dp[n-1] //取最后一楼,即求得n层楼梯的可能爬法.
}
func TestClimbStairs(t *testing.T) {
	tests := []struct {
		data   int //数据
		expect int //预期值
	}{
		{2, 2},
		{3, 3},
	}
	for index, item := range tests {
		if actual := ClimbStairs(item.data); actual != item.expect {
			index++
			t.Errorf("index:%d, expect:%d, actual:%d\n", index, item.expect, actual)
		}
	}
}
