package main

import (
	"fmt"
	"log"
)

func test() {
	arm := newArm()

	fmt.Println("at state: ", arm.getCurrentStateAsString())
	err := arm.moveCounterClockwise()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("at state: ", arm.getCurrentStateAsString())
	err = arm.atPickup()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("at state: ", arm.getCurrentStateAsString())
	err = arm.moveClockwise()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("at state: ", arm.getCurrentStateAsString())
	err = arm.atStart()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("at state: ", arm.getCurrentStateAsString())
	fmt.Println()

	/* err = vendingMachine.requestItem()
	   if err != nil {
	       log.Fatalf(err.Error())
	   }

	   err = vendingMachine.insertMoney(10)
	   if err != nil {
	       log.Fatalf(err.Error())
	   }

	   err = vendingMachine.dispenseItem()
	   if err != nil {
	       log.Fatalf(err.Error())
	   } */
}
