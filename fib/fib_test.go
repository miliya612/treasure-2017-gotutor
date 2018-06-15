package main

import "testing"

func TestFib(t *testing.T) {
	type Case struct {
		in, want int
	}

	cases := []Case{
		{0, 0},
		{1, 1},
		{10, 55},
		{-1, 0},
	}
	for i, c := range cases {
		if got := fib(c.in); got != c.want {
			t.Errorf("#%d: fib(%d) want %d, got %d\n", i, c.in, c.want, got)
		}
	}
}
