package main

import (
	"encoding/json"
	"fmt"
	"time"

	arrowhead "github.com/johankristianss/arrowhead/pkg/arrowhead"
	//"github.com/johankristianss/arrowhead/pkg/rpc"
)

type Sensor struct {
	Status string `json:"status"`
}

type Sensors struct {
	eSensor     Sensor
	pSensor1    Sensor
	pSensor2    Sensor
	vacumSensor Sensor
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func pollingService(framework arrowhead.Framework, sensors *Sensors, pusher *Piston) {

	for {

		resEsen, err := framework.SendRequest("esen-get-status", arrowhead.EmptyParams())
		checkError(err)
		err = json.Unmarshal(resEsen, &sensors.eSensor)
		checkError(err)

		resPsen1, err := framework.SendRequest("psen1-get-status", arrowhead.EmptyParams())
		checkError(err)
		err = json.Unmarshal(resPsen1, &sensors.pSensor1)
		checkError(err)

		resPsen2, err := framework.SendRequest("psen2-get-status", arrowhead.EmptyParams())
		checkError(err)
		err = json.Unmarshal(resPsen2, &sensors.pSensor2)
		checkError(err)

		resVac, err := framework.SendRequest("vac-get-status", arrowhead.EmptyParams())
		checkError(err)
		err = json.Unmarshal(resVac, &sensors.vacumSensor)
		checkError(err)

		ps1Status := sensors.pSensor1.Status
		ps2Status := sensors.pSensor2.Status
		esStatus := sensors.eSensor.Status
		vacStatus := sensors.vacumSensor.Status

		// check for start
		if esStatus == "True" && ps1Status == "True" {
			err = pusher.startPush()
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Println("Poll successful at start")
		}

		if ps2Status == "True" {
			err = pusher.atEnd()
			if err != nil {
				fmt.Println(err)
			}
		}
		if vacStatus == "True" && ps2Status == "True" {
			err = pusher.startRetract()
			if err != nil {
				fmt.Println(err)
			}

		}
		if esStatus == "False" && ps1Status == "True" {
			err = pusher.atStart()
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Println("Poll successful at return")
		}
		fmt.Println("Current state: ", pusher.getCurrentStateAsString())
		time.Sleep(1000 * time.Millisecond)

	}

}

func main() {
	framework, err := arrowhead.CreateFramework()
	checkError(err)

	sensors := Sensors{
		eSensor:     Sensor{Status: "False"},
		pSensor1:    Sensor{Status: "False"},
		pSensor2:    Sensor{Status: "False"},
		vacumSensor: Sensor{Status: "False"},
	}

	piston := newPiston()

	go pollingService(*framework, &sensors, piston)

	err = framework.ServeForever()
	checkError(err)

}
