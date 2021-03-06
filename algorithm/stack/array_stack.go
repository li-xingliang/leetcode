package stack

type arrayStack struct {
	data []int

	cap    int // 容量
	length int // 大小
	top    int // 栈顶
}

// NewArrayStack new array stack
func NewArrayStack(cap int) Stack {
	return &arrayStack{
		data:   make([]int, cap, cap),
		cap:    cap,
		length: 0,
		top:    0,
	}
}

func (s *arrayStack) Push(data int) error {
	if s.Len() == s.cap {
		return ErrStackFull
	}

	s.data[s.top] = data
	s.length++
	s.top++

	return nil
}

func (s *arrayStack) Pop() (int, error) {
	if s.Len() == 0 {
		return 0, ErrStackEmpty
	}

	s.top--
	tmp := s.data[s.top]
	s.length--

	return tmp, nil
}

func (s *arrayStack) Len() int {
	return s.length
}

func (s *arrayStack) Empty() error {
	s.data = make([]int, 0)
	s.length = 0
	s.top = -1

	return nil
}
