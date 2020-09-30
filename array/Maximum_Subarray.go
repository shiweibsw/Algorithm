package main

import "fmt"

/**
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
思路：题目要求输出数组中某个区间内数字之和最大的那个值。dp[i] 表示 [0,i] 区间内各个子区间和的最大值，状态转移方程是 dp[i] = nums[i] + dp[i-1] (dp[i-1] > 0)，dp[i] = nums[i] (dp[i-1] ≤ 0)。
时间复杂度：O(n) 空间复杂度：O(n)
*/
func main() {
	aar := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(aar))
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		res = max1(res, dp[i])
	}
	return res
}

func max1(x1 int, x2 int) int {
	if x1 >= x2 {
		return x1
	} else {
		return x2
	}
}
