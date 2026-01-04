package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/simonvetter/modbus"
)

func main() {
	// Create a coil handler(we are using boolean sensors)
	handler := &CoilHandler{}

	// Configure the Server
	server, err := modbus.NewServer(&modbus.ServerConfiguration{
		URL:        "tcp://0.0.0.0:5026",
		Timeout:    30 * time.Second,
		MaxClients: 5,
	}, handler)

	if err != nil {
		fmt.Printf("failed to create server: %v\n", err)
		os.Exit(1)
	}

	// Start the Server
	err = server.Start()
	if err != nil {
		fmt.Printf("failed to start server: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("--------------------------------------------------")
	fmt.Println("Modbus Slave (COILS) Running on port 5026")
	fmt.Println("Interactive Mode:")
	fmt.Println("  [1] + Enter -> Set Sensor TRUE  (ON)")
	fmt.Println("  [2] + Enter -> Set Sensor FALSE (OFF)")
	fmt.Println("  [Q] + Enter -> Quit")
	fmt.Println("--------------------------------------------------")

	// manual input loop
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter Command: ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())

		handler.lock.Lock()

		switch input {
		case "1":
			handler.coilState = true
			fmt.Println(" >> STATUS UPDATED: ON")
		case "2":
			handler.coilState = false
			fmt.Println(" >> STATUS UPDATED: OFF")
		case "q", "Q":
			handler.lock.Unlock()
			fmt.Println("Stopping server...")
			server.Stop()
			return
		default:
			fmt.Println("Invalid command. Press 1 or 2.")
		}

		handler.lock.Unlock()
	}
}

// ---------------------------------------------------------
// CUSTOM HANDLER FOR COILS
// ---------------------------------------------------------

type CoilHandler struct {
	lock      sync.RWMutex
	coilState bool // <--- Changed from uint16 to bool
}

// HandleCoils is now the active method
// It handles Read Coils (FC01) and Write Single Coil (FC05)
func (h *CoilHandler) HandleCoils(req *modbus.CoilsRequest) (res []bool, err error) {
	// Optional: Filter Unit ID
	if req.UnitId != 1 {
		// allow default
	}

	h.lock.RLock()
	defer h.lock.RUnlock()

	for i := 0; i < int(req.Quantity); i++ {
		currentAddr := req.Addr + uint16(i)
		switch currentAddr {
		case 0:
			if req.IsWrite {
				// If a client writes to us, update our state
				h.coilState = req.Args[i]
				fmt.Printf("[Network Write] Client set Coil 0 to: %v\n", h.coilState)
			}
			res = append(res, h.coilState)
		default:
			err = modbus.ErrIllegalDataAddress
			return
		}
	}
	return
}

// ---------------------------------------------------------
// UNUSED HANDLERS
// ---------------------------------------------------------

// We disable Holding Registers now because we are strictly using Coils
func (h *CoilHandler) HandleHoldingRegisters(req *modbus.HoldingRegistersRequest) (res []uint16, err error) {
	err = modbus.ErrIllegalFunction
	return
}

func (h *CoilHandler) HandleDiscreteInputs(req *modbus.DiscreteInputsRequest) (res []bool, err error) {
	err = modbus.ErrIllegalFunction
	return
}

func (h *CoilHandler) HandleInputRegisters(req *modbus.InputRegistersRequest) (res []uint16, err error) {
	err = modbus.ErrIllegalFunction
	return
}
