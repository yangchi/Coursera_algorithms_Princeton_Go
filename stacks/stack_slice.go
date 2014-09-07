package stack

/*
A slice based stack.
*/

type StackS struct {
	values []string
	current uint
}

func (stack *StackS) init() {
	stack.values = make([]string, 1) 
	stack.current = 0
}

func (stack *StackS) push(str string) {
	if stack.current == uint(len(stack.values) - 1) {
		newslice := make([]string, 2 * len(stack.values))
		copy(newslice, stack.values)
		stack.values = newslice
	}
	stack.values[stack.current] = str
	stack.current++
}

func (stack *StackS) pop() (string) {
	stack.current--
	str := stack.values[stack.current]
	stack.values[stack.current] = ""
	if stack.current < uint(len(stack.values) / 4) {
		newslice := make([]string, len(stack.values)/2)
		copy(newslice, stack.values)
		stack.values = newslice
	}
	return str
}

func (stack *StackS) isEmpty() (bool) {
	return stack.current == 0
}
