// +build ignore

package models

import (
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all"
)

type SPI_Device struct {
	Name					string
	ClockEnablePin 			*embd.DigitalPin		// High to Load, Low to Serialize
	ParallelLoadPin 		*embd.DigitalPin		// Low to Load, High to Serialize
}

func NewSPI_Device(clock_enable_index int, parallel_load_index int) *SPI_Device {
	sd := new(SPI_Device)
	sd.ClockEnablePin = clock_enable
	sd.ParallelLoadPin = parallel_load

	return sd
}

func (sd *SPI_Device) SPILoadState() {
	sd.ClockEnablePin.Write(High)
	sd.ParallelLoadPin.Write(Low)
}

func (sd *SPI_Device) SPIShiftState() {
	sd.ClockEnablePin.Write(Low)
	sd.ParallelLoadPin.Write(High)
}

