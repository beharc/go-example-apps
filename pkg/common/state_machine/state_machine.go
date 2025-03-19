package state_machine

import (
	"fmt"
	"sync"
)

type StateMachine struct {
	currentState     string
	validTransitions map[string][]string
	version          int64
	mu               sync.RWMutex
}

func NewStateMachine(initialState string, transitions map[string][]string) *StateMachine {
	return &StateMachine{
		currentState:     initialState,
		validTransitions: transitions,
	}
}

func (sm *StateMachine) GetStateAndVersion() (string, int64) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.currentState, sm.version
}

func (sm *StateMachine) GetState() string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.currentState
}

func (sm *StateMachine) CanTransition(newState string) bool {
	for _, validState := range sm.validTransitions[sm.currentState] {
		if validState == newState {
			return true
		}
	}
	return false
}

func (sm *StateMachine) Transition(newState string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if sm.CanTransition(newState) {
		sm.currentState = newState
		sm.version++
		return nil
	}
	return fmt.Errorf("invalid transition from %s to %s", sm.currentState, newState)
}
