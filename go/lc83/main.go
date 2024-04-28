package main

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	appeared := make(map[int]bool)
	cur := head.Next
	prev := head
	appeared[head.Val] = true
	for cur != nil {
		_, ok := appeared[cur.Val]
		if ok {
			cur = cur.Next
			prev.Next = cur
		} else {
			appeared[cur.Val] = true
			cur = cur.Next
			prev = prev.Next
		}
	}
	return head
}

func main() {

}
