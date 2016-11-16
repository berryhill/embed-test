// +build ignore

package main

import (
	"fmt"
	"flag"
	"time"
	
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all"
)

const (
	address = 0x77

	calAc1          = 0xAA
	calAc2          = 0xAC
	calAc3          = 0xAE
	calAc4          = 0xB0
	calAc5          = 0xB2
	calAc6          = 0xB4
	calB1           = 0xB6
	calB2           = 0xB8
	calMB           = 0xBA
	calMC           = 0xBC
	calMD           = 0xBE
	control         = 0xF4
	tempData        = 0xF6
	pressureData    = 0xF6
	readTempCmd     = 0x2E
	readPressureCmd = 0x34

	tempReadDelay = 5 * time.Millisecond

	p0 = 101325

	pollDelay = 250
)

func main() {
	if err := embd.InitSPI(); err != nil {
		panic(err)
	}
	defer embd.CloseSPI()

	spiBus := embd.NewSPIBus(embd.SPIMode0, 0, 1000000, 8, 0)
	defer spiBus.Close()

	flag.Parse()

	if err := embd.InitI2C(); err != nil {
		panic(err)
	}
	defer embd.CloseI2C()

	i2cBus := embd.NewI2CBus(1)

	dataByte := byte(85)
	fmt.Println("starting")

	for k := 0; k < 100000; k++ { 
		dataReceived, err := spiBus.TransferAndReceiveByte(dataByte)
		if err != nil {
			panic(err)
		}

		if err :=  i2cBus.WriteByteToReg(address, control, readTempCmd); err != nil {
			panic(err)
		}
		time.Sleep(50 * time.Millisecond)
		temp, err := i2cBus.ReadWordFromReg(address, tempData)
		if err != nil {
			panic(err)
		}

		fmt.Println(temp)
		fmt.Println(dataReceived)
		time.Sleep(50 * time.Millisecond)

	}
}

