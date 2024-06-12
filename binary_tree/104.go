package binary_tree

import (
	"container/list"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 递归写法
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 迭代写法
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var level int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
		}
		level++
	}

	return level
}
