package models

import (
	"fmt"

	"github.com/hybridgroup/gobot/platforms/gpio"
	//"github.com/hybridgroup/gobot/platforms/raspi"
)

type Thermometer struct {
	Temperature		float32                		`json:"temperature"`
	Gpio 			*gpio.AnalogSensorDriver 	`json:"gpio"`
}

/*
func NewThermometer (r *raspi.RaspiAdaptor) *Thermometer {
	th := new(Thermometer)
	//th.Gpio = gpio.NewAnalogSensorDriver(r, "Thermometer", "0")
	return th
}
*/

func (th *Thermometer) SetTemperature(f float32) {
	th.Temperature = f
	fmt.Print("Temperature = ", th.Temperature)
}



