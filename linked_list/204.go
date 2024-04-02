package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	cur := head
	if cur.Next != nil {
		next := cur.Next
		nextNext := next.Next
		next.Next = cur
		cur.Next = nextNext
		cur = nextNext
	}

	return head
}
