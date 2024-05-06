package main

type ListNode struct {
    Val int
    Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	head = nextMax(head)
	return head
}

func nextMax(node *ListNode) *ListNode {
	if node.Next == nil {
		return node
	}
	if node.Val < node.Next.Val {
		return nextMax(node.Next)
	}
	node.Next = nextMax(node.Next)
	if node.Val < node.Next.Val {
		return node.Next
	}
	return node
}

func main() {
	l := makeList([]int{5,2,13,3,8})
	printList(removeNodes(l))

	l = makeList([]int{1,1,1,1})
	printList(removeNodes(l))
}

func makeList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0], Next: nil}
	node := head
	for i := 1; i < len(vals); i++ {
		node.Next = &ListNode{Val: vals[i], Next: nil}
		node = node.Next
	}
	return head
}

func printList(head *ListNode) {
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}
