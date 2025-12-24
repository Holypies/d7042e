package main

import (
	"encoding/json"
	"fmt"

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

func main() {
    framework, err := arrowhead.CreateFramework()
    checkError(err)

    // For POST services: never use EmptyParams() directly
    params := arrowhead.EmptyParams()
		car := Sensor{Status: "False"}

		carJSON, err := json.Marshal(car)
		checkError(err)
		params.Payload = carJSON
    // Toggle the sensor
    _, err = framework.SendRequest("update-sensor-status", params)
    checkError(err)
    fmt.Println("Sensor toggled successfully")

    // Now GET the current status
    res, err := framework.SendRequest("get-sensor-status", arrowhead.EmptyParams())
    checkError(err)

    var sensor Sensor
    err = json.Unmarshal(res, &sensor)
    checkError(err)
    fmt.Printf("Current sensor status: %s\n", sensor.Status)
}
