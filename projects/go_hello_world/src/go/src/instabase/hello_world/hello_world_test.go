package go_hello_world

import "testing"

func TestCreeter(t *testing.T) {
	expected := "Hello World!"
	actual := HelloWorld()
	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}
