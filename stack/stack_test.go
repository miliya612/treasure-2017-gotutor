package main

import "testing"

func TestPush(t *testing.T) {
	type Case struct {
		in, want []string
	}
	cs := []Case{
		{
			[]string{"a"},
			[]string{"a"},
		},
		{
			[]string{"a", "b", "c"},
			[]string{"a", "b", "c"},
		},
	}
	for i, c := range cs {
		s := &Stack{}
		for _, str := range c.in {
			s.Push(str)
		}
		got := s
		if !testEq(c.want, got.content) {
			t.Errorf("#%d: Push(%#v) want %#v, got %#v\n", i, c.in, c.want, got)
		}
	}
}

func TestPop(t *testing.T) {
	type Case struct {
		in   []string
		want string
	}
	cs := []Case{
		{
			[]string{"a"},
			"a",
		},
		{
			[]string{"a", "b", "c"},
			"c",
		},
	}
	for i, c := range cs {
		s := &Stack{}

		// initialize: add item to Stack
		for _, str := range c.in {
			s.Push(str)
		}
		got := s.Pop()
		if got != c.want {
			t.Errorf("#%d: Pop() want %#v, got %#v\n", i, c.want, got)
		}
	}
}

func testEq(a, b []string) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
