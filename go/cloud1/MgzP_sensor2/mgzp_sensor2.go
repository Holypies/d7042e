package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	arrowhead "github.com/johankristianss/arrowhead/pkg/arrowhead"
	"github.com/johankristianss/arrowhead/pkg/rpc"
	"github.com/simonvetter/modbus"
)

type Sensor struct {
	Status string `json:"status"`
}

type InMemorySensorRepository struct {
	sync.RWMutex
	sensor *Sensor
}

func (r *InMemorySensorRepository) GetStatus() string {
	r.RLock() // Lock for reading
	defer r.RUnlock()
	return r.sensor.Status
}

func (r *InMemorySensorRepository) UpdateStatus(newStatus string) {
	r.Lock() // Lock for writing
	defer r.Unlock()
	r.sensor.Status = newStatus
}

type UpdateSensorService struct {
	inMemorySensorRepository *InMemorySensorRepository
}

func (s *UpdateSensorService) HandleRequest(params *arrowhead.Params) ([]byte, error) {
	fmt.Println("UpdateSensorService called, triggering status toggle.")

	// Check for proper initialization (Guardrail against nil pointer panic)
	if s.inMemorySensorRepository.sensor == nil {
		return nil, fmt.Errorf("repository sensor is not initialized (nil)")
	}

	// Println. Get the current status
	currentStatus := s.inMemorySensorRepository.sensor.Status

	// Toggle the status based on its current value
	var newStatus string
	if currentStatus == "True" {
		newStatus = "False"
	} else {
		// This handles "False" and the initial zero-value state ""
		newStatus = "True"
	}
	fmt.Println("2")
	// Update the persistent state in the repository
	s.inMemorySensorRepository.sensor.Status = newStatus

	fmt.Printf("Sensor Status Updated: Toggled from **%s** to **%s**\n", currentStatus, newStatus)

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

func pollingService(repo *InMemorySensorRepository) {
	// create tcp modbus client
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     "tcp://localhost:5022",
		Timeout: 1 * time.Second,
	})
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	// Open the Connection
	err = client.Open()
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	defer client.Close() // Close connection when done
	fmt.Println("Connected to Modbus server at localhost:5022")

	// Read from coil 0
	var regAddr uint16 = 0

	for {

		results, err := client.ReadCoil(regAddr)
		if err != nil {
			log.Fatalf("Read failed: %v", err)
		}

		var polledStatusString string
		var polledSensorStatus = results

		if polledSensorStatus {
			polledStatusString = "True"
		} else {
			polledStatusString = "False"
		}

		// check if polled status is the same as current status and update the status if it isnt
		currentStatus := repo.GetStatus()
		if currentStatus != polledStatusString {
			repo.UpdateStatus(polledStatusString)
			fmt.Println("Sensor status updated to: ", polledStatusString)
		}

		// poll every 1ms (plc scanrate)
		time.Sleep(1 * time.Second)
		fmt.Println("Loop successful")
	}

}

func main() {
	fmt.Println("main")
	framework, err := arrowhead.CreateFramework()
	checkError(err)

	inMemorySensorRepository := &InMemorySensorRepository{
		sensor: &Sensor{Status: "False"},
	}

	fmt.Println("1")
	// monitor sensor in background
	go pollingService(inMemorySensorRepository)

	//	updateSensorService := &UpdateSensorService{inMemorySensorRepository: inMemorySensorRepository}
	getSensorService := &GetSensorService{inMemorySensorRepository: inMemorySensorRepository}

	//	framework.HandleService(updateSensorService, rpc.POST, "update-sensor-status", "/magazine")
	framework.HandleService(getSensorService, rpc.GET, "psen2-get-status", "/magazine")

	err = framework.ServeForever()
	checkError(err)

}
