package stack_go

/*
非线程安全的
*/
type Stack struct {
	top  int
	data []int
}

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Init(size int) {
	if size <= 0 {
		panic("size is too little!!")
	}
	s.top = 0
	s.data = make([]int, size, size)
}

func (s Stack) Size() int {
	return cap(s.data)
}

func (s Stack) Length() int {
	return s.top
}

func (s Stack) IsEmpty() bool {
	return s.top <= 0
}

func (s Stack) IsFull() bool {
	return s.top > s.Size()
}

func (s *Stack) Push(v int) bool {
	if !s.IsFull() {
		s.data[s.top] = v
		s.top++
		return true
	}
	return false
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	v := s.data[s.top-1]
	s.top--
	return v, true
}

func (s *Stack) Seek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	v := s.data[s.top-1]
	return v, true
}
