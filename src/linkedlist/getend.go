package linkedlist

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	p, _ := getend(head, k)
	return p
}
func getend(head *ListNode, k int) (*ListNode, int) {
	//如果节点的下一个节点为空，说明他是最后一个节点
	if head.Next == nil {
		return head, 1
	}
	//递归直到最后一个节点
	p, count := getend(head.Next, k)
	//判断为倒数第几个节点，若是要找的则将指针指向并返回，若找到之后递归回去传的一直为找到的节点，count会发生变化
	count++
	if count == k {
		p = head
	}
	return p, count
}
