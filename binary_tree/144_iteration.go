package binary_tree

import (
	"container/list"
)

/*
前序遍历的迭代写法
*/
func preorderTraversal2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		res = append(res, node.Val)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return res
}
