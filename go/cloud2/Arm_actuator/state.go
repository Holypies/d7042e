package main

type State interface {
	atStart() error
	moveCounterClockwise() error
	moveClockwise() error
	atPickup() error

	/* not needed for simple demonstation
	stop() error
	reset() error */
}
