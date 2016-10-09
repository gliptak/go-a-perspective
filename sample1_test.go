package main

import "testing"

func TestAddition(t *testing.T) {
	v := 1.0 + 2.0
	if v != 3.0 {
		t.Error("Expected 3, got ", v)
	}
}

func TestAdditionFail(t *testing.T) {
	v := 1.0 + 2.0
	if v != 2.0 {
		t.Error("Expected 2, got ", v)
	}
}
