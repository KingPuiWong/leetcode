package binary_tree

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	}

	if root.Right == nil && root.Left != nil {
		return 1 + minDepth(root.Left)
	}

	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
