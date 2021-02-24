package linkedlist

func reversePrint(head *ListNode) []int {
	//头指针为空直接返回Nil
	if head == nil {
		return nil
	}
	//计算所需slice大小
	p := head
	count := 0
	for p != nil {
		count++
	}
	//再次遍历链表，将值倒叙赋给数组
	sli := make([]int, count)
	p = head
	for p != nil {
		count--
		sli[count] = p.Val
		p = p.Next
	}
	return sli
}
