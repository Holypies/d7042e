package main

import "fmt"

type offState struct {
	vac *Vacum
}

func (s *offState) turnOn() error {
	s.vac.setState(s.vac.on)
	fmt.Println("turning on")
	return nil
}
func (s *offState) turnOff() error {
	return fmt.Errorf("allready")
}
