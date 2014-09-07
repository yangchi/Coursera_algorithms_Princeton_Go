package stack

/*
A linked-list based stack. This one uses the built-in list container from golang.
*/

import ("container/list")

type StackL struct {
	values *list.List
}

func (stack *StackL) init() {
	stack.values = list.New()
}

func (stack *StackL) size () (uint) {
	return uint(stack.values.Len())
}

func (stack *StackL) isEmpty() (bool) {
	return stack.size() == 0
}

func (stack *StackL) push(str string) {
	stack.values.PushBack(str)
}

func (stack *StackL) pop () (string) {
	elem := stack.values.Back()
	str := elem.Value.(string)
	stack.values.Remove(elem)
	return str

}
