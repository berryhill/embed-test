// +build ignore

package main

import (
	"fmt"
	
	"github.com/kidoman/embd"
	"github.com/kidoman/embd/sensor/bmp085"
	_ "github.com/kidoman/embd/host/all"
)

func main() {
	if err := embd.InitSPI(); err != nil {
		panic(err)
	}
	defer embd.CloseSPI()

	spiBus := embd.NewSPIBus(embd.SPIMode0, 0, 1000000, 8, 0)
	defer spiBus.Close()

	flag.Parse()

	if err = embd.InitI2C(); err != nil {
		panic(err)
	}
	defer embd.CloseI2C()

	bus := embd.NewI2CBus(1)

	baro := bmp085.New(bus)
	defer baro.Close()

	dataByte := byte(85)
	fmt.Println("starting")

	for k := 0; k < 100000; k++ { 
		dataReceived, err := spiBus.TransferAndReceiveByte(dataByte)
		if err != nil {
			panic(err)
		}

		temp, err := baro.Temperature()
		if err != nil {
			panic(err)
		}
		
		fmt.Println(dataReceived)

		temp, err := baro.Temperature()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Temp is %v\n", temp)
		pressure, err := baro.Pressure()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Pressure is %v\n", pressure)
		altitude, err := baro.Altitude()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Altitude is %v\n", altitude)

		time.Sleep(50 * time.Millisecond)

	}
}

