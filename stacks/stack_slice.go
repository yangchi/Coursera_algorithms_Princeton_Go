package stack

/*
A slice based stack.
*/

type StackS struct {
	values []string
	current uint
}

func (stack *StackS) init() {
	stack.values = make([]string, 10)
	stack.current = 0
}

func (stack *StackS) push(str string) {
}
