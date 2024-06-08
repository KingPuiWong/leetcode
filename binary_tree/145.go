package binary_tree

/*
后序遍历
*/

func postorderTraversal(root *TreeNode) []int {
	var res []int
	reversal := func(node *TreeNode) {}
	reversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		reversal(node.Left)
		reversal(node.Right)
		res = append(res, node.Val)
	}

	reversal(root)
	return res
}
