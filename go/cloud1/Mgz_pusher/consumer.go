package main

import (
	"encoding/json"
	"fmt"
    "time"
	"log"

	arrowhead "github.com/johankristianss/arrowhead/pkg/arrowhead"
)

type Sensor struct {
	Status string `json:"status"`
}

type Sensors struct{
	eSensor Sensor
	pSensor1 Sensor
	pSensor2 Sensor
	vacumSensor Sensor

} 

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func pollingService(framework arrowhead.Framework,sensors *Sensors, pusher *Piston){
	
	for {

    resEsen, err := framework.SendRequest("get-sensor-status", arrowhead.EmptyParams())
    checkError(err)
    err = json.Unmarshal(resEsen, &sensors.eSensor)
    checkError(err)

	resPsen1, err := framework.SendRequest("get-sensor-status", arrowhead.EmptyParams())
    checkError(err)
    err = json.Unmarshal(resPsen1, &sensors.pSensor1)
    checkError(err)

	resPsen2, err := framework.SendRequest("get-sensor-status", arrowhead.EmptyParams())
    checkError(err)
    err = json.Unmarshal(resPsen2, &sensors.pSensor2)
    checkError(err)

	/* check vacum sensor (allways true when testing without cloud 2) 
	resVac, err := framework.SendRequest("get-sensor-status", arrowhead.EmptyParams())
    checkError(err)
    err = json.Unmarshal(resVac, &sensor.vacumSensor)
    checkError(err) */

	// check for start
	if resEsen == "True" && resPsen1 == "True"{
		err = &piston.startPush()
		if err != nil {
			log.Fatalf(err.Error())
    	}
		fmt.Println("Poll successful at start")
	}

	if resPsen2 == "True"{
		err = &piston.atEnd()
		if err != nil {
			log.Fatalf(err.Error())
    	}
		fmt.Println("Poll successful at push")
	}
	if resVac == "True"{
		err = &piston.startRetract()
		if err != nil {
			log.Fatalf(err.Error())
    	}
		fmt.Println("Poll successful at end")
	}
	if resEsen == "False" && resPsen1 == "True" {
		err = &piston.atStart()
		if err != nil {
			log.Fatalf(err.Error())
    	}
		fmt.Println("Poll successful at return")
	}


}



func main(){
    framework, err := arrowhead.CreateFramework()
    checkError(err)

    sensors := Sensors{
        eSensor:     Sensor{Status: "False"},
        pSensor1:    Sensor{Status: "False"},
        pSensor2:    Sensor{Status: "False"},
        vacumSensor: Sensor{Status: "True"},
    }

	piston := newPiston()

  	go pollingService(*framework,&sensors, &piston)
    

    err = framework.ServeForever()
	  checkError(err)

}
