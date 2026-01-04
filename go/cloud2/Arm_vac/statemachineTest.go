package main

import (
	"fmt"
	"log"
)

func test() {
	vac := newVacum()

	err := vac.turnOn()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vac.turnOff()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vac.turnOn()
	if err != nil {
		log.Fatalf(err.Error())
	}

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
