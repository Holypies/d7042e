package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	arrowhead "github.com/johankristianss/arrowhead/pkg/arrowhead"
	"github.com/johankristianss/arrowhead/pkg/rpc"
)

type Sensor struct {
	Status string `json:"status"`
}

type Sensors struct {
	aSen1 Sensor
	aSen2 Sensor
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func pollingService(framework arrowhead.Framework, sensors *Sensors, vac *Vacum, vac_sen *InMemorySensorRepository) {
	fmt.Println("Current state: ", vac.getCurrentStateAsString())
	for {

		resaSen1, err := framework.SendRequest("asen1-get-status", arrowhead.EmptyParams())
		checkError(err)
		err = json.Unmarshal(resaSen1, &sensors.aSen1)
		checkError(err)

		resaSen2, err := framework.SendRequest("asen2-get-status", arrowhead.EmptyParams())
		checkError(err)
		err = json.Unmarshal(resaSen2, &sensors.aSen2)
		checkError(err)

		as1Status := sensors.aSen1.Status
		as2Status := sensors.aSen2.Status

		old_state := vac.getCurrentStateAsString()
		// check for start
		if as1Status == "True" {
			err = vac.turnOn()
			vac_sen.sensor.Status = "True"
			if err != nil {

			}
		}

		if as2Status == "True" {
			err = vac.turnOff()
			vac_sen.sensor.Status = "False"
			if err != nil {

			}

		}

		if vac.getCurrentStateAsString() != old_state {
			fmt.Println("Current state: ", vac.getCurrentStateAsString())
		}
		time.Sleep(1000 * time.Millisecond)

	}

}

type InMemorySensorRepository struct {
	sync.RWMutex
	sensor *Sensor
}
type GetVacumService struct {
	inMemorySensorRepository *InMemorySensorRepository
}

func (s *GetVacumService) HandleRequest(params *arrowhead.Params) ([]byte, error) {
	sensorJSON, err := json.Marshal(s.inMemorySensorRepository.sensor)
	if err != nil {
		return nil, err
	}
	return sensorJSON, nil
}

func main() {
	framework, err := arrowhead.CreateFramework()
	checkError(err)

	vacum := newVacum()

	sensors := Sensors{
		aSen1: Sensor{Status: "False"},
		aSen2: Sensor{Status: "False"},
	}

	inMemorySensorRepository := &InMemorySensorRepository{
		sensor: &Sensor{Status: "False"},
	}

	go pollingService(*framework, &sensors, vacum, inMemorySensorRepository)

	getVacumService := &GetVacumService{inMemorySensorRepository: inMemorySensorRepository}

	framework.HandleService(getVacumService, rpc.GET, "vac-get-status", "/arm")

	err = framework.ServeForever()
	checkError(err)

}
