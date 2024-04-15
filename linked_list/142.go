package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// fast 和 slow 同时走,fast 走两步， slow 走一步
// 因为可能环内相遇,所以 slow 继续走,head这时候也继续走,如果相遇则是环的入口
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			for head != slow {
				head = head.Next
				slow = slow.Next
			}

			return head
		}
	}
	return nil
}
