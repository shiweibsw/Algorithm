package main

import "fmt"

/**
2020-08-25

如果数组里面有重复数字就输出 true，否则输出 flase
Example 1:
Input: [1,2,3,1]
Output: true
Example 2:
Input: [1,2,3,4]
Output: false
思路:
用 map 判断即可
时间复杂度：O(n) 空间复杂度O(n)
*/

func main() {
	arr := []int{1, 2, 3, 1}
	fmt.Println(ContainsDuplicate(arr))
}

func ContainsDuplicate(nums []int) bool {
	record := make(map[int]bool, len(nums))
	for _, i := range nums {
		if _, found := record[i]; found {
			return true
		}
		record[i] = true
	}
	return false
}
