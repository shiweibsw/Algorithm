package main

import "fmt"

/**
合并两个已经有序的数组，结果放在第一个数组中，第一个数组假设空间足够大。要求算法时间复杂度足够低。
Example:

Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]

思路：
为了不大量移动元素，就要从2个数组长度之和的最后一个位置开始，依次选取两个数组中大的数，
从第一个数组的尾巴开始往头放，只要循环一次以后，就生成了合并以后的数组了。
时间复杂度：O(n) 空间复杂度O(1)
*/
func main() {
	arr1 := []int{1, 2, 3, 0, 0, 0}
	arr2 := []int{2, 5, 6}
	merge(arr1, 3, arr2, 3)
	fmt.Println(arr1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	i := m - 1
	j := n - 1
	k := m + n - 1
	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
	//因为nums2 可能比nums1 要长，所有需要把nums2 剩下的数据也放在nums1中
	for ; j >= 0; k-- {
		nums1[k] = nums2[j]
		j--
	}
}
