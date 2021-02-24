package linkedlist

type Node struct {
     Val int
     Next *Node
     Random *Node
 }

func CopyRandomList(head *Node) *Node {
	//头指针为空直接返回
	if head==nil{
		return head
	}
	//将复制的新节点放在复制节点的后面
	p:=head
	for p!=nil {
		node:=&Node{p.Val,nil,nil}
		node.Next=p.Next
		p.Next=node
		//向后移动
		p=p.Next.Next
	}
	//将random的指针也复制过来
	p=head
	for p!=nil {
		if p.Random!=nil{
			p.Next.Random=p.Random.Next
		}
		p=p.Next.Next
	}
	//将链表分为两个链表，第二个链表就是复制完后的链表
	p=head
	newList:=p.Next
	q:=p.Next
	for p!=nil {
		p.Next=p.Next.Next
		p=p.Next
		if q.Next!=nil {
			q.Next=q.Next.Next
			q=q.Next
		}
	}
	return newList
}