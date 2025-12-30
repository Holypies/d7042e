package main

import "fmt"

type pickupState struct {
	arm *Arm
}

func (s *pickupState) atStart() error {
	return fmt.Errorf("not at start")
}
func (s *pickupState) moveCounterClockwise() error {
	return fmt.Errorf("can't move more clockwise")
}
func (s *pickupState) moveClockwise() error {
	s.arm.setState(s.arm.clockwise)
	fmt.Println("Changing state to: clockwise")
	return nil
}
func (s *pickupState) atPickup() error {
	return fmt.Errorf("not at the pickup point")
}
