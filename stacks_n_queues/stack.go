package stack

type Stack interface {
	push(str string)
	pop() (string)
	isEmpty() (bool)
	size() (int)
}
