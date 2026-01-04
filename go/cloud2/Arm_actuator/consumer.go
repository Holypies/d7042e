package main

import (
	"encoding/json"
	"fmt"
	"time"

	arrowhead "github.com/johankristianss/arrowhead/pkg/arrowhead"
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

		/* check vacum sensor (allways true when testing without cloud 2)
			resVac, err := framework.SendRequest("get-sensor-status", arrowhead.EmptyParams())
		    checkError(err)
		    err = json.Unmarshal(resVac, &sensor.vacumSensor)
		    checkError(err) */

		ps1Status := sensors.pSensor1.Status
		ps2Status := sensors.pSensor2.Status
		esStatus := sensors.eSensor.Status
		vacStatus := sensors.vacumSensor.Status

		// check for start
		if esStatus == "True" && ps1Status == "True" {
			err = pusher.startPush()
			if err != nil {
				//log.Fatalf(err.Error())
			}
			//fmt.Println("Poll successful at start")
		}

		if ps2Status == "True" {
			err = pusher.atEnd()
			if err != nil {
				//log.Fatalf(err.Error())
			}
			//fmt.Println("Poll successful at push")
		}
		if vacStatus == "True" {
			err = pusher.startRetract()
			if err != nil {
				//log.Fatalf(err.Error())
			}
			//fmt.Println("Poll successful at end")
		}
		if esStatus == "False" && ps1Status == "True" {
			err = pusher.atStart()
			if err != nil {
				//log.Fatalf(err.Error())
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
		vacumSensor: Sensor{Status: "True"},
	}

	piston := newPiston()

	go pollingService(*framework, &sensors, piston)

	err = framework.ServeForever()
	checkError(err)

}
