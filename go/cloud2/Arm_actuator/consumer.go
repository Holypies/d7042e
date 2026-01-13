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
	pSensor2    Sensor
	aSensor1    Sensor
	aSensor2    Sensor
	vacumSensor Sensor
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func pollingService(framework arrowhead.Framework, sensors *Sensors, arm *Arm) {
	fmt.Println("Current state: ", arm.getCurrentStateAsString())
	for {
		
			resPsen2, err := framework.SendRequest("psen2-get-status", arrowhead.EmptyParams())
			checkError(err)
			err = json.Unmarshal(resPsen2, &sensors.pSensor2)
			checkError(err)
		

		resAsen1, err := framework.SendRequest("asen1-get-status", arrowhead.EmptyParams())
		checkError(err)
		err = json.Unmarshal(resAsen1, &sensors.aSensor1)
		checkError(err)

		resAsen2, err := framework.SendRequest("asen2-get-status", arrowhead.EmptyParams())
		checkError(err)
		err = json.Unmarshal(resAsen2, &sensors.aSensor2)
		checkError(err)

		resVac, err := framework.SendRequest("vac-get-status", arrowhead.EmptyParams())
		checkError(err)
		err = json.Unmarshal(resVac, &sensors.vacumSensor)
		checkError(err)

		ps2Status := sensors.pSensor2.Status
		as1Status := sensors.aSensor1.Status
		as2Status := sensors.aSensor2.Status
		vacStatus := sensors.vacumSensor.Status

		old_state := arm.getCurrentStateAsString()
		if ps2Status == "True" && vacStatus == "False" {
			err = arm.moveCounterClockwise()
			if err != nil {
				//log.Fatalf(err.Error())
			}
			//fmt.Println("intra cloud success")
		}

		if as1Status == "True" {
			err = arm.atPickup()
			// start vac
			if err != nil {
				//log.Fatalf(err.Error())
			}
			//fmt.Println("Poll successful at start")
		}

		if vacStatus == "True" {
			err = arm.moveClockwise()
			if err != nil {
				//log.Fatalf(err.Error())
			}
			//fmt.Println("Poll successful at push")
		}
		if as2Status == "True" {
			err = arm.atStart()
			// vacum off
			if err != nil {
				//log.Fatalf(err.Error())
			}
			//fmt.Println("Poll successful at end")
		}

		if arm.getCurrentStateAsString() != old_state {
			fmt.Println("Current state: ", arm.getCurrentStateAsString())
		}

		time.Sleep(1000 * time.Millisecond)

	}

}

func main() {
	framework, err := arrowhead.CreateFramework()
	checkError(err)

	sensors := Sensors{
		pSensor2:    Sensor{Status: "False"},
		aSensor1:    Sensor{Status: "False"},
		aSensor2:    Sensor{Status: "False"},
		vacumSensor: Sensor{Status: "False"},
	}

	arm := newArm()

	go pollingService(*framework, &sensors, arm)

	err = framework.ServeForever()
	checkError(err)

}
