package main

/*
#include "stack_c.h"
*/
import "C"

import (
	"fmt"
	. "github.com/toukii/assert"
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
		fmt.Printf("%d ", C.indexing(stk.base, length))
	}
	fmt.Println("]")
}

func TestStack() {
	var stk C.struct_Stack
	C.Init(&stk, 11)
	C.Display(stk)
	ok := C.Push(&stk, 1)
	Equal(nil, C.int(1), C.Length(stk), C.bool(1), ok)
	C.Display(stk)
}

func TestInit() {
	var stk C.struct_Stack
	C.Init(&stk, 3)
	Equal(nil, C.int(3), C.Capacity(stk))
	Equal(nil, C.bool(1), C.IsEmpty(stk))
}

func TestPush() {
	var stk C.struct_Stack
	C.Init(&stk, 2)
	C.Push(&stk, 1)
	C.Push(&stk, 2)
	Equal(nil, C.int(2), C.Length(stk))
	var ok C.bool
	v := C.Seek(stk, &ok)
	Equal(nil, C.int(2), v, C.bool(1), ok, C.int(2), C.Length(stk))
	ok = C.Push(&stk, 3)
	Equal(nil, C.int(2), C.Length(stk), C.bool(0), ok)
	v = C.Pop(&stk, &ok)
	Equal(nil, C.int(2), v, C.bool(1), ok, C.int(1), C.Length(stk))
}

func main() {
	TestInit()
	TestPush()
	TestStack()
}
