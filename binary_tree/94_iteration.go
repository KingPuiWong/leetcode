package binary_tree

import (
	"container/list"
)

/*
中序遍历 迭代写法
*/

func inorderTraversal2(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	st := list.New()

	for root != nil || st.Len() > 0 {
		for root != nil {
			st.PushBack(root)
			root = root.Left
		}
		root = st.Remove(st.Back()).(*TreeNode)
		ans = append(ans, root.Val)
		root = root.Right
	}
	return ans
}
