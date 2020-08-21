package main

import "fmt"

/**
给定一个数组 nums 和一个数值 val，将数组中所有等于 val 的元素删除，并返回剩余的元素个数。
Given nums = [0,1,2,2,3,0,4,2], val = 2,

Your function should return length = 5, with the first five elements of nums containing 0, 1, 3, 0, and 4.

Note that the order of those five elements can be arbitrary.

It doesn't matter what values are set beyond the returned length.
思路：
这里数组的删除并不是真的删除，只是将删除的元素移动到数组后面的空间内
时间复杂度：O(n) 空间复杂度O(1)
*/
func main() {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(arr, 2))
	fmt.Println(arr)
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	return j
}
