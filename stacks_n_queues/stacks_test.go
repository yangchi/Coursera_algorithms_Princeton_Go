package stack

import "testing"

func genericStackTest (stack Stack) (bool, string) {
	if !stack.isEmpty() {
		return false, "isEmpty gives wrong answer after init"
	}
	if stack.size() != 0 {
		return false, "Stack size should be 0 after init"
	}
	stack.push("123")
	if stack.size() != 1 {
		return false, "Push failed"
	}
	stack.pop()
	if stack.size() != 0 {
		return false, "Pop failed"
	}
	stack.push("123")
	if stack.size() != 1 {
		return false, "2nd Push failed"
	}
	if stack.isEmpty() {
		return false, "isEmpty gives wrong answer when not empty"
	}
	for i := 0; i < 1; i++ {
		stack.pop()
	}
	if !stack.isEmpty() {
		return false, "isEmpty gives wrong answer when empty"
	}
	return true, ""
}

func TestStackL (t *testing.T) {
	var stack Stack
	var stackL StackL
	stackL.init()
	stack = &stackL
	res, msg := genericStackTest(stack)
	if !res {
		t.Error(msg)
	}
}
