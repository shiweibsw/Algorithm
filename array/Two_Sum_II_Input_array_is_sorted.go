package main

import "fmt"

/**
找出两个数之和等于 target 的两个数字，要求输出它们的下标。注意一个数字不能使用 2 次。
下标从小到大输出。假定题目一定有一个解。
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]

Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2

思路：
指针对撞或者map集合判断
时间复杂度：O(n) 空间复杂度O(1)
*/

func main() {
	arr := []int{2, 7, 11, 15}
	fmt.Println(twoSum167(arr, 9))
}

func twoSum167(num []int, target int) []int {
	i, j := 0, len(num)-1
	for i < j {
		if num[i]+num[j] == target {
			return []int{i + 1, j + 1}
		} else if num[i]+num[j] < target {
			i++
		} else {
			j--
		}
	}
	return []int{-1, -1}
}

func twoSum167_1(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		another := target - num
		if _, ok := m[another]; ok {
			return []int{m[another] + 1, i + 1}
		} else {
			m[num] = i
		}
	}
	return nil
}
