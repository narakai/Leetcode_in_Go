package reverse_list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	one []int
}

type ans struct {
	one []int
}

//链表转数组
func l2s(head *ListNode) []int {
	res := []int{}

	for head != nil{
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

//数组转链表
func s2l(nums []int) *ListNode {
	if len(nums) == 0{
		return nil
	}
	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}
	return res
}

func Test(t *testing.T)  {
	ast := assert.New(t)

	qs := []question{
		{
			para{[]int{1, 2, 3, 4, 5}},
			ans{[]int{5, 4, 3, 2, 1}},
		},
		//test2
		{
			para{[]int{1, 2, 3, 4, 5, 6}},
			ans{[]int{6, 5, 4, 3, 2, 1}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("%v\n", p)
		ast.Equal(a.one, l2s(reverseList(s2l(p.one))), "输入:%v", p)
	}

}