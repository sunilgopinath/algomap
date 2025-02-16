package minstack_test

import (
	"testing"

	minstack "github.com/sunilgopinath/algomap/stacks/155MinStack"
)

func TestMinStack(t *testing.T) {
	// Initialize MinStack
	obj := minstack.Constructor()

	// Test case: ["push", "push", "push", "getMin", "pop", "top", "getMin"]
	// Inputs: [[-2],[0],[-3],[],[],[],[]]
	// Expected Outputs: [null,null,null,null,-3,null,0,-2]

	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)

	// Check getMin
	if min := obj.GetMin(); min != -3 {
		t.Errorf("expected GetMin to return -3, but got %d", min)
	}

	// Pop the top element
	obj.Pop()

	// Check top
	if top := obj.Top(); top != 0 {
		t.Errorf("expected Top to return 0, but got %d", top)
	}

	// Check getMin again
	if min := obj.GetMin(); min != -2 {
		t.Errorf("expected GetMin to return -2, but got %d", min)
	}
}
