package stack_test

import (
	"go-algorithms/stack"
	"testing"
)

func TestIsValidClosure(t *testing.T) {
	input := "{[]()}"
	if stack.IsValidClosure(input) != true {
		t.Fatalf("Failed test %s", input)
	}
	input = "{[(]}"

	if stack.IsValidClosure(input) != false {
		t.Fatalf("Failed test %s", input)
	}

	input = "{[}"
	if stack.IsValidClosure(input) != false {
		t.Fatalf("Failed test %s", input)
	}
}
