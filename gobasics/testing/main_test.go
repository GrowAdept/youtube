package main

import (
	"fmt"
	"testing"
)

var numTests = []struct {
	input    int
	expected int
}{
	{1, 3},
	{100, 102},
	{0, 2},
	{-10, -8},
}

func TestAdd2(t *testing.T) {
	for _, n := range numTests {
		got := add2(n.input)
		if got != n.expected {
			t.Errorf("got: %v expected: %v\n", got, n.expected)
		}
	}
}

func TestHello(t *testing.T) {
	got := Hello()
	if got != "Hello" {
		t.Errorf("got: \"%v\" expected: \"Hello\"\n", got)
	}
}

func BenchmarkAdd2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add2(-10)
	}
}

func Example() {
	fmt.Println("The bigger the interface, the weaker the abstraction.") // if this doesn't match comment below, it will fail
	// Output: The bigger the interface, the weaker the abstraction.
}
