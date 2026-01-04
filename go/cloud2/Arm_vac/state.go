package main

type State interface {
	turnOn() error
	turnOff() error

	/* not needed for simple demonstation
	stop() error
	reset() error */
}
