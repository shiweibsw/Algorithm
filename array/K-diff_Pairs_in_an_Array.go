package main

import "fmt"

/**
2020-08-26
给定一个数组，在数组里面找到几组不同的 pair 对，每个 pair 对相差 K 。问能找出多少组这样的 pair 对
Example 1:
Input: [3, 1, 4, 1, 5], k = 2
Output: 2
Explanation: There are two 2-diff pairs in the array, (1, 3) and (3, 5).
Although we have two 1s in the input, we should only return the number of unique pairs.

Example 2:

Input:[1, 2, 3, 4, 5], k = 1
Output: 4
Explanation: There are four 1-diff pairs in the array, (1, 2), (2, 3), (3, 4) and (4, 5).

Example 3:

Input: [1, 3, 1, 5, 4], k = 0
Output: 1
Explanation: There is one 0-diff pair in the array, (1, 1).

思路：
这一题可以用 map 记录每个数字出现的次数。
重复的数字也会因为唯一的 key，不用担心某个数字会判断多次。
遍历一次 map，每个数字都加上 K 以后，判断字典里面是否存在，
如果存在， count ++，如果 K = 0 的情况需要单独判断，
如果字典中这个元素频次大于 1，count 也需要 ++

时间复杂度：O(n) 空间复杂度O(n)
*/
func main() {
	arr := []int{3, 1, 4, 1, 5}
	fmt.Println(findPairs(arr, 3))
}
func findPairs(nums []int, k int) int {
	if k < 0 || len(nums) == 0 {
		return 0
	}
	var count int
	m := make(map[int]int, len(nums))
	for _, value := range nums {
		m[value]++
	}
	for key := range m {
		if k == 0 && m[key] > 1 {
			count++
			continue
		}
		if k > 0 && m[key+k] > 0 {
			count++
		}
	}
	return count
}
