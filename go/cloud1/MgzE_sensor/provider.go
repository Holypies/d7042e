package main

import (
	"encoding/json"
	"fmt"

	arrowhead "github.com/johankristianss/arrowhead/pkg/arrowhead"
	"github.com/johankristianss/arrowhead/pkg/rpc"
)

type Sensor struct {
	Status string `json:"status"`
}

type InMemorySensorRepository struct {
	sensor *Sensor
}

type UpdateSensorService struct {
	inMemorySensorRepository *InMemorySensorRepository
}

//func (s *UpdateSensorService) HandleRequest(params *arrowhead.Params) ([]byte, error) {
//	fmt.Println("UpdateSesorService called, updating sensor status")
//	sensor := Sensor{}
//	err := json.Unmarshal(params.Payload, &sensor)
//	fmt.Println(sensor)
//	if err != nil {
//		return nil, err
//	}
//	if sensor.Status == "True"{
//		s.inMemorySensorRepository.sensor.Status = "False"
//	}else{
//		s.inMemorySensorRepository.sensor.Status = "True"
//	}
//	
//	fmt.Println("test")
//	fmt.Println(s.inMemorySensorRepository.sensor.Status)
//	//s.inMemoryCarRepository.sensor = ""
//	return nil, nil
//}
func (s *UpdateSensorService) HandleRequest(params *arrowhead.Params) ([]byte, error) {
	fmt.Println("➡️ UpdateSensorService called, triggering status toggle.")

	// Check for proper initialization (Guardrail against nil pointer panic)
	if s.inMemorySensorRepository.sensor == nil {
		return nil, fmt.Errorf("repository sensor is not initialized (nil)")
	}

	// 1Println. Get the current status
	currentStatus := s.inMemorySensorRepository.sensor.Status
  fmt.Println("1")
	// 2. Toggle the status based on its current value
	var newStatus string
	if currentStatus == "True" {
		newStatus = "False"
	} else {
		// This handles "False" and the initial zero-value state ""
		newStatus = "True"
	}
	fmt.Println("2")
	// 3. Update the persistent state in the repository
	s.inMemorySensorRepository.sensor.Status = newStatus

	fmt.Printf("✅ Sensor Status Updated: Toggled from **%s** to **%s**\n", currentStatus, newStatus)

	// In Arrowhead, returning nil, nil indicates a successful service execution
	return nil, nil
}
type GetSensorService struct {
	inMemorySensorRepository *InMemorySensorRepository
}

func (s *GetSensorService) HandleRequest(params *arrowhead.Params) ([]byte, error) {
	sensorJSON, err := json.Marshal(s.inMemorySensorRepository.sensor)
	if err != nil {
		return nil, err
	}
	return sensorJSON, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	framework, err := arrowhead.CreateFramework()
	checkError(err)


	inMemorySensorRepository := &InMemorySensorRepository{
    sensor: &Sensor{Status: "False"},
	}
	
	updateSensorService := &UpdateSensorService{inMemorySensorRepository: inMemorySensorRepository}
	getSensorService := &GetSensorService{inMemorySensorRepository: inMemorySensorRepository}

	framework.HandleService(updateSensorService, rpc.POST, "update-sensor-status", "/magazine")
	framework.HandleService(getSensorService, rpc.GET, "get-sensor-status", "/magazine")

	err = framework.ServeForever()
	checkError(err)
}

