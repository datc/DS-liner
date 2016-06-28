package stack_go

import (
	. "github.com/bmizerany/assert"
	"testing"
)

var (
	stack *Stack
)

func init() {
	stack = NewStack()
}

func TestInit(t *testing.T) {
	stack.Init(3)
	Equal(t, 3, stack.Size())
	Equal(t, true, stack.IsEmpty())
}

func TestPush(t *testing.T) {
	stack.Init(2)
	stack.Push(1)
	stack.Push(2)
	Equal(t, 2, stack.Length())
	v, ok := stack.Seek()
	Equal(t, 2, v, true, ok)
	Equal(t, 2, stack.Length())

	v, ok = stack.Pop()
	Equal(t, 2, v, true, ok)
	Equal(t, 1, stack.Length())
}

// 测试用例相互依赖，重用测试用例
// 一次调用Assert，判断多个条件：Equal(t, 2, v, true, ok)
