package collections

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack[int]()
	if s == nil {
		t.Error("NewStack() should not return nil")
	}
	var v = 2
	s.Push(v)
	if s.Len() != 1 {
		t.Error("NewStack() should return a stack with length 1")
	}
	if s.Pop() != v {
		t.Errorf("NewStack() should return value %d", v)
	}
	if s.Len() != 0 {
		t.Error("NewStack() should return a stack with length 0")
	}
}
