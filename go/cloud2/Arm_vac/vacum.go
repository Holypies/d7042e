package main

type Vacum struct {
	off State
	on  State

	currentState State
}

func newVacum() *Vacum {
	v := &Vacum{}

	offState := &offState{
		vac: v,
	}
	onState := &onState{
		vac: v,
	}

	v.off = offState
	v.on = onState

	v.setState(offState)

	return v
}

func (p *Vacum) setState(s State) {
	p.currentState = s
}

func (v *Vacum) getCurrentStateAsString() string {
	currState := "state"
	switch v.currentState {
	case v.on:
		currState = "ON"
	case v.off:
		currState = "OFF"
	default:
		currState = "error"
	}
	return currState
}

func (v *Vacum) turnOn() error {
	return v.currentState.turnOn()
}
func (v *Vacum) turnOff() error {
	return v.currentState.turnOff()
}
