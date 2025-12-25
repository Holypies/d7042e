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

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func pollingService(framework arrowhead.Framework,sensor *Sensor){
	
	for {

    res, err := framework.SendRequest("get-sensor-status", arrowhead.EmptyParams())
    checkError(err)
    
    err = json.Unmarshal(res, &sensor)
    checkError(err)
    fmt.Printf("Current sensor status: %s\n", sensor.Status)
		time.Sleep(1*time.Second)
		
	}


}



func main() {
    framework, err := arrowhead.CreateFramework()
    checkError(err)

    var sensor Sensor
  	go pollingService(*framework,&sensor)
    

    err = framework.ServeForever()
	  checkError(err)

}
