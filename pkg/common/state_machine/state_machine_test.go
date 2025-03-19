package state_machine

import (
	"testing"
)

// Test the state machine transitions
func TestStateMachine(t *testing.T) {
	// Define allowed transitions
	transitions := map[string][]string{
		"Created":    {"Processing"},
		"Processing": {"Completed"},
	}

	// Initialize state machine
	sm := NewStateMachine("Created", transitions)

	// Test valid transition
	err := sm.Transition("Processing")
	if err != nil {
		t.Errorf("Expected valid transition, got error: %v", err)
	}

	// Test another valid transition
	err = sm.Transition("Completed")
	if err != nil {
		t.Errorf("Expected valid transition, got error: %v", err)
	}

	// Test invalid transition
	err = sm.Transition("Created")
	if err == nil {
		t.Errorf("Expected error for invalid transition, but got none")
	}
}
