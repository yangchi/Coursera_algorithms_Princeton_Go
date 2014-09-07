package stack

import (
	"testing"
	"math/rand"
)

func randomString (length uint8) (string) {
	const alphadigits = "0123456789abcdefghjiklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(123)
	buf := make([]byte, length)
	for index, _:= range buf {
		buf[index] = alphadigits[rand.Intn(len(alphadigits))]
	}
	return string(buf)
}

func genericRandomTest (stack Stack) (bool, string) {
	rounds := rand.Intn(100)
	for rounds > 0 {
		rounds--
		if rand.Int() % 2 == 1 {
			/* Push */
			str := randomString(uint8(rand.Int()))
			cur_length := stack.size()
			stack.push(str)
			if stack.size() != cur_length + 1 {
				return false, "Push failed on string: " + str
			}
		} else {
			/* Pop */
			if stack.isEmpty() {
				continue
			}
			cur_length := stack.size()
			str := stack.pop()
			if stack.size() != cur_length - 1 {
				return false, "Pop failed with string returned: " + str
			}
		}
	}
	return true, ""
}

func TestStackLRandomString (t *testing.T) {
	var stack Stack
	var stackL StackL
	stackL.init()
	stack = &stackL
	res, mes := genericRandomTest(stack)
	if !res {
		t.Error(mes)
	}
}

func genericSimpleTest (stack Stack) (bool, string) {
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
	res, msg := genericSimpleTest(stack)
	if !res {
		t.Error(msg)
	}
}
