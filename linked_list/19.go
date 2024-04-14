package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 使用快慢指针的方法
// 要求第 len - n 个节点,快指针先走 n,则快 = 慢 + n , 当快 = len 时,慢 = len - n
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	soft, fast := dummy, dummy
	for n > 0 && fast != nil {
		fast = fast.Next
		n--
	}

	fast = fast.Next

	for fast != nil {
		fast = fast.Next
		soft = soft.Next
	}

	soft.Next = soft.Next.Next

	return dummy.Next
}
