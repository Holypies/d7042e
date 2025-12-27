package main

type State interface {
	atStart() error
	push() error
	retract() error
	atEnd() error

	stop() error
	reset() error
}
