package main

import (
	"fmt"
	"sort"
)

/**
给一个已经有序的数组，返回的数组也必须是有序的，且数组中的每个元素是由原数组中每个数字的平方得到的
Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]
*/
func main() {
	arr := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(arr))
	fmt.Println(sortedSquares1(arr))

}

/**
解法一
*/
func sortedSquares(A []int) []int {
	ans := make([]int, len(A))
	for i, k, j := 0, len(A)-1, len(ans)-1; i <= j; k-- {
		if A[i]*A[i] > A[j]*A[j] {
			ans[k] = A[i] * A[i]
			i++
		} else {
			ans[k] = A[j] * A[j]
			j--
		}
	}
	return ans
}

/**
解法二
*/
func sortedSquares1(A []int) []int {
	for i, value := range A {
		A[i] = value * value
	}
	sort.Ints(A)
	return A
}
