package stack_go

import (
	. "github.com/toukii/equal"
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
	Equal(t, 2, v, true, ok, 2, stack.Length())
	ok = stack.Push(3)
	Equal(t, 2, stack.Length(), false, ok)
	v, ok = stack.Pop()
	Equal(t, 2, v, true, ok, 1, stack.Length())
}

// 测试用例相互依赖，重用测试用例 -- 测试用例相互调用
// 一次调用Assert，判断多个条件：Equal(t, 2, v, true, ok)
