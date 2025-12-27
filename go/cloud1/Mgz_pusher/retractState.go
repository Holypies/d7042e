package main 

import "fmt"

type retractState struct{
	piston *Piston
}

func (s *retractState)push() error {
	return fmt.Errorf("invalid command")
} 
func (s *retractState)retract() error {
	return fmt.Errorf("already retracting")
} 
func (s *retractState)stop() error {
	s.piston.setState(s.piston.stop)
	fmt.Println("Changing state to: Stop")
	return nil 
} 
func (s *retractState)reset() error {
	return fmt.Errorf("Stop before reseting")
} 

func (s *retractState)atStart() error {
	s.piston.setState(s.piston.start)
	fmt.Println("Changing state to: start")
	return nil 
} 

func (s *retractState)atEnd() error {
	return fmt.Errorf("invald command, is not at end state")
} 