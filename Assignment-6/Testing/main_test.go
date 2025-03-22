package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(3, 2)
	expected := 5
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
