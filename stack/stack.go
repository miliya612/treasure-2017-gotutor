package main

import "fmt"

// Stack represents a structure of stack.
type Stack struct {
	content []string
	limit   int
}

// Pop remove the last item in Stack.
func (s *Stack) Pop() string {
	if size := len(s.content); size > 0 {
		str := s.content[size-1]
		s.content = s.content[:size-1]
		return str
	}
	return "\"\""
}

// Push add the arg to Stack as the last item.
func (s *Stack) Push(cc string) {
	if len(s.content) >= s.limit {
		s.content = s.content[len(s.content)-s.limit+1:]
	}
	s.content = append(s.content, cc)
}

func main() {
	s := &Stack{limit: 2}
	s.Push("dataA")
	s.Push("dataB")
	s.Push("dataC")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	s.Push("dataD")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
