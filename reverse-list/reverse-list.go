package reverse_list

//定义list节点
type ListNode struct {
	Val int
	Next *ListNode
}

//https://algorithm.yuanbin.me/zh-hans/linked_list/reverse_linked_list.html

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}

	//return prev
	return prev
}


