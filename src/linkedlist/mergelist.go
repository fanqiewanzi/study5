package linkedlist


type ListNode struct {
     Val int
     Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists==nil{
		return nil
	}
	//两两结合，结合到最后一个链表为止
	//这里设第一个链表为最终合并排序后的链表
	min:=lists[0]
	for i:=1;i< len(lists);i++ {
		min=merge(min,lists[i])
	}
	return min
}

func merge(list1 *ListNode,list2 *ListNode) *ListNode {
	//当其中任何一个链表走到头时返回另一个链表
	if list1==nil||list2==nil{
		if list1==nil{
			return list2
		}else{
			return list1
		}
	}
	//对链表的值进行比较，比另一个小则对他下一个指针进行递归比较，以此来达到重置将链表的指针排序的效果
	if list1.Val<list2.Val{
		list1.Next=merge(list1.Next,list2)
		return list1
	}else {
		list2.Next=merge(list1,list2.Next)
		return list2
	}

}