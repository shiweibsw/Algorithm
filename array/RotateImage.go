package main

import "fmt"

/**
给定一个 n × n 的二维矩阵表示一个图像。将图像顺时针旋转 90 度。说明：你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
Given input matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

rotate the input matrix in-place such that it becomes:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
时间复杂度：O(n) 空间复杂度：O(1)


*/
func main() {
	aar := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(aar)
	fmt.Println(aar)
}
func rotate(matrix [][]int) {
	row := len(matrix)
	if row <= 0 {
		return
	}
	fmt.Println(matrix)
	column := len(matrix[0])
	for i := 0; i < row; i++ { // 对角线变换
		for j := i + 1; j < column; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
	fmt.Println(matrix)
	halfColumn := column / 2
	for i := 0; i < row; i++ {
		for j := 0; j < halfColumn; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[i][column-j-1]
			matrix[i][column-j-1] = tmp
		}
	}
}
