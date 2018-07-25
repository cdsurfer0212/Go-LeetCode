package mergetwosortedlists

import "../utils"

func mergeTwoLists(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var result *utils.ListNode
	if l1.Val < l2.Val {
		result = l1
		result.Next = mergeTwoLists(l1.Next, l2)
	} else {
		result = l2
		result.Next = mergeTwoLists(l1, l2.Next)
	}
	return result
}
