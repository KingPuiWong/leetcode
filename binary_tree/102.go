package binary_tree

import (
	"container/list"
)

func levelOrder(root *TreeNode) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		var arr []int
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}

			arr = append(arr, node.Val)
		}

		res = append(res, arr)
	}
	return res
}
