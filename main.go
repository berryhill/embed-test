// +build ignore

package main

import (
	"fmt"
	
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all"
)

func main() {
	if err := embd.InitSPI(); err != nil {
		panic(err)
	}
	defer embd.CloseSPI()

	spiBus := embd.NewSPIBus(embd.SPIMode0, 0, 1000000, 8, 0)
	defer spiBus.Close()

//	receivedByte, err := spiBus.ReceiveByte()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("received byte is:", receivedByte)
//	count := 0

	dataByte := byte(0)
	fmt.Println("starting")

	for k := 0; k < 100000; k++ { 
		dataReceived, err := spiBus.TransferAndReceiveByte(dataByte)
		if err != nil {
			panic(err)
		}
		
		fmt.Println(dataReceived)
	}
}

