package main

func main() {

}

type TreeNode struct {
	val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode, output []int) []int {
	inorder(root, &output)
	return output
}

func inorder(root *TreeNode, output *[]int) {
	if root != nil {
		inorder(root.Left, output)
		*output = append(*output, root.val)
		inorder(root.Right, output)
	}
}
