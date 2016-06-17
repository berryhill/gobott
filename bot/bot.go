package bot

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"

	"github.com/gobott/models"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

type Bot struct {
	raspiAdaptor		*raspi.RaspiAdaptor
}

var HeartBeat *gobot.Gobot
var Gbot *gobot.Gobot
var Operator *gobot.Robot
var BotHandler *gobot.Robot

func NewBot() {
	bot := new(Bot)
	Gbot = gobot.NewGobot()

	r, _, leds := bot.initMicrocontroller()
	fmt.Println(leds)

	Operator = NewOperator()

	Gbot.AddRobot(Operator)
	Gbot.AddRobot(NewMachineBot(r))

	Gbot.Start()
}

func (b *Bot) initMicrocontroller() (*raspi.RaspiAdaptor, []*models.Button, *gpio.LedDriver) {
	r := raspi.NewRaspiAdaptor("raspi")

	var buttons []*models.Button
	buttons = append(buttons, models.NewButton(r))

	led := gpio.NewLedDriver(r, "led", "7")

	return r, buttons, led
}

func initSensors (r *raspi.RaspiAdaptor) {
	//return models.NewThermometer(r)
	//TODO implement
}