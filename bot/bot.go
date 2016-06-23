package bot

import (
	"github.com/hybridgroup/gobot"

	"github.com/hybridgroup/gobot/platforms/raspi"
)

var HeartBeat *gobot.Gobot
var Gbot *gobot.Gobot
var Operator *gobot.Robot
var Machine *gobot.Robot
var BotHandler *gobot.Robot

func NewBot() {
	gateway := NewGateway()
	Gbot = gobot.NewGobot()

	r  := gateway.raspiAdaptor

	Operator = NewOperator()
	Machine = NewMachineBot(r)

	Gbot.AddRobot(Operator)
	Gbot.AddRobot(Machine)

	Gbot.Start()
}

func initSensors (r *raspi.RaspiAdaptor) {
	//return models.NewThermometer(r)
	//TODO implement
}