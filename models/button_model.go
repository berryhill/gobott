package models

import (
	"fmt"
	"encoding/json"

	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
	"github.com/hybridgroup/gobot"
)

type Sensor interface {
	Listen()
	MarshalJson()
}

type Button struct {
	Name 		string                  `json:"name"`
	Pressed		bool                	`json:"pressed"`
	Gpio		*gpio.ButtonDriver	`json:"gpio"`
	//Action	 	func()
}

func NewButton(r *raspi.RaspiAdaptor) *Button {
	b := new(Button)
	b.Name = "Test"
	b.Gpio = gpio.NewButtonDriver(r, "button", "11")
	fmt.Println(b)
	return b
}

func (b *Button) Listen() {
	gobot.On(b.Gpio.Event("push"), func(data interface{}) {
		b.Pressed = true
		fmt.Println("button pressed")
	})

	gobot.On(b.Gpio.Event("release"), func(data interface{}) {
		b.Pressed = false
		fmt.Println("button released")
	})
}

func (b *Button) MarshalJson() []byte {
	json, err := json.Marshal(b)
	//json, _ := json.MarshalIndent(b, "", "   ")

	if err != nil {
		fmt.Println(err)
		return json
	}

	return json
}






