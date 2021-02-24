package linkedlist

func reverseList(head *ListNode) *ListNode {
	//声明三个变量，将this指向当前指针
	var pre, this, next *ListNode
	this = head
	for this != nil {
		//当前指针的next
		next = this.Next
		//上一个指针变为当前指针的next
		this.Next = pre
		//当前指针变为下一个指针的上一指针
		pre = this
		//当前指针向后移
		this = next
	}
	//返回最后一个下一指针的上一指针就为新链表的头指针
	return pre
}
