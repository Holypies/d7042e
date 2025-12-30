package main

import "fmt"

type counterClockwiseState struct {
	arm *Arm
}

func (s *counterClockwiseState) atStart() error {
	return fmt.Errorf("not at start")
}
func (s *counterClockwiseState) moveClockwise() error {
	return fmt.Errorf("already moving clockwise")
}
func (s *counterClockwiseState) moveCounterClockwise() error {
	return fmt.Errorf("can't change direction")
}
func (s *counterClockwiseState) atPickup() error {
	s.arm.setState(s.arm.pickup)
	fmt.Println("Changing state to: pickup")
	return nil
}
