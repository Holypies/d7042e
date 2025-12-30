package main

import "fmt"

type startState struct {
	arm *Arm
}

func (s *startState) atStart() error {
	return fmt.Errorf("already at start")
}
func (s *startState) moveClockwise() error {
	return fmt.Errorf("can't move more clockwise")
}
func (s *startState) moveCounterClockwise() error {
	s.arm.setState(s.arm.counterClockwise)
	fmt.Println("Changing state to: counterclockwise")
	return nil
}
func (s *startState) atPickup() error {
	return fmt.Errorf("not at the pickup point")
}

/*
func (s *startState) reset() error {
	return fmt.Errorf("invald command, is not at start state")
}
func (s *startState) stop() error {
	return fmt.Errorf("invald command, is already at end state")
}
*/
