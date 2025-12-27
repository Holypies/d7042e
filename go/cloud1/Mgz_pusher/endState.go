package main 

import "fmt"

type endState struct{
	piston *Piston
}

func (s *endState)push() error {
	return fmt.Errorf("In end state, can not push further")
} 
func (s *endState)retract() error {
	s.piston.setState(s.piston.retract)
	fmt.Println("Changing state to: Retracting")
	return nil 
} 
func (s *endState)stop() error {
	s.piston.setState(s.piston.stop)
	fmt.Println("Changing state to: Stop")
	return nil 
} 
func (s *endState)reset() error {
	return fmt.Errorf("Stop before reseting")
} 

func (s *endState)atStart() error {
	return fmt.Errorf("invald command, is not at start state")
} 
func (s *endState)atEnd() error {
	return fmt.Errorf("invald command, is already at end state")
} 