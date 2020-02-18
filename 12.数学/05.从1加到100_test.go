/*
 * @Author: your name
 * @Date: 2020-02-18 10:11:01
 * @LastEditTime: 2020-02-18 11:12:01
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \leetcode\12.数学\05.从1加到100_test.go
 */
package mathematics

import "testing"

//time: time:O(n), space:O(1)
func Sum100(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func Sum100V2(n int) int {
	return n * (n + 1) / 2
}

func TestSum100(t *testing.T) {
	if Sum100(100) != 5050 {
		t.Error("sum is err")
	}
}
func TestSum100V1(t *testing.T) {

	if Sum100V2(100) != 5050 {
		t.Error("sum is err")
	}
}
