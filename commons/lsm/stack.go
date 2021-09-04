/*
 * @Author: your name
 * @Date: 2021-09-05 01:09:47
 * @LastEditTime: 2021-09-05 01:10:47
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /lsm_tree_pro/commons/lsm/stack.go
 */
package lsm

type Stack struct {
	Vals []int
}

/** initialize your data structure here. */
func InitStack() *Stack {
	stack := &Stack{}
	stack.Vals = make([]int, 0)
	return stack
}

func (s *Stack) Push(x int) {
	s.Vals = append(s.Vals, x)
}

func (s *Stack) Pop() {
	if s.Size() != 0 {
		s.Vals = s.Vals[:s.Size()-1]
	}
}

func (s *Stack) Top() int {
	if s.Size() == 0 {
		return 2147483647
	}
	return s.Vals[s.Size()-1]
}

func (s *Stack) Size() int {
	return len(s.Vals)
}

type MinStack struct {
	MStack *Stack
	SStack *Stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	minStack := MinStack{}
	minStack.MStack = InitStack()
	minStack.SStack = InitStack()
	return minStack
}

func (s *MinStack) Push(x int) {
	s.MStack.Push(x)
	if s.SStack.Top() >= x {
		s.SStack.Push(x)
	}
}

func (s *MinStack) Pop() {
	if s.MStack.Size() > 0 {
		val := s.MStack.Top()
		if s.SStack.Top() == val {
			s.SStack.Pop()
		}
		s.MStack.Pop()
	}
}

func (s *MinStack) Top() int {
	return s.MStack.Top()
}

func (s *MinStack) Min() int {
	return s.SStack.Top()
}
