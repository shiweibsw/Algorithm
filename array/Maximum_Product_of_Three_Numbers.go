package main

import (
	"fmt"
	"sort"
)

/**
2020-08-26
给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积
Example 1:

Input: [1,2,3]
Output: 6
Example 2:

Input: [1,2,3,4]
Output: 24

思路：给出一个数组，要求求出这个数组中任意挑 3 个数能组成的乘积最大的值。
题目的 test case 数据量比较大，如果用排序的话，时间复杂度高，可以直接考虑模拟，挑出 3 个数组成乘积最大值，
必然是一个正数和二个负数，或者三个正数。那么选出最大的三个数和最小的二个数，对比一下就可以求出最大值了。时间复杂度 O(n)
时间复杂度：O(n) 空间复杂度O(1)
*/
func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(maximumProduct(arr))
	fmt.Println(maximumProduct1(arr))
}

// O(nlogn)
func maximumProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 1
	if len(nums) <= 3 {
		for i := range nums {
			res = res * nums[i]
		}
		return res
	}

	sort.Ints(nums)
	if nums[len(nums)-1] <= 0 {
		return 0
	}
	return max(nums[0]*nums[1]*nums[len(nums)-1], nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3])
}

func max(x1 int, x2 int) int {
	if x1 >= x2 {
		return x1
	} else {
		return x2
	}
}

// O(n)
func maximumProduct1(nums []int) int {
	n1, n2, n3 := -1<<63, -1<<63, -1<<63
	n4, n5 := 0, 0
	for _, v := range nums {
		if v > n1 {
			n3 = n2
			n2 = n1
			n1 = v
		} else if v > n2 {
			n3 = n2
			n2 = v
		} else if v > n3 {
			n3 = v
		} else if v < n4 {
			n5 = n4
			n4 = v
		} else if v < n5 {
			n5 = v
		}
	}

	if n2*n3 > n4*n5 {
		return n1 * n2 * n3
	}
	return n1 * n4 * n5
}
