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

func NewBot() {
	bot := new(Bot)
	gbot := gobot.NewGobot()

	r, buttons, leds := bot.initMicrocontroller()
	fmt.Println(leds)

	gbot.AddRobot(NewOperator())
	gbot.AddRobot(NewPeripheralPanel(r, buttons))
	gbot.Start()
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
	// TODO implement
}