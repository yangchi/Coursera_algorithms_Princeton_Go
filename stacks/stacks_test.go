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
	if !stack.isEmpty() {
		return false, "Should be empty to begin with"
	}
	rounds := rand.Intn(100)
	for rounds > 0 {
		rounds--
		/* Push */
		str := randomString(uint8(rand.Int()))
		stack.push(str)
		if stack.isEmpty() {
			return false, "Push failed"
		}
		/* Pop */
		str_popped := stack.pop()
		if !stack.isEmpty() {
			return false, "Pop failed"
		}
		if str != str_popped {
			return false, "Strings do not match"
		}
	}
	return true, ""
}

func genericSimpleTest (stack Stack) (bool, string) {
	if !stack.isEmpty() {
		return false, "isEmpty gives wrong answer after init"
	}
	stack.push("123")
	if stack.isEmpty() {
		return false, "Push failed"
	}
	str := stack.pop()
	if !stack.isEmpty() {
		return false, "Pop failed"
	}
	if str != "123" {
		return false, "Push and Pop do not match"
	}
	return true, ""
}

func TestStackLL (t *testing.T) {
	var stack Stack
	var stackLL StackLL
	stack = &stackLL
	res, msg := genericSimpleTest(stack)
	if !res {
		t.Error(msg)
	}
	res, msg = genericRandomTest(stack)
	if !res {
		t.Error(msg)
	}
}

func TestStackS (t *testing.T) {
	var stack Stack
	var stackS StackS
	stackS.init()
	stack = &stackS
	res, msg := genericSimpleTest(stack)
	if !res {
		t.Error(msg)
	}
	res, msg = genericRandomTest(stack)
	if !res {
		t.Error(msg)
	}
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
	res, msg = genericRandomTest(stack)
	if !res {
		t.Error(msg)
	}
}
