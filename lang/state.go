package lang

import (
	"fmt"
	"github.com/jfixby/pin"
)

type StateChecker interface {
	CheckStateIs(state string)
	CheckStateIsNot(state string)
	SwitchState(nextState string)
	SwitchStateFromTo(stateFrom, stateTo string)
}

type basicStateChecker struct {
	state string
}

func (b *basicStateChecker) CheckStateIs(state string) {
	pin.AssertTrue(fmt.Sprintf("Current state<%v> is <%v>", b.state, state), b.state == state)
}

func (b *basicStateChecker) CheckStateIsNot(state string) {
	pin.AssertTrue(fmt.Sprintf("Current state<%v> is not <%v>", b.state, state), b.state != state)
}

func (b *basicStateChecker) SwitchState(nextState string) {
	AssertNotNil("nextState", nextState)
	AssertNotEmpty("nextState", nextState)
	b.state = nextState
}

func (b *basicStateChecker) SwitchStateFromTo(stateFrom, stateTo string) {
	b.CheckStateIs(stateFrom)
	b.SwitchState(stateTo)
}

func NewStateChecker(startState string) StateChecker {
	AssertNotNil("startState", startState)
	AssertNotEmpty("startState", startState)
	return &basicStateChecker{state: startState}
}
