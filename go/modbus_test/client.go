package main

import (
	"fmt"
	"log"
	"time"

	"github.com/simonvetter/modbus"
)

func main() {
	// Configure modbus Client
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:      "tcp://localhost:5020",
		Timeout:  1 * time.Second,
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
	fmt.Println("Connected to Modbus server at localhost:5020")

	// --- TEST 2: READ DATA ---
	// Read from coil 0
	var regAddr uint16 = 0

	results, err := client.ReadCoil(regAddr)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	fmt.Println("[READ] Current Sensor Value: ", results)

	// --- TEST 2: WRITE DATA ---
	// set new value to opposite of current value
    var newValue bool = !results  

    fmt.Printf("[WRITE] Changing value to: %t\n", newValue)
	
	err = client.WriteCoil(regAddr, newValue)
	if err != nil {
		log.Fatalln("Write failed: %v", err)
	}

	// --- TEST 3: VERIFY CHANGE ---
	results, err = client.ReadCoil(regAddr)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	fmt.Printf("[READ] New Sensor Value: %t\n", results) 
}