package main

import "fmt"

/**
在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1]

思路：

这道题最优的做法时间复杂度是 O(n)。
顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字，
如果找到了，直接返回 2 个数字的下标即可。如果找不到，就把这个数字存入 map 中，
等待扫到“另一半”数字的时候，再取出来返回结果。
map 中key为数组元素，value为下标。

时间复杂度：O(n)，空间复杂度：O(n)
*/

func main() {
	arr := []int{2, 7, 11, 15}
	fmt.Println(twoSum(arr, 9))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}
