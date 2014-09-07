package stack

/*
A linked-list based stack. This one doesn't use the built-in list container from golang.
*/

type node struct {
	value string
	next *node
}

type StackLL struct {
	head *node
}

func (stack *StackLL) push(str string) {
	newnode := new(node)
	newnode.value = str
	newnode.next = stack.head
	stack.head = newnode
}

func (stack *StackLL) pop() (string) {
	str := stack.head.value
	stack.head = stack.head.next
	return str
}

func (stack *StackLL) isEmpty() (bool) {
	return stack.head == nil
}
