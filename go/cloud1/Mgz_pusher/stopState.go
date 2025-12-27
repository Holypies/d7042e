package main 

import "fmt"

type stopState struct{
	piston *Piston
}

func (s *stopState)push() error {
	return fmt.Errorf("In stop state, please reset before proceeding")
} 
func (s *stopState)retract() error {
	return fmt.Errorf("In stop state, please reset before proceeding")
} 
func (s *stopState)stop() error {
	return fmt.Errorf("In stop state, please reset before proceeding")
} 
func (s *stopState)reset() error {
	s.piston.setState(s.piston.retract)
	return nil
} 
func (s *stopState)atStart() error {
	return fmt.Errorf("In stop state, please reset before proceeding")
} 
func (s *stopState)atEnd() error {
	return fmt.Errorf("In stop state, please reset before proceeding")
} 