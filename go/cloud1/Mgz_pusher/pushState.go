package main 

import "fmt"

type pushState struct{
	piston *Piston
}

func (s *pushState)atStart() error {
	return fmt.Errorf("invald command, not at start")
} 
func (s *pushState)push() error {
	return fmt.Errorf("Allready in push state")
} 
func (s *pushState)retract() error {
	return fmt.Errorf("Cannot retract before endState")
} 
func (s *pushState)atEnd() error {
	s.piston.setState(s.piston.end)
	fmt.Println("Changing state to: end")
	return nil 
} 
func (s *pushState)stop() error {
	s.piston.setState(s.piston.stop)
	fmt.Println("Changing state to: Stop")
	return nil 
} 
func (s *pushState)reset() error {
	return fmt.Errorf("Stop before reseting")
} 