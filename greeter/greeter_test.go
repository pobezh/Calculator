package greeter

import "testing"

func TestGreet(t *testing.T) {
	g := New()
	expected := "Hello, Guest! Welcome to the multi-file Go project from CodeInterview."
	if msg := g.Greet(); msg != expected {
		t.Errorf("Greet() = %v, want %v", msg, expected)
	}
}

func TestUpdateName(t *testing.T) {
	g := New()
	g.UpdateName("Alice")
	if g.Name != "Alice" {
		t.Errorf("UpdateName() failed, got %v, want %v", g.Name, "Alice")
	}
}

func TestIsNameValid(t *testing.T) {
	g := New()
	if !g.IsNameValid("Alice") {
		t.Errorf("IsNameValid() should be true for 'Alice'")
	}
	if g.IsNameValid("Alice123") {
		t.Errorf("IsNameValid() should be false for 'Alice123'")
	}
}
