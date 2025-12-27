package main 

import "fmt"

type startState struct{
	piston *Piston
}


func (s *startState)atStart() error {
	return fmt.Errorf("invald command, is already at start state")
} 
func (s *startState)push() error {
	s.piston.setState(s.piston.push)
	fmt.Println("Changing state to: Push")
	return nil 
} 
func (s *startState)retract() error {
	return fmt.Errorf("In start state, can not retract further")
} 
func (s *startState)stop() error {
	s.piston.setState(s.piston.stop)
	fmt.Println("Changing state to: Stop")
	return nil 
} 
func (s *startState)reset() error {
	return fmt.Errorf("Already reset")
} 

func (s *startState)atEnd() error {
	return fmt.Errorf("invald command, is not at end state")
} 