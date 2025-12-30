package main

import "fmt"

type clockwiseState struct {
	arm *Arm
}

func (s *clockwiseState) atStart() error {
	s.arm.setState(s.arm.start)
	fmt.Println("Changing state to: start")
	return nil
}
func (s *clockwiseState) moveClockwise() error {
	return fmt.Errorf("already moving clockwise")
}
func (s *clockwiseState) moveCounterClockwise() error {
	return fmt.Errorf("can't change direction")
}
func (s *clockwiseState) atPickup() error {
	return fmt.Errorf("not at start")
}
