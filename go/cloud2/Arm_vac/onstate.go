package main

import "fmt"

type onState struct {
	vac *Vacum
}

func (s *onState) turnOn() error {
	return fmt.Errorf("allready on")
}
func (s *onState) turnOff() error {
	s.vac.setState(s.vac.off)
	fmt.Println("turning off")
	return nil
}
