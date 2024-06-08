package binary_tree

func inorderTraversal(root *TreeNode) []int {
	var res []int
	traversal := func(node *TreeNode) {}
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return res
}
