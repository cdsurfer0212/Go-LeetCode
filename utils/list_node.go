package utils

// ListNode ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func ConvertIntArrayToListNode(intArray []int) *ListNode {
	var result *ListNode
	var temp *ListNode
	for _, element := range intArray {
		if result == nil {
			result = &ListNode{
				Val: element,
			}
			temp = result
		} else {
			temp.Next = &ListNode{
				Val: element,
			}
			temp = temp.Next
		}
	}
	return result
}

func ConvertListNodeToIntArray(listNode *ListNode) []int {
	var result []int
	for listNode != nil {
		result = append(result, listNode.Val)
		listNode = listNode.Next
	}
	return result
}
