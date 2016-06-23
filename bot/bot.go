package bot

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/raspi"

	"github.com/gobott/store"
	"fmt"
)

var Adaptor *raspi.RaspiAdaptor
var HeartBeat *gobot.Gobot
var Gbot *gobot.Gobot
var Operator *gobot.Robot
var Machine *gobot.Robot
var BotHandler *gobot.Robot
//var Gateway *Gateway

type Bot struct {
	Gateway 		*Gateway
}

func NewBot() *Bot {
	store.InitDb()
	Adaptor = raspi.NewRaspiAdaptor("raspi")

	bot := new(Bot)
	gateway := new(Gateway)

	if bot.Gateway, _ = gateway.Retrieve(); bot.Gateway == nil {
		bot.Gateway = NewGateway()
	}

	fmt.Println("Machine ID: " + bot.Gateway.Id.String())

	Gbot = gobot.NewGobot()

	Operator = NewOperator()
	Machine = NewMachineBot(Adaptor)

	Gbot.AddRobot(Operator)
	Gbot.AddRobot(Machine)

	Gbot.Start()

	return bot
}