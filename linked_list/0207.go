package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 我们求出两个链表的长度，并求出两个链表长度的差值，然后让curA移动到，和curB 末尾对齐的位置
// 此时我们就可以比较curA和curB是否相同，如果不相同，同时向后移动curA和curB，如果遇到curA == curB，则找到交点。
// 否则循环退出返回空指针。
// !!! 指针相同,而不是数值相同
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	lenA, lenB := 0, 0
	for curA != nil {
		lenA++
		curA = curA.Next
	}

	for curB != nil {
		lenB++
		curB = curB.Next
	}

	newCurA, newCurB := headA, headB
	if lenA > lenB {
		subLen := lenA - lenB
		for subLen > 0 {
			newCurA = newCurA.Next
			subLen--
		}
	} else {
		subLen := lenB - lenA
		for subLen > 0 {
			newCurB = newCurB.Next
			subLen--
		}
	}

	for newCurA != newCurB {
		newCurA = newCurA.Next
		newCurB = newCurB.Next
	}
	return newCurA
}
