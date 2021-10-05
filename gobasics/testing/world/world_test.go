package world

import "testing"

// TestHello checks if "Hello World" is returned
func TestWorld(t *testing.T) {
	got := World()
	if got != "World" {
		t.Errorf("got: %v expected: \"Hello\n", got)
	}
}
