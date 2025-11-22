package pkg

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// this pointer keeps changing to the next digit every iter
	var rolling_listNode *ListNode

	// pointer to starting ListNode
	var res *ListNode
	for list1 != nil || list2 != nil {
		// if list1 is null then we need to put all the items of list2 in the result
		if list1 == nil {
			for list2 != nil {
				set(&rolling_listNode, &res, list2.Val)
				list2 = list2.Next
			}
			break
		}

		// if list2 is null then we need to put all the items of list1 in the result
		if list2 == nil {
			for list1 != nil {
				set(&rolling_listNode, &res, list1.Val)
				list1 = list1.Next
			}
			break

		}
		// put list1 in result
		if list1.Val < list2.Val {
			set(&rolling_listNode, &res, list1.Val)
			list1 = list1.Next

		} else if list1.Val > list2.Val { // put list2 in result
			set(&rolling_listNode, &res, list2.Val)
			list2 = list2.Next
		} else {

			// they are the same , put both in result
			set(&rolling_listNode, &res, list1.Val)
			set(&rolling_listNode, &res, list1.Val)
			list1 = list1.Next
			list2 = list2.Next

		}

	}
	return res
}

func set(this **ListNode, final_rv **ListNode, val int) {
	if *this == nil {
		*this = &ListNode{Val: val}
		*final_rv = *this
	} else {
		(*this).Next = &ListNode{Val: val}
		*this = (*this).Next
	}
}
