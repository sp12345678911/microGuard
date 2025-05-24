package greet

import "testing"

func TestGreeting(t *testing.T) {
	result := Greeting("測試")
	expected := "你好, 測試!"

	if result != expected {
		t.Errorf("Greeting() = %q, 預期 %q", result, expected)
	}
}
