package main

/*
#include "stack_c.h"
*/
import "C"

import (
	"fmt"
	. "github.com/toukii/equal"
)

//export Display
func Display(stk C.struct_Stack) {
	fmt.Printf("Stack size:%d\n", stk.size)
	length := C.before(stk.top, stk.base)
	if length <= 0 {
		return
	}
	fmt.Print("[")
	for {
		if length <= 0 {
			break
		}
		length--
		fmt.Printf("%d ", C.index(stk.base, length))
	}
	fmt.Println("]")
}

func TestStack() {
	var stk C.struct_Stack
	C.Init(&stk, 11)
	C.Display(stk)
	ok := C.Push(&stk, 1)
	fmt.Println("ok:", ok)
	if ok == 1 {
		fmt.Println("...")
	}
	C.Display(stk)
}

func TestInit() {
	var stk C.struct_Stack
	C.Init(&stk, 3)
	Equal(nil, 2, C.Size(stk))
	// Equal(t, true, stk.IsEmpty())
}

func TestPush() {
	var stk C.struct_Stack
	C.Init(&stk, 2)
	C.Push(&stk, 1)
	C.Push(&stk, 2)
	Equal(nil, 2, C.Length(stk))
	var ok C.bool
	v := C.Seek(stk, &ok)
	Equal(nil, 2, v, true, ok, 2, C.Length(stk))
	ok = C.Push(&stk, 3)
	Equal(nil, 2, C.Length(stk), false, ok)
	v = C.Pop(&stk, &ok)
	Equal(nil, 2, v, true, ok, 1, C.Length(stk))
}

func main() {
	// TestInit()
	TestStack()
}
