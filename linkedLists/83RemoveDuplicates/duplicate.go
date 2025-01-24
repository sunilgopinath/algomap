package removeduplicates

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/

type ListNode struct {
	Val int
	Next *ListNode
}

// DeleteDuplicates removes duplicates from a sorted linked list
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current := head
	for current.Next != nil {
		if current.Val == current.Next.Val { // found duplicate as list is sorted
			current.Next = current.Next.Next
		} else {
			current = current.Next // increment "counter"
		}
	}
	return head
}