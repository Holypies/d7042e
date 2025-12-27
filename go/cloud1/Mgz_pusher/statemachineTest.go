package main

import (
    "fmt"
    "log"
)

func test() {
    piston := newPiston()

    err := piston.startPush()
    if err != nil {
        log.Fatalf(err.Error())
    }

    err = piston.atEnd()
    if err != nil {
        log.Fatalf(err.Error())
    }

    err = piston.startRetract()
    if err != nil {
        log.Fatalf(err.Error())
    }


    err = piston.atStart()
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