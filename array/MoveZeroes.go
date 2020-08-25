package main

import "fmt"

/**
2020-08-25
要求不能采用额外的辅助空间，将数组中 0 元素都移动到数组的末尾，并且维持所有非 0 元素的相对位置
Example 1:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]

思路:
可以只扫描数组一遍，不断的用 i，j 标记 0 和非 0 的元素，然后相互交换，最终到达题目的目的
时间复杂度：O(n) 空间复杂度O(1)
*/
func main() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	fmt.Println(arr)
}

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	j := 0
	for i, n := range nums {
		if n != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
}
