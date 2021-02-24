package linkedlist

var count=0
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head!=nil{
		getKthFromEnd(head.Next,k)
	}
	count++
	if count==k{
		return head
	}
	return nil
}
