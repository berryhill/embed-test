// +build ignore

package main

import (
	"fmt"
	//"flag"
	//"time"

	"github.com/embed-test/i2c"
	
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
	dataByte := byte(85)

	fmt.Println("starting")
	i2cBus, err := i2c.New(byte(15), 1)
	if err != nil {
		panic(err)
	}

	for {
		i2cBus.WriteByte(byte(85))
		_, err := spiBus.TransferAndReceiveByte(dataByte)
		if err != nil {
			panic(err)
		}
	}
}


