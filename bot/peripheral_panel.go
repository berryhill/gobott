package bot

import (
	"github.com/hybridgroup/gobot"

	"github.com/gobott/models"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

var Buttons []*models.Button
var ppwork func()

func init() {
	ppwork = func() {
		// Buttons[0].Listen
	}
}

func NewPeripheralPanel(r *raspi.RaspiAdaptor, buttons []*models.Button) *gobot.Robot {
	Buttons = buttons
	robot := gobot.NewRobot("Peripheral Bot", []gobot.Connection{r},
		[]gobot.Device{Buttons[0].Gpio}, ppwork,
	)

	return robot
}

func GetJson(inter models.Sensor) {
	inter.MarshalJson()
}

