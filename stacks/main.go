package main

type stack struct {
	items []int
}

func (s *stack) push(value int) {
	s.items = append(s.items, value)
}

func (s *stack) pop(value int) int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {
	stack := stack{}
	stack.push(10)
	stack.pop(10)
}
