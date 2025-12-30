package main

type Arm struct {
	start            State
	counterClockwise State
	pickup           State
	clockwise        State
	stop             State

	currentState State
}

func newArm() *Arm {
	a := &Arm{}

	startState := &startState{
		arm: a,
	}
	clockwiseState := &clockwiseState{
		arm: a,
	}
	pickupState := &pickupState{
		arm: a,
	}
	counterClockwiseState := &counterClockwiseState{
		arm: a,
	}

	a.start = startState
	a.counterClockwise = counterClockwiseState
	a.clockwise = clockwiseState
	a.pickup = pickupState

	a.setState(startState)

	return a
}

func (a *Arm) setState(s State) {
	a.currentState = s
}

func (a *Arm) getCurrentStateAsString() string {
	currState := "state"
	switch a.currentState {
	case a.start:
		currState = "Start state"
	case a.clockwise:
		currState = "Clockwise state"
	case a.counterClockwise:
		currState = "Counterclockwise state"
	case a.pickup:
		currState = "Pickup state"
	default:
		currState = "error: unaccounted state"
	}
	return currState
}

func (a *Arm) atStart() error {
	return a.currentState.atStart()
}
func (a *Arm) moveClockwise() error {
	return a.currentState.moveClockwise()
}
func (a *Arm) atPickup() error {
	return a.currentState.atPickup()
}
func (a *Arm) moveCounterClockwise() error {
	return a.currentState.moveCounterClockwise()
}

/* func (a *Arm) startStop() error {
	return a.currentState.stop()
}
func (a *Arm) startReset() error {
	return a.currentState.reset()
} */
