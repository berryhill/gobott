package bot

import (
	"github.com/hybridgroup/gobot"

	"github.com/gobott/store"
	"fmt"
	"github.com/hybridgroup/gobot/platforms/beaglebone"
	"gopkg.in/mgo.v2/bson"
)

var Adaptor *beaglebone.BeagleboneAdaptor
var HeartBeat *gobot.Gobot
var Gbot *gobot.Gobot
var Operator *gobot.Robot
var Machine *gobot.Robot
var BotHandler *gobot.Robot
//var Gateway *Gatewayq

type Bot struct {
	Gateway 		*Gateway
}

func NewBot() *Bot {
	store.InitDb()
	Adaptor = beaglebone.NewBeagleboneAdaptor("beaglebone")

	bot := new(Bot)
	gateway := new(Gateway)
	if bot.Gateway, _ = gateway.Retrieve(); bot.Gateway == nil {
		bot.Gateway = NewGateway()
	}

	if bot.Gateway.Id == "" {
		bot.Gateway.Id = bson.NewObjectId()
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