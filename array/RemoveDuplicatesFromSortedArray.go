package main

import "fmt"

/**
给定一个有序数组 nums，对数组中的元素进行去重，使得原数组中的每个元素只有一个。最后返回去重以后数组的长度值,和重新排序后的数组
Given nums = [0,0,1,1,1,2,2,3,3,4],
Your function should return length = 5，with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.
It doesn't matter what values are set beyond the returned length.
思路：两个指针从头开始遍历，相等的话finder指针+1，不相等的话交换last+1和finder指针元素的位置，最后返回last+1的长度即可
时间复杂度：O(n) 空间复杂度O(1)
*/
func main() {
	var arr = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(arr))
	fmt.Println(arr)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}
