package main

type Piston struct {
	start   State
	push    State
	retract State
	end     State
	stop    State

	currentState State
}

func newPiston() *Piston {
	p := &Piston{}

	startState := &startState{
		piston: p,
	}
	pushState := &pushState{
		piston: p,
	}
	endState := &endState{
		piston: p,
	}
	retractState := &retractState{
		piston: p,
	}
	stopState := &stopState{
		piston: p,
	}

	p.start = startState
	p.retract = retractState
	p.push = pushState
	p.stop = stopState
	p.end = endState

	p.setState(startState)

	return p
}

func (p *Piston) setState(s State) {
	p.currentState = s
}

func (p *Piston) getCurrentStateAsString() string {
	currState := "state"
	switch p.currentState {
	case p.start:
		currState = "Start state"
	case p.retract:
		currState = "Retract state"
	case p.end:
		currState = "End state"
	case p.push:
		currState = "Push state"
	default:
		currState = "error"
	}
	return currState
}

func (p *Piston) atStart() error {
	return p.currentState.atStart()
}
func (p *Piston) startPush() error {
	return p.currentState.push()
}
func (p *Piston) atEnd() error {
	return p.currentState.atEnd()
}
func (p *Piston) startRetract() error {
	return p.currentState.retract()
}
func (p *Piston) startStop() error {
	return p.currentState.stop()
}
func (p *Piston) startReset() error {
	return p.currentState.reset()
}
