package main

import "fmt"

/**
2020-08-25
如果数组里面有重复数字，并且重复数字的下标差值小于等于 K 就输出 true，如果没有重复数字或者下标差值超过了 K ，则输出 flase。
Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true

Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true

Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false
思路：这道题可以维护一个只有 K 个元素的 map，每次只需要判断这个 map 里面是否存在这个元素即可。
如果存在就代表重复数字的下标差值在 K 以内。map 的长度如果超过了 K 以后就删除掉 i-k 的那个元素，这样一直维护 map 里面只有 K 个元素。

*/
func main() {
	arr := []int{1, 0, 1, 1}
	//arr := []int{1, 2, 3, 1, 2, 3}
	fmt.Println(containsNearbyDuplicate(arr, 1))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}
	if k <= 0 {
		return false
	}
	record := make(map[int]bool, len(nums))
	for i, n := range nums {
		if _, found := record[n]; found {
			return true
		}
		record[n] = true
		if len(record) == k+1 {
			delete(record, nums[i-k])
		}
	}
	return false
}
