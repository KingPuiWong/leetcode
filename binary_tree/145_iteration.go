package binary_tree

import (
	"container/list"
)

/*
后序遍历
*/

func postorderTraversal2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		res = append(res, node.Val)
		if node.Left != nil {
			st.PushBack(node.Left)
		}
		if node.Right != nil {
			st.PushBack(node.Right)
		}
	}

	reverse(res)
	return res
}

func reverse(res []int) {
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l += 1
		r -= 1
	}
}
